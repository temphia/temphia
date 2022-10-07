import type { ApiBase } from "../base";

export class BprintAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/admin/bprint/");
  }

  create(data: any) {
    return this.base.post("/admin/bprint/", data);
  }

  get(bid: string) {
    return this.base.get(`/admin/bprint/${bid}`);
  }

  post(bid: string, data: any) {
    return this.base.post(`/admin/bprint/${bid}`, data);
  }

  delete(bid: string) {
    return this.base.get(`/admin/bprint/${bid}`);
  }

  list_file(bid: string) {
    return this.base.get(`/admin/bprint/${bid}/file`);
  }

  get_file(bid: string, file: string) {
    return this.base.get(`/admin/bprint/${bid}/file/${file}`);
  }

  add_file(bid: string, file: string, data: any) {
    // fixme => make this formdata
    return this.base.post(`/admin/bprint/${bid}/file/${file}`, data);
  }

  update_file(bid: string, file: string, data: any) {
    // fixme => make this formdata
    return this.base.post(`/admin/bprint/${bid}/file/${file}`, data);
  }

  delete_file(bid: string, file: string) {
    return this.base.delete(`/admin/bprint/${bid}/file/${file}`);
  }

  instance(bid: string, data: any) {
    return this.base.post(`/admin/bprint/${bid}/instance`, data);
  }
}
