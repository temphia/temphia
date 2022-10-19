import type { ApiBase } from "../base";

export class RepoAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/admin/repo/");
  }
  new(data: any) {
    return this.base.post("/admin/repo/", data);
  }

  get(rid: string) {
    return this.base.get(`/admin/repo/${rid}`);
  }

  update(rid: string, data: any) {
    return this.base.post(`/admin/repo/${rid}`, data);
  }
  delete(rid: string) {
    return this.base.get(`/admin/repo/${rid}`);
  }
}