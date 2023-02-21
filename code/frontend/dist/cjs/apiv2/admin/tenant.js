"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminTenantAPI = void 0;
class AdminTenantAPI {
    constructor(base) {
        this.base = base;
    }
    edit(data) {
        return this.base.post("/admin/tenant/", data);
    }
    get() {
        return this.base.get("/admin/tenant/");
    }
    get_domains() {
        return this.base.get("/admin/tenant/domain");
    }
    new_domain(data) {
        return this.base.post(`/admin/tenant/domain`, data);
    }
    get_domain(did) {
        return this.base.get(`/admin/tenant/domain/${did}`);
    }
    edit_domain(did, data) {
        return this.base.post(`/admin/tenant/domain/${did}`, data);
    }
    delete_domain(did) {
        return this.base.delete(`/admin/tenant/domain/${did}`);
    }
    domain_issue_adapter_editor(did) {
        return this.base.get(`/admin/tenant/domain/${did}/issue/adapter_editor`);
    }
    domain_adapter_reset(did) {
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
exports.AdminTenantAPI = AdminTenantAPI;
