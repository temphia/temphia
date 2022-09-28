import type { ApiBase } from "../base";

export class TenantAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  edit(data: any) {
    return this.base.post("/admin/tenant/", data);
  }

  get() {
    return this.base.get("/admin/tenant/");
  }

  get_domain() {
    return this.base.get("/admin/tenant/domain");
  }
  new_domain(data: any) {
    return this.base.post(`/admin/tenant/domain`, data);
  }

  edit_domain(did: string, data: any) {
    return this.base.post(`/admin/tenant/domain/${did}`, data);
  }

  delete_domain(did: string) {
    return this.base.delete(`/admin/tenant/domain/${did}`);
  }
}
