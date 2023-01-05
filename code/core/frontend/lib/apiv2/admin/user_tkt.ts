import type { ApiBase } from "../base";

export class AdminUserTktAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/admin/user_tkt/");
  }

  get(uid: string) {
    return this.base.get(`/admin/user_tkt/${uid}`);
  }

  new(data: any) {
    return this.base.post(`/admin/user_tkt/`, data);
  }

  update(uid: string, data: any) {
    return this.base.post(`/admin/user_tkt/${uid}`, data);
  }

  delete(uid: string) {
    return this.base.delete(`/admin/user_tkt/${uid}`);
  }
}
