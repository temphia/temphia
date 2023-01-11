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

  update(bid: string, data: any) {
    return this.base.post(`/admin/bprint/${bid}`, data);
  }

  delete(bid: string) {
    return this.base.delete(`/admin/bprint/${bid}`);
  }

  list_file(bid: string) {
    return this.base.get(`/admin/bprint/${bid}/file`);
  }

  get_file(bid: string, file: string) {
    return this.base.get(`/admin/bprint/${bid}/file/${file}`);
  }

  add_file(bid: string, file: string, data: any) {    
    return this.base.postForm(`/admin/bprint/${bid}/file/${file}`, true, data);
  }
  
  update_file(bid: string, file: string, data: any) {
    return this.base.patchForm(`/admin/bprint/${bid}/file/${file}`, true, data);
  }

  delete_file(bid: string, file: string) {
    return this.base.delete(`/admin/bprint/${bid}/file/${file}`);
  }

  instance(bid: string, data: any) {
    return this.base.post(`/admin/bprint/${bid}/instance`, data);
  }

  issue(bid: string, data: any) {
    return this.base.post(`/admin/bprint/${bid}/issue`, data);
  }

  issue_encoded(bid: string, data: any) {
    return this.base.post(`/admin/bprint/${bid}/issue/encoded`, data);
  }

}
