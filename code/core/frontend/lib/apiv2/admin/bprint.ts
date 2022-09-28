import type { ApiBase } from "../base";

export class BprintAPI {
  apibase: ApiBase;
  constructor(base: ApiBase) {
    this.apibase = base;
  }

  list() {
    return this.apibase.get("/admin/bprint/");
  }

  create(data: any) {
    return this.apibase.post("/admin/bprint/", data);
  }

  get(bid: string) {
    return this.apibase.get(`/admin/bprint/${bid}`);
  }

  post(bid: string, data: any) {
    return this.apibase.post(`/admin/bprint/${bid}`, data);
  }

  delete(bid: string) {
    return this.apibase.get(`/admin/bprint/${bid}`);
  }

  list_file(bid: string) {
    return this.apibase.get(`/admin/bprint/${bid}/file`);
  }

  get_file(bid: string, file: string) {
    return this.apibase.get(`/admin/bprint/${bid}/file/${file}`);
  }

  add_file(bid: string, file: string, data: any) {
    // fixme => make this formdata
    return this.apibase.post(`/admin/bprint/${bid}/file/${file}`, data);
  }

  update_file(bid: string, file: string, data: any) {
    // fixme => make this formdata
    return this.apibase.post(`/admin/bprint/${bid}/file/${file}`, data);
  }

  delete_file(bid: string, file: string) {
    return this.apibase.delete(`/admin/bprint/${bid}/file/${file}`);
  }

  instance(bid: string, data: any) {
    return this.apibase.post(`/admin/bprint/${bid}/instance`, data);
  }
}
