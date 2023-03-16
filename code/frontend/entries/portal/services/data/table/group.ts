import { DataAPI, FolderTktAPI } from "../../../../../lib/apiv2";
import {
  MESSAGE_SERVER_PUBLISH,
  Sockd,
  SockdMessage,
} from "../../../../../lib/sockd";
import type { SockdService } from "../../sockd/sockd";
import { TableService } from "./table_service";
import type { DataModification } from "./table_types";

export class GroupService {
  source: string;
  name: string;
  data_api: DataAPI;
  api_base_url: string;
  tables_services: Map<string, TableService>;
  active_table: string;

  tables: object[];
  close_modal: any;
  open_modal: any;
  folder_api: FolderTktAPI;

  // sockd
  sockd_builder: SockdService;
  sockd_conn: Sockd;
  event_timer: number;

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
    this.tables_services = new Map();
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
    if (msg.type !== MESSAGE_SERVER_PUBLISH) {
      return;
    }
    
    const payload = msg.payload as DataModification;
    const tablesvc = this.tables_services.get(payload.table);
    if (tablesvc) {
      tablesvc.on_sockd(payload);
    }
  };

  default_table = () => {
    return (this.tables[0] || {})["slug"] || undefined;
  };

  table_service = async (table: string) => {
    let tservice = this.tables_services.get(table);
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
      this.tables_services.set(table, tservice);
    }

    this.active_table = table;

    return tservice;
  };

  run = async () => {
    this.event_timer = setInterval(() => {
      this.tables_services.forEach((tservice) => {
        tservice.poll();
      });
    }, 10000);
  };

  close = async () => {
    clearInterval(this.event_timer);

    this.tables_services.forEach((tsvc) => {
      tsvc.close();
    });
  };
}
