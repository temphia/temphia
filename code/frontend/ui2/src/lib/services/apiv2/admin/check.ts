import type { ApiBase } from "../base";

export class AdminCheckAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  bprint(bid: string) {
    this.base.get(`check/bprint/${bid}`);
  }

  plug(pid: string) {
    this.base.get(`check/plug/${pid}`);
  }

  dataGroup(gid: string) {
    this.base.get(`check/dgroup/${gid}`);
  }
}