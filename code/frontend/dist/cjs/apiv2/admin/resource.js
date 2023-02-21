"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminResourceAPI = void 0;
class AdminResourceAPI {
    constructor(base) {
        this.base = base;
    }
    list() {
        return this.base.get("/admin/resource/");
    }
    new(data) {
        return this.base.post("/admin/resource/", data);
    }
    get(rid) {
        return this.base.get(`/admin/resource/${rid}`);
    }
    update(rid, data) {
        return this.base.post(`/admin/resource/${rid}`, data);
    }
    delete(rid) {
        return this.base.delete(`/admin/resource/${rid}`);
    }
}
exports.AdminResourceAPI = AdminResourceAPI;
