"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminUserGroupAPI = void 0;
class AdminUserGroupAPI {
    constructor(base) {
        this.base = base;
    }
    list() {
        return this.base.get("/admin/ugroup/");
    }
    new(data) {
        return this.base.post("/admin/ugroup/", data);
    }
    get(ugroup) {
        return this.base.get(`/admin/ugroup/${ugroup}`);
    }
    update(ugroup, data) {
        return this.base.post(`/admin/ugroup/${ugroup}`, data);
    }
    delete(ugroup) {
        return this.base.delete(`/admin/ugroup/${ugroup}`);
    }
    // data
    listData(ugroup) {
        return this.base.get(`/admin/ugroup/${ugroup}/data`);
    }
    newData(ugroup, data) {
        return this.base.post(`/admin/ugroup/${ugroup}/data`, data);
    }
    getData(ugroup, id) {
        return this.base.get(`/admin/ugroup/${ugroup}/data/${id}`);
    }
    updateData(ugroup, id, data) {
        return this.base.post(`/admin/ugroup/${ugroup}/data/${id}`, data);
    }
    deleteData(ugroup, id) {
        return this.base.delete(`/admin/ugroup/${ugroup}/data/${id}`);
    }
    // auth
    listAuth(ugroup) {
        return this.base.get(`/admin/ugroup/${ugroup}/auth`);
    }
    newAuth(ugroup, data) {
        return this.base.post(`/admin/ugroup/${ugroup}/auth`, data);
    }
    getAuth(ugroup, id) {
        return this.base.get(`/admin/ugroup/${ugroup}/auth/${id}`);
    }
    updateAuth(ugroup, id, data) {
        return this.base.post(`/admin/ugroup/${ugroup}/auth/${id}`, data);
    }
    deleteAuth(ugroup, id) {
        return this.base.delete(`/admin/ugroup/${ugroup}/auth/${id}`);
    }
}
exports.AdminUserGroupAPI = AdminUserGroupAPI;
