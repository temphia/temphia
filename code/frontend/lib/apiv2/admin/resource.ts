import type { ApiBase } from "../base";

export class AdminResourceAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/admin/resource/");
  }
  new(data: any) {
    return this.base.post("/admin/resource/", data);
  }

  get(rid: string) {
    return this.base.get(`/admin/resource/${rid}`);
  }

  update(rid: string, data: any) {
    return this.base.post(`/admin/resource/${rid}`, data);
  }
  delete(rid: string) {
    return this.base.get(`/admin/resource/${rid}`);
  }
}