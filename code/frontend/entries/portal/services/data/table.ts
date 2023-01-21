import { get, writable, Writable } from "svelte/store";
import type { DataAPI, FolderTktAPI } from "../../../../lib/apiv2";
import { generate_column_order } from "./data_utils";
import { Column, defaultViewData } from "./dtypes";
import type { DirtyData, FilterItem, NavData, ViewData } from "./dtypes";

export interface DataState {
  indexed_column: { [_: string]: Column };
  column_order: string[];
  reverse_ref_column: object[];

  rows: number[];
  indexed_rows: { [_: string]: object };

  sparse_rows: number[];
  remote_dirty: { [_: string]: true };
  views: object[];
}

export class TableService {
  all_tables: object[];
  table_slug: string;
  group_slug: string;
  data_api: DataAPI;
  folder_api: FolderTktAPI;
  state: TableState;

  _open_modal: (compo: any, props: object) => void;
  _close_modal: () => void;

  row_service: RowService;

  constructor(opts: {
    all_tables: object[];
    table_slug: string;
    group_slug: string;
    data_api: DataAPI;
    folder_api: FolderTktAPI;
    open_modal: (compo: any, props: object) => void;
    close_modal: () => void;
  }) {
    this.all_tables = opts.all_tables;
    this.table_slug = opts.table_slug;
    this.group_slug = opts.group_slug;
    this.data_api = opts.data_api;
    this.folder_api = opts.folder_api;
    this._open_modal = opts.open_modal;
    this._close_modal = opts.close_modal;

    this.state = new TableState(this);
    this.row_service = new RowService(this, this.state);
  }

  init = async () => {
    const data = await this.do_query({
      ...defaultViewData(),
      load_extra_meta: true,
    });

    this.state.set_rows_data(data, false);
  };

  apply_remote_changes = () => {};
  poll = async () => {};
  close = () => {};

  get_row_service = () => {
    return this.row_service;
  };

  apply_view = async (name: string, view: ViewData) => {
    const data = await this.do_query({
      ...defaultViewData(),
      ...view,
      load_extra_meta: false,
    });
    if (!data) {
      console.warn("Could not fetch rows");
      return;
    }

    this.state.set_rows_data(data, false);
  };

  // private

  private do_query = async (query: {
    count: number;
    filter_conds: object[];
    page: number;
    selects: string[];
    search_term: string;
    load_extra_meta: boolean;
  }) => {
    this.state.set_loading();

    const resp = await this.data_api.simple_query(this.table_slug, query);
    if (resp.status !== 200) {
      this.state.set_err_loading(resp.data);
      return;
    }

    let last_page = false;
    if (query.count > resp.data["rows"].length) {
      last_page = true;
    }

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
      load_extra_meta: false,
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

      loading_error: "",
      active_page: 0,
      last_page: false,
      active_view: defaultViewData(),

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
    current_page: number
  ) => {
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
    }));
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

      let reverse_ref_column = old["reverse_ref_column"] || [];
      let views = old["views"] || [];
      let hooks = old["hooks"] || [];

      const extra_meta = data["extra_meta"];

      if (extra_meta) {
        reverse_ref_column = extra_meta["reverse_refs"] || [];
        views = extra_meta["views"] || [];
        hooks = extra_meta["hooks"] || [];
      }

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
        reverse_ref_column,
        rows,
        hooks,
        views,
      };
    });
  };

  start_row_edit = (row_id: number) => {
    this.dirty_store.set({
      data: {},
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

export class DirtyRowService {
  dirtyStore: Writable<DirtyData>;
  callbacks: Map<string, () => void>;
  constructor(store: Writable<DirtyData>) {
    this.dirtyStore = store;
    this.callbacks = new Map();
  }

  register_before_save(field: string, callback: () => void): void {
    this.callbacks.set(field, callback);
  }

  on_ohange(_field: string, _value: any): void {
    this.set_value(_field, _value);
  }

  // row stuff
  start_modify_row = (row: number) => {
    this.callbacks.clear();
    this.dirtyStore.set({ rowid: row, data: {} });
  };

  start_new_row = () => {
    this.callbacks.clear();
    this.dirtyStore.set({ rowid: 0, data: {} });
  };

  set_value = (_filed: string, value: any) => {
    this.dirtyStore.update((old) => ({
      ...old,
      data: { ...old.data, [_filed]: value },
    }));
  };

  clear_dirty_row = () => {
    this.dirtyStore.set({ rowid: 0, data: {} });
  };

  set_ref_copy(column: string, value: any) {
    this.dirtyStore.update((old) => ({
      ...old,
      data: { ...old.data, [column]: value },
    }));
  }

  before_save() {
    this.callbacks.forEach((val) => val());
  }
}
