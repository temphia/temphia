import { get, writable, Writable } from "svelte/store";
import type { EngineService, ExecInstance } from "../engine";
import type { SockdService } from "../sockd";
import type { DtableAPI } from "../../core/api";
import {
  defaultViewData,
  DirtyData,
  FilterItem,
  NavData,
  ViewData,
} from "./data_types";
import { RowEditor } from "./roweditor";
import type { CommonStore } from "./store";
import { FolderTktAPI } from "../../core/tktapi";

export interface TableOptions {
  tables: object[];
  group: string;
  source: string;
  cabinet_ticket: string;
  sockd_ticket: string;
  api: DtableAPI;
  current_table: string;
  store: CommonStore;
  engine_service: EngineService;
  sockd_svc: SockdService;
}

export class DataTableService {
  api: DtableAPI;
  dtable: string;
  store: CommonStore;
  tableData: object;

  dirtyStore: Writable<DirtyData>;
  navStore: Writable<NavData>;
  lastLoading: number;
  groupOpts: TableOptions;
  FolderTktAPI: FolderTktAPI;
  loading: boolean;
  row_editor: RowEditor;
  hook_executor: HookExecutor;
  sockd_svc: SockdService;

  constructor(opts: TableOptions) {
    this.api = opts.api;
    this.dtable = opts.current_table;
    this.store = opts.store;
    this.sockd_svc = opts.sockd_svc;

    this.groupOpts = opts;
    this.FolderTktAPI = new FolderTktAPI(opts.api._api_base_url, opts.cabinet_ticket);

    this.navStore = writable({
      loading: true,
      lastTry: null,

      loading_error: "",
      active_page: 0,
      last_page: false,
      active_view: defaultViewData(),

      views: {},
    });

    this.navStore.subscribe((val) => console.log("NAV_STORE @=> ", val));
    this.navStore.subscribe((navd) => (this.loading = navd.loading));
    this.lastLoading = 0;
    this.navStore.subscribe((data) => console.log(data));
    this.dirtyStore = writable({ data: {}, rowid: 0 });
    this.dirtyStore.subscribe((data) => console.log("DIRTY DATA", data));
    this.row_editor = new RowEditor(this.dirtyStore);
    this.hook_executor = new HookExecutor(opts.engine_service, this);
  }

  init = async () => {
    const data = await this.do_query({
      ...defaultViewData(),
      load_extra_meta: true,
    });
    if (!data) {
      console.warn("Could not fetch rows");
      return;
    }
    this.saveData(data, false);

    // const sockdroom = this.sockd_svc.get_dyn_room();
    // sockdroom.onMessage(this.sockd_data_sync);
    // this.sockd_svc.change_group(
    //   this.groupOpts.source,
    //   this.groupOpts.group,
    //   this.groupOpts.sockd_ticket
    // );
  };

  reset = async () => {
    await this.init();
  };

  reload = async () => {
    const data = await this.do_query({
      ...defaultViewData(),
      load_extra_meta: true,
    });
    if (!data) {
      console.warn("Could not fetch rows");
      return;
    }
    this.saveData(data, false);
  };

  saveData = (data: object, append: boolean) => {
    this.store.set_rows_data(this.dtable, data, append);
  };

  fetchRowLatest = async (rowid: number) => {
    const resp = await this.api.get_row(this.dtable, rowid);
    if (resp.status !== 200) {
      return;
    }
    this.store.set_row_data(this.dtable, resp.data);
  };

  deleteRow = async (rowid: number) => {
    let resp = await this.api.delete_row(this.dtable, rowid);
    if (resp.status !== 200) {
      console.warn("could not delete row", resp);
      return;
    }
  };

  applyView = async (name: string, view: ViewData) => {
    const data = await this.do_query({
      ...defaultViewData(),
      ...view,
      load_extra_meta: false,
    });
    if (!data) {
      console.warn("Could not fetch rows");
      return;
    }
    this.saveData(data, false);
  };

  reachedTop = async () => {
    console.log("TOP REACHED");
    if (this.skipLoading()) {
      return;
    }
    console.log("FETCH MORE");
  };

