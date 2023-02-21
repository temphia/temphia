"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminRepoAPI = void 0;
class AdminRepoAPI {
    constructor(base) {
        this.base = base;
    }
    list() {
        return this.base.get("/admin/repo/");
    }
    new(data) {
        return this.base.post("/admin/repo/", data);
    }
    get(rid) {
        return this.base.get(`/admin/repo/${rid}`);
    }
    update(rid, data) {
        return this.base.post(`/admin/repo/${rid}`, data);
    }
    delete(rid) {
        return this.base.delete(`/admin/repo/${rid}`);
    }
}
exports.AdminRepoAPI = AdminRepoAPI;
