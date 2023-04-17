import { get, writable, Writable } from "svelte/store";
import type { DataAPI, FolderTktAPI } from "../../../../../lib/apiv2";
import { DirtyRowService } from "./dirty_service";
import { generate_column_order } from "./formatter";
import {
  DataState,
  DirtyData,
  NavData,
  ViewData,
  defaultViewData,
  ViewModeType,
} from "./state_types";
import {
  DataModification,
  DataModTypeDelete,
  DataModTypeInsert,
  DataModTypeUpdate,
  DataWidget,
  FilterItem,
  TableExecData,
} from "./table_types";

export class TableService {
  all_tables: object[];
  table_slug: string;
  group_slug: string;
  data_api: DataAPI;
  folder_api: FolderTktAPI;
  state: TableState;

  data_widgets: object[];
  rev_ref_columns: object[];

  _open_modal: (compo: any, props: object) => void;
  _close_modal: () => void;

  profile_generator: (string: any) => string;

  row_service: RowService;

  constructor(opts: {
    all_tables: object[];
    table_slug: string;
    group_slug: string;
    data_api: DataAPI;
    folder_api: FolderTktAPI;
    open_modal: (compo: any, props: object) => void;
    close_modal: () => void;
    profile_generator: (string: any) => string;
  }) {
    this.all_tables = opts.all_tables;
    this.table_slug = opts.table_slug;
    this.group_slug = opts.group_slug;
    this.data_api = opts.data_api;
    this.folder_api = opts.folder_api;
    this._open_modal = opts.open_modal;
    this._close_modal = opts.close_modal;
    this.profile_generator = opts.profile_generator;
    this.state = new TableState(this);
    this.row_service = new RowService(this, this.state);
  }

  refresh = () => {
    return this.init();
  };

  init = async (filters?: object[]) => {
    this.state.set_loading();

    const resp = await this.data_api.load_table(
      this.table_slug,
      filters ? { view_filters: filters } : {}
    );
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }

    const count = resp.data["count"] || 0;
    let page = 0;
    let selects = [];
    let filter_conds = [];

    const queryresp = resp.data["query_response"] || {};
    let last_page = queryresp["final"] || false;

    let view_mode: ViewModeType = "NONE";
    const views = resp.data["views"] || ([] as object[]);
    const active_view = resp.data["active_view"];

    if (resp.data["active_view"]) {
      const active = views.filter((v) => v["name"] !== active_view)[0];
      if (active) {
        view_mode = "STATIC";
      }
    } else if (filters) {
      view_mode = "MANUAL";
    }

    this.rev_ref_columns = resp.data["reverse_refs"] || [];
    this.data_widgets = resp.data["data_widgets"] || [];

