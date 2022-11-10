import { sleep } from "yootils";
import type { DataAPI } from "../../../../lib/apiv2";
import type { TableService } from "./table";

export class GroupService {
  source: string;
  name: string;
  api: DataAPI;
  active_tables: Map<string, TableService>;

  tables: object[];
  cabinet_ticket: string;

  /*
  
  fixme
    - sockd  
  */

  constructor(opts: { source: string; name: string; api: DataAPI }) {
    this.source = opts.source;
    this.name = opts.name;
    this.api = opts.api;
    this.active_tables = new Map();
    this.tables = [];
  }

  init = async () => {
    const resp = await this.api.load();
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

  table_service = () => {};

  run = async () => {
    console.log("Starting event loop");
    while (true) {
      console.log("Sleeping");
      await sleep(100);
    }
  };

  close = async () => {};
}
