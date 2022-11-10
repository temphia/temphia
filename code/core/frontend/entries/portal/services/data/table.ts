import { writable, Writable } from "svelte/store";
import type { DataAPI } from "../../../../lib/apiv2";
import type { DirtyData, NavData } from "./dtypes";
import { defaultViewData } from "./dtypes";

export class TableService {
  state: TableState;

  all_tables: object[];
  table_name: string;
  group_name: string;
  data_api: DataAPI;

  constructor(opts: {
    all_tables: object[];
    table_name: string;
    group_name: string;
    data_api: DataAPI;
  }) {
    this.all_tables = opts.all_tables;
    this.table_name = opts.table_name;
    this.group_name = opts.group_name;
    this.data_api = opts.data_api;

    this.state = new TableState(this);
  }

  apply_remote_changes = () => {};
  poll = async () => {};
}

export class TableState {
  dirty_store: Writable<DirtyData>;
  nav_store: Writable<NavData>;
  last_loading: number;
  svc: TableService;

  // reflected from subcribe to nav store.
  _loading: boolean;

  constructor(service: TableService) {
    this.svc = service;
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

    this.nav_store.subscribe((ndata) => console.log("NAV_DATA |>", ndata));
    this.nav_store.subscribe((ndata) => (this._loading = ndata.loading));
  }
}
