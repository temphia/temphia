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
    add(key, value) {
        return this.base.post(`/admin/plug_state/key`, {
            key,
            value,
        });
    }
    update(key, value) {
        return this.base.post(`/admin/plug_state/key/${key}`, {
            key,
            value,
        });
    }
    delete(key) {
        return this.base.delete(`/admin/plug_state/key/${key}`, {
            key,
        });
    }
    get(key) {
        return this.base.delete(`/admin/plug_state/key/${key}`);
    }
}
exports.AdminPlugStateTktAPI = AdminPlugStateTktAPI;