    this.state.set_ok_loading(
      count,
      selects,
      filter_conds,
      last_page,
      page,
      view_mode
    );
    this.state.set_rows_data(queryresp, false);
  };

  on_sockd = (data: DataModification) => {
    if (data.mod_type === "comment") {
      console.log("@comment on row", data);
      return;
    }

    console.log("@processing data event", data);

    if (data.mod_type === DataModTypeDelete) {
      this.state.data_store.update((old) => {
        data.rows.forEach((row_id) => {
          delete old.indexed_rows[row_id];
          old.rows = old.rows.filter((v) => v !== row_id);
          console.log("@deleting_record", row_id);
        });

        return { ...old };
      });
      return;
    }

    if (this.state._nav_store.view_mode === "NONE") {
      if (data.mod_type === DataModTypeUpdate) {
        this.state.data_store.update((old) => {
          // if we do batch update data.data needs to be array

          const rows: object[] =
            data.rows.length == 1 ? [data.data] : data.data;

          data.rows.forEach((row_id, idx) => {
            const row = rows[idx];

            const oldrow = old.indexed_rows[row_id];
            if (!oldrow) {
              return;
            }

            const newrow = { ...oldrow, ...row };

            old.indexed_rows[row_id] = newrow;
          });

          console.log("@applying_update", data.data);

          return { ...old, indexed_rows: { ...old.indexed_rows } };
        });
      } else if (data.mod_type === DataModTypeInsert) {
        console.log("@applying_insert");
        if (this.state._nav_store.last_page) {
          console.log("@row/push");

          this.state.data_store.update((old) => {
            // if we support batch insert data.data needs to be array

            const rows: object[] =
              data.rows.length == 1 ? [data.data] : data.data;

            data.rows.forEach((row_id, idx) => {
              const row = rows[idx];
              old.indexed_rows[row_id] = row;

              console.log("@row", row, row_id);

              old.rows.push(row_id);
            });
            return { ...old };
          });
        } else {
          this.state.set_needs_refresh();
        }
      }
    } else {
      this.state.data_store.update((old) => {
        data.rows.forEach((row) => {
          old.marked_rows[row] = data.mod_type === "update" ? "blue" : "green";
        });

        console.log("@refresh/mark");

        return old;
      });
      this.state.set_needs_refresh();
    }
  };

  fts = (opts: { search_term; search_column; count }) => {
    return this.data_api.fts_query(this.table_slug, opts);
  };

  poll = async () => {};

  close = () => {};

  get_row_service = () => {
    return this.row_service;
  };

  apply_view = async (name: string, view: ViewData, modified: boolean) => {
    const data = await this.do_query({
      ...defaultViewData(),
      ...view,
    });
    if (!data) {
      console.warn("Could not fetch rows");
      return;
    }

    this.state.nav_store.update((state) => ({
      ...state,
      view_mode: modified ? "MANUAL" : "STATIC",
    }));

    this.state.set_rows_data(data, false);
  };

  // private

  private do_query = async (query: {
    count: number;
    filter_conds: object[];
    page: number;
    selects: string[];
    search_term: string;
  }) => {
    this.state.set_loading();

    const resp = await this.data_api.simple_query(this.table_slug, query);
    if (resp.status !== 200) {
      this.state.set_err_loading(resp.data);
      return;
    }

    let last_page = resp["final"] || false;

    const active_filter_conds = query.filter_conds as FilterItem[];

    this.state.set_ok_loading(
      query.count,
      active_filter_conds,
      query.selects,
      last_page,
      query.page
    );

    return resp.data;
  };

  reached_top = async () => {
    console.log("TOP REACHED");
    if (this.skip_loading()) {
      return;
    }
    console.log("FETCH MORE");
  };

  reached_buttom = async () => {
    console.log("@start_fetch_more");
    const navdata = this.state._nav_store;
    if (navdata.last_page) {
      if (this.skip_loading()) {
        console.warn("already last page");
        return;
      }
    }

    const data = await this.do_query({
      ...navdata.active_view,
      page: navdata.active_page + 1,
      count: 50,
    });
    if (!data) {
      console.warn("Could not fetch rows");
      return;
    }
    this.state.set_rows_data(data, true);
    console.log("@end_fetch_more");
  };

  private skip_loading = () => {
    if (this.state._nav_store.loading) {
      return true;
    }
    const now = new Date().valueOf();
    if (now - this.state.last_loading < 1000 * 10) {
      return true;
    }
    this.state.last_loading = now;
    return false;
  };

  get_invoker(widget: DataWidget) {
    return new TableInvoker(this, widget);
  }

  get_exec_data(rows: number[]): TableExecData {
    const state = get(this.state.data_store);

    const cells = {};

    rows.forEach((r) => {
      cells[r] = state.indexed_rows[r];
    });

    return {
      cells,
      data_group: this.group_slug,
      invoker_type: "data_table",
      rows: rows,
      table_id: this.table_slug,
      columns: Object.values(state.indexed_column),
    };
  }
}

export class TableState {
  dirty_store: Writable<DirtyData>;
  nav_store: Writable<NavData>;
  data_store: Writable<DataState>;

  last_loading: number;
  service: TableService;

  // reflected from subcribe to nav store.
  _nav_store: NavData;

  constructor(service: TableService) {
    this.service = service;
    this.last_loading = 0;

    this.dirty_store = writable({
      data: {},
      rowid: 0,
    });

    this.nav_store = writable({
      loading: true,
      lastTry: null,
      view_mode: "NONE",
      loading_error: "",
      active_page: 0,
      last_page: false,
      active_view: defaultViewData(),
      needs_refresh: false,
      views: {},
    });

    this.nav_store.subscribe((ndata) => (this._nav_store = ndata));
    this.data_store = writable({
      indexed_column: {},
      column_order: [],
      reverse_ref_column: [],
      rows: [],
      indexed_rows: {},
      sparse_rows: [],
      remote_dirty: {},
      views: [],
      ref_rows_cache: {},
      marked_rows: {},
    });

    // debug
    this.nav_store.subscribe((ndata) => console.log("NAV_STORE|>", ndata));
    this.data_store.subscribe((dstate) => console.log("DATA_STORE|>", dstate));
  }

  set_loading = () => {
    this.nav_store.update((old) => ({ ...old, loading: true }));
  };

  set_err_loading = (emessage: string) => {
    this.nav_store.update((old) => ({
      ...old,
      loading: false,
      loading_error: emessage,
    }));
  };

  set_ok_loading = (
    count: number,
    filter_conds: any,
    selects: string[],
    is_last_page: boolean,
    current_page: number,
    view_mode?: ViewModeType
  ) => {
    view_mode = view_mode || "NONE";
    this.nav_store.update((old) => ({
      ...old,
      loading: false,
      active_view: {
        count: count,
        filter_conds: filter_conds,
        main_column: "",
        search_term: "",
        selects: selects,
      },
      last_page: is_last_page,
      active_page: current_page,
      loading_error: "",
      view_mode,
    }));
  };

  set_needs_refresh = () => {
    this.nav_store.update((old) => ({ ...old, needs_refresh: true }));
  };

