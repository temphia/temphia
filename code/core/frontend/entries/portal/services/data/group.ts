import { sleep } from "yootils";
import type { DataAPI } from "../../../../lib/apiv2";
import { TableService } from "./table";

export class GroupService {
  source: string;
  name: string;
  data_api: DataAPI;
  active_tables: Map<string, TableService>;

  tables: object[];
  cabinet_ticket: string;

  /*
  
  fixme
    - sockd  
    - cabinet/folder
  */

  constructor(opts: { source: string; name: string; api: DataAPI }) {
    this.source = opts.source;
    this.name = opts.name;
    this.data_api = opts.api;
    this.active_tables = new Map();
    this.tables = [];
  }

  init = async () => {
    const resp = await this.data_api.load();
    if (!resp.ok) {
      console.log("GROUP_SERVICE_INIT_ERR", resp);
      return;
    }

    this.tables = resp.data["tables"] || [];
    this.cabinet_ticket = resp.data["cabinet_ticket"] || "";
  };

  default_table = () => {
    return this.tables[0]["slug"] || undefined;
  };

  table_service = async (table: string) => {
    let tservice = this.active_tables.get(table);
    if (!tservice) {
      tservice = new TableService({
        all_tables: this.tables,
        data_api: this.data_api,
        group_slug: this.name,
        table_slug: table,

        close_modal: null,
        open_modal: null,
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
      await sleep(100);
    }
  };

  close = async () => {
    this.active_tables.forEach((tsvc) => {
      tsvc.close();
    });
  };
}
