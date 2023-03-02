"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminPlugStateTktAPI = void 0;
class AdminPlugStateTktAPI {
    constructor(base) {
        this.base = base;
    }
    query(options) {
        return this.base.post(`/admin/plug_state/query`, options);
    }
    add(key, value, opts) {
        return this.base.post(`/admin/plug_state/key`, {
            key,
            value,
            options: opts,
        });
    }
    update(key, value, opts) {
        return this.base.post(`/admin/plug_state/key/${key}`, {
            key,
            value,
            options: opts
        });
    }
    delete(key) {
        return this.base.delete(`/admin/plug_state/key/${key}`, {
            key,
        });
    }
    get(key) {
        return this.base.get(`/admin/plug_state/key/${key}`);
    }
}
exports.AdminPlugStateTktAPI = AdminPlugStateTktAPI;
