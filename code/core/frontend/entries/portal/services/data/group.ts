import { sleep } from "yootils";
import type { DataAPI } from "../../../../lib/apiv2";

export class GroupService {
  source: string;
  name: string;
  api: DataAPI;

  tables: object[];

  constructor(opts: { source: string; name: string; api: DataAPI }) {
    this.source = opts.source;
    this.name = opts.name;
    this.api = opts.api;
  }

  init = async (table?: string) => {
    const resp = await this.api.load();
    if (!resp.ok) {
      return;
    }

    this.tables = resp.data;

    console.log("TABLES", this.tables)

    if (!table) {
      table = this.tables[0]["slug"];
    }
    if (!table) {
      return;
    }
  };

  default_table() {
    return this.tables[0]["slug"];
  }

  run = async () => {
    console.log("Starting event loop");
    while (true) {
      console.log("Sleeping");
      await sleep(100);
    }
  };

  close = async () => {};
}
