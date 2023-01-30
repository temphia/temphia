import { sleep } from "yootils";
import { DataAPI, FolderTktAPI } from "../../../../../lib/apiv2";
import type { Sockd, SockdMessage } from "../../../../../lib/sockd";
import type { SockdService } from "../../sockd/sockd";
import { TableService } from "./table";

export class GroupService {
  source: string;
  name: string;
  data_api: DataAPI;
  api_base_url: string;
  active_tables: Map<string, TableService>;

  tables: object[];
  close_modal: any;
  open_modal: any;
  folder_api: FolderTktAPI;

  // sockd
  sockd_builder: SockdService;
  sockd_conn: Sockd;

  constructor(opts: {
    source: string;
    name: string;
    api: DataAPI;
    sockd_builder: SockdService;
    api_base_url: string;
    close_modal: any;
    open_modal: any;
  }) {
    this.source = opts.source;
    this.name = opts.name;
    this.data_api = opts.api;
    this.active_tables = new Map();
    this.tables = [];
    this.api_base_url = opts.api_base_url;

    this.sockd_builder = opts.sockd_builder;
    this.close_modal = opts.close_modal;
    this.open_modal = opts.open_modal;
  }

  init = async () => {
    const resp = await this.data_api.load();
    if (!resp.ok) {
      console.log("GROUP_SERVICE_INIT_ERR", resp);
      return;
    }

    this.tables = resp.data["tables"] || [];
    const folder_ticket = resp.data["folder_ticket"] || "";
    this.folder_api = new FolderTktAPI(this.api_base_url, folder_ticket);

    this.sockd_conn = await this.sockd_builder.build(
      this.data_api.sockd_url(),
      this.__sockd_handle
    );
  };

  __sockd_handle = (msg: SockdMessage) => {
    console.log("MESSAGE FROM DATA WS SOCKD THING|>", msg);
  };

  default_table = () => {
    return (this.tables[0] || {})["slug"] || undefined;
  };

  table_service = async (table: string) => {
    let tservice = this.active_tables.get(table);
    if (!tservice) {
      tservice = new TableService({
        all_tables: this.tables,
        data_api: this.data_api,
        group_slug: this.name,
        table_slug: table,
        folder_api: this.folder_api,
        close_modal: this.close_modal,
        open_modal: this.open_modal,
      });

      await tservice.init();
      this.active_tables.set(table, tservice);
    }

    return tservice;
  };

  run = async () => {
    console.log("Starting event loop");
    while (true) {
      console.log("Sleeping");
      await sleep(1000);
    }
  };

  close = async () => {
    this.active_tables.forEach((tsvc) => {
      tsvc.close();
    });
  };
}
