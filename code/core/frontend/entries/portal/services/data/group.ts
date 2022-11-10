import { sleep } from "yootils";
import type { DataAPI } from "../../../../lib/apiv2";

export class GroupService {
  source: string;
  name: string;
  api: DataAPI;

  tables: object[];

  constructor(opts: {
    source: string;
    name: string;
    api: DataAPI;
  }) {
    this.source = opts.source;
    this.name = opts.name;
    this.api = opts.api;
  }

  init = async (table?: string) => {
    const resp = await this.api.load();
    if (!resp.ok) {
        return
    }

    this.tables = resp.data;

    if (!table) {
      table = this.tables[0]["slug"];
    }
    if (!table) {
      return;
    }
  };

  run = async () => {
    console.log("Starting event loop");
    while (true) {
      console.log("Sleeping");
      await sleep(100);
    }
  };

  close = async () => {

  };
}
