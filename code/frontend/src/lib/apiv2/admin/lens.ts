import type { ApiBase } from "../base";

export class AdminLensAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  query(qopts: any) {
    return this.base.post(`/admin/lens/query`, qopts);
  }
}
