"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminUserTktAPI = void 0;
class AdminUserTktAPI {
    constructor(base) {
        this.base = base;
    }
    list() {
        return this.base.get("/admin/user_tkt/");
    }
    get(uid) {
        return this.base.get(`/admin/user_tkt/${uid}`);
    }
    new(data) {
        return this.base.post(`/admin/user_tkt/`, data);
    }
    update(uid, data) {
        return this.base.post(`/admin/user_tkt/${uid}`, data);
    }
    delete(uid) {
        return this.base.delete(`/admin/user_tkt/${uid}`);
    }
}
exports.AdminUserTktAPI = AdminUserTktAPI;
