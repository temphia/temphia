"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminPlugStateTktAPI = void 0;
class AdminPlugStateTktAPI {
    constructor(base) {
        this.base = base;
    }
    list(qparms) {
        const u = new URLSearchParams();
        u.set("page", String(qparms.page || 0));
        u.set("page_count", String(qparms.page_count || 0));
        u.set("key_cursor", qparms.key_cursor || "");
        return this.base.get(`/admin/plug_state/`);
    }
    add(key, value) {
        return this.base.post(`/admin/plug_state/`, {
            key,
            value,
        });
    }
    update(key, value) {
        return this.base.post(`/admin/plug_state/${key}`, {
            key,
            value,
        });
    }
    delete(key) {
        return this.base.delete(`/admin/plug_state/${key}`, {
            key,
        });
    }
    get(key) {
        return this.base.delete(`/admin/plug_state/${key}`);
    }
}
exports.AdminPlugStateTktAPI = AdminPlugStateTktAPI;
