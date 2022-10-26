import type { ApiBase } from "../base";

export class AdminBprintAPI {
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

  import(data: any) {
    return this.base.put("/admin/bprint/", data);
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

  // fixme => make this formdata
  add_file(bid: string, file: string, data: any) {
    
    return this.base.post(`/admin/bprint/${bid}/file/${file}`, data);
  }

  // fixme => make this formdata
  update_file(bid: string, file: string, data: any) {
    
    return this.base.post(`/admin/bprint/${bid}/file/${file}`, data);
  }

  delete_file(bid: string, file: string) {
    return this.base.delete(`/admin/bprint/${bid}/file/${file}`);
  }

  instance(bid: string, data: any) {
    return this.base.post(`/admin/bprint/${bid}/instance`, data);
  }
}
