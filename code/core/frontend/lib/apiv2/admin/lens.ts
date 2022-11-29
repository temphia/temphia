import type { ApiBase } from "../base";

export class AdminLensAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  query(index: string, qopts: any) {
    return this.base.post(`/admin/lens/${index}`, qopts);
  }
}
