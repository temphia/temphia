"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminTargetAPI = void 0;
class AdminTargetAPI {
    constructor(base) {
        this.base = base;
    }
    // app
    listApp() {
        return this.base.get(`/admin/target/app`);
    }
    listAppByType(ttype, target = "") {
        return this.base.get(`/admin/target/app/${ttype}?target=${target}`);
    }
    newApp(ttype, data) {
        return this.base.post(`/admin/target/app/${ttype}`, data);
    }
    getApp(ttype, id) {
        return this.base.get(`/admin/target/app/${ttype}/${id}`);
    }
    updateApp(ttype, id, data) {
        return this.base.post(`/admin/target/app/${ttype}/${id}`, data);
    }
    deleteApp(ttype, id) {
        return this.base.delete(`/admin/target/app/${ttype}/${id}`);
    }
    // hooks
    listHook() {
        return this.base.get(`/admin/target/hook`);
    }
    listHookByType(ttype, target = "") {
        return this.base.get(`/admin/target/hook/${ttype}?target=${target}`);
    }
    newHook(ttype, data) {
        return this.base.post(`/admin/target/hook/${ttype}`, data);
    }
    getHook(ttype, id) {
        return this.base.get(`/admin/target/hook/${ttype}/${id}`);
    }
    updateHook(ttype, id, data) {
        return this.base.post(`/admin/target/hook/${ttype}/${id}`, data);
    }
    deleteHook(ttype, id) {
        return this.base.delete(`/admin/target/hook/${ttype}/${id}`);
    }
}
exports.AdminTargetAPI = AdminTargetAPI;
