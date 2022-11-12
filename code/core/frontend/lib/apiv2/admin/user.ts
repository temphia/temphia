import type { ApiBase } from "../base";

export class AdminUserAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/admin/user/");
  }

  get(uid: string) {
    return this.base.get(`/admin/user/${uid}`);
  }
}