  reachedButtom = async () => {
    console.log("@start_fetch_more");

    const navdata = get(this.navStore);
    if (navdata.last_page) {
      if (this.skipLoading()) {
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
    this.saveData(data, true);
    console.log("@end_fetch_more");
  };

  private skipLoading = () => {
    if (this.loading) {
      return true;
    }
    const now = new Date().valueOf();
    if (now - this.lastLoading < 1000 * 10) {
      return true;
    }
    this.lastLoading = now;
    return false;
  };

  saveDirtyRow = () => {
    return this._saveRow();
  };

  private _saveRow = async () => {
    const dirtyData = get(this.dirtyStore);
    if (dirtyData.rowid === 0) {
      const resp = await this.api.new_row(this.dtable, {
        cells: dirtyData.data,
      });
      return resp.data;
    }

    const rowid = dirtyData.rowid;

    const data = get(this.store.states);

    const version = data[this.dtable].indexed_rows[rowid]["__version"];

    const resp = await this.api.update_row(this.dtable, rowid, {
      cells: dirtyData.data,
      version,
    });

    if (resp.status !== 200) {
      console.warn("could not update row", resp);
      return resp.data;
    }

    this.store.set_row_data(this.dtable, resp.data);
    return resp.data;
  };

  do_query = async (query: {
    count: number;
    filter_conds: object[];
    page: number;
    selects: string[];
    search_term: string;
    load_extra_meta: boolean;
  }) => {
    this.navStore.update((old) => ({ ...old, loading: true }));
    const resp = await this.api.simple_query(this.dtable, query);
    if (resp.status !== 200) {
      this.navStore.update((old) => ({
        ...old,
        loading: false,
        loading_error: resp.data,
      }));
      return;
    }

    let last_page = false;
    if (query.count > resp.data["rows"].length) {
      last_page = true;
    }

    const active_filter_conds = query.filter_conds as FilterItem[];
    this.navStore.update((old) => ({
      ...old,
      loading: false,
      active_view: {
        count: query.count,
        filter_conds: active_filter_conds,
        main_column: "",
        search_term: query.search_term,
        selects: query.selects,
      },

      last_page,
      active_page: query.page,
      loading_error: "",
    }));
    return resp.data;
  };

  close = () => {
    this.hook_executor.close();
  };

  ref_load = async (data: any) => {
    return this.api.ref_load(this.dtable, data);
  };

  ref_resolve_pri = async (
    ref_type: string,
    target_table: string,
    target_column: string,
    ids: number[]
  ) => {
    return this.api.ref_resolve(this.dtable, {
      type: ref_type,
      target: target_table,
      column: target_column,
      row_ids: ids,
    });
  };

  list_activity = async (rowId: number) => {
    return this.api.list_activity(this.dtable, rowId);
  };

  rev_ref_load = async (
    target_table: string,
    target_column: string,
    rowid: number
  ) => {
    return this.api.rev_ref_load(this.dtable, {
      current_table: this.dtable,
      target_table: target_table,
      column: target_column,
      current_item: rowid,
    });
  };

  comment_row = async (rowId: number, message: string) => {
    return this.api.comment_row(this.dtable, rowId, message);
  };

  set_ref_callback = (fn: () => HTMLElement) => {
    this.hook_executor.get_target_ref = fn;
  };

  sockd_data_sync = (message) => {
    /*
                    {
                "room": "sys.dtable",
                "type": "server_publish",
                "xid": "cadr9e0m4q742ae05jug",
                "payload": {
                  "table": "acc",
                  "mod_type": "update",
                  "data": {
                    "__mod_sig": "{\"user_id\":\"superuser\",\"table_name\":\"default1_sb1_acc\"}",
                    "__version": 1,
                    "fullname": "brunch😜"
                  }
                }
              }
    
        */
    if (message["type"] !== "server_publish") {
      return;
    }

    const payload = message["payload"];

    switch (payload["mode_type"]) {
      case "update":
        this.store.set_row_data(this.dtable, {
          __id: payload["rows"][0],
          ...payload["data"],
        });

        break;
      default:
        break;
    }
  };
}

export class HookExecutor {
  _engine: EngineService;
  _table_service: DataTableService;
  _active_execs: Map<number, ExecInstance>;
  get_target_ref: () => HTMLElement;

  constructor(e: EngineService, dts: DataTableService) {
    this._engine = e;
    this._table_service = dts;
    this._active_execs = new Map();
  }

  execute_hook = async (hook: object) => {
    // const hid = hook["id"];
    // let pexec: PlugExec = this._active_execs[hid];
    // if (!pexec) {
    //   pexec = await this._engine.instance_dataplug(hook);
    //   this._active_execs.set(hid, pexec);
    //   pexec.set_handler((xid: string, action: string, data: any) => {
    //     this.on_message(hid, data); // fixme => send action too
    //   });
    // }
    // // run or re_run
    // await pexec.run(this.get_target_ref(), {});
  };

  close = () => {
    this._active_execs.forEach((pexec) => {
      // pexec.close();
    });
  };

  // commands

  on_message = (hookid: number, data: any) => {
    console.log("HOOK@", hookid, "DATA", data);
  };

  command_data_hello = (message: string) => {
    console.log(message);
  };
}