  set_row_data = (data: any) => {
    this.data_store.update((state) => {
      const row_id = data["__id"];

      let old_row = state.indexed_rows[row_id];
      if (old_row) {
        state.indexed_rows[row_id] = { ...old_row, ...data };
        return {
          ...state,
          indexed_rows: { ...state.indexed_rows },
        };
      }

      return state;
    });
  };

  set_rows_data = (data: any, append: boolean) => {
    if (!data["rows"]) {
      return;
    }

    this.data_store.update((old) => {
      const indexed_column = data["columns"];
      const column_order = generate_column_order(indexed_column);

      const old_rows = append ? old["rows"] || [] : [];
      const old_indexed = append ? old["indexed_rows"] || {} : {};
      let views = old["views"] || [];

      const _raw_rows = data["rows"]; //  [{ "__id": 1, xyz: "mno" }]
      const _rows = _raw_rows.map((row) => row["__id"]);

      // fixme => implement aesc and desc
      const rows = Array.from(new Set([..._rows, ...old_rows])).sort(
        (a, b) => a - b
      ); // only works for order_by "__id"

      const indexed_rows = _raw_rows.reduce(
        (result, curr) => {
          result[curr.__id] = curr;

          return result;
        },
        { ...old_indexed }
      );

      return {
        ...old,
        column_order,
        indexed_column,
        indexed_rows,
        rows,
        views,
      };
    });
  };

  start_row_edit = (row_id: number, data?: any) => {
    this.dirty_store.set({
      data: data || {},
      rowid: row_id,
    });
  };
}

export class RowService {
  service: TableService;
  state: TableState;

  constructor(service: TableService, state: TableState) {
    this.service = service;
    this.state = state;
  }

  get_dirty_service = () => {
    return new DirtyRowService(this.state.dirty_store);
  };

  ref_load = async (data: any) => {
    return this.service.data_api.ref_load(this.service.table_slug, data);
  };

  ref_resolve_pri = async (
    ref_type: string,
    target_table: string,
    target_column: string,
    ids: number[]
  ) => {
    return this.service.data_api.ref_resolve(this.service.table_slug, {
      type: ref_type,
      target: target_table,
      column: target_column,
      row_ids: ids,
    });
  };

  list_activity = async (rowId: number) => {
    return this.service.data_api.list_activity(
      this.service.table_slug,
      String(rowId)
    );
  };

  rev_ref_load = async (
    target_table: string,
    target_column: string,
    rowid: number
  ) => {
    return this.service.data_api.reverse_ref_load(this.service.table_slug, {
      current_table: this.service.table_slug,
      target_table: target_table,
      column: target_column,
      current_item: rowid,
    });
  };

  comment_row = async (rowId: number, message: string) => {
    return this.service.data_api.comment_row(
      this.service.table_slug,
      String(rowId),
      message
    );
  };

  delete_row = async (rowid: number) => {
    let resp = await this.service.data_api.delete_row(
      this.service.table_slug,
      String(rowid)
    );
    if (resp.status !== 200) {
      console.warn("could not delete row", resp);
      return;
    }
  };

  save_row = async () => {
    const api = this.service.data_api;
    const table_slug = this.service.table_slug;

    const dirtyData = get(this.state.dirty_store);
    if (dirtyData.rowid === 0) {
      const resp = await api.new_row(table_slug, {
        cells: dirtyData.data,
      });
      return resp.data;
    }

    const rowid = dirtyData.rowid;
    const data = get(this.state.data_store);

    const version = data.indexed_rows[rowid]["__version"];

    const resp = await api.update_row(table_slug, String(rowid), {
      cells: dirtyData.data,
      version,
    });

    if (resp.status !== 200) {
      console.warn("could not update row", resp);
      return resp.data;
    }

    this.state.set_row_data(resp.data);
    return resp.data;
  };

  fetch_row_latest = async (rowid: number) => {
    const resp = await this.service.data_api.get_row(
      this.service.table_slug,
      String(rowid)
    );
    if (resp.status !== 200) {
      return;
    }

    this.state.set_row_data(resp.data);
  };

  list_user = (colid: string) => {
    return this.service.data_api.list_users({
      target_type: "table",
      target: `${this.service.table_slug}/${colid}`,
    });
  };

  open_model(compo: any, opts: object) {
    this.service._open_modal(compo, opts);
  }

  close_model() {
    this.service._close_modal();
  }

  folder_api() {
    return this.service.folder_api;
  }
}

export class TableInvoker {
  service: TableService;
  widget: DataWidget;
  constructor(service: TableService, widget: DataWidget) {
    this.service = service;
    this.widget = widget;
  }

  handle = (instance_id: string, msg_id: string, data: any) => {
    console.log("@instance_handle", instance_id, msg_id, data);
  };

  close = (instance_id: string) => {
    console.log("@close", instance_id);
  };
}
