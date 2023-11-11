import type { ApiBase } from "../base";

export class AdminTenantAPI {
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

  get_domains() {
    return this.base.get("/admin/tenant/domain");
  }

  new_domain(data: any) {
    return this.base.post(`/admin/tenant/domain`, data);
  }

  get_domain(did: string) {
    return this.base.get(`/admin/tenant/domain/${did}`);
  }

  edit_domain(did: string, data: any) {
    return this.base.post(`/admin/tenant/domain/${did}`, data);
  }

  delete_domain(did: string) {
    return this.base.delete(`/admin/tenant/domain/${did}`);
  }

  domain_adapter_reset(did: string) {
    return this.base.get(`/admin/tenant/domain/${did}/reset`);
  }

  list_system_kv({ last, etype, prefix }) {
    const u = new URLSearchParams();
    u.set("last", last || "");
    u.set("etype", etype || "");
    u.set("prefix", prefix || "");

    return this.base.get(`/admin/tenant/system/kv?${u.toString()}`);
  }

  list_system_event({ last, etype }) {
    const u = new URLSearchParams();
    u.set("last", last || "");
    u.set("etype", etype || "");

    return this.base.get(`/admin/tenant/system/event?${u.toString()}`);
  }
}
