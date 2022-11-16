import type { ApiBase } from "../base";

export class AdminLensAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  query_app(qopts: any) {
    return this.base.post("/admin/lens/app", qopts);
  }
  query_engine(qopts: any) {
    return this.base.post("/admin/lens/engine", qopts);
  }
  query_site(qopts: any) {
    return this.base.post("/admin/lens/site", qopts);
  }
}
