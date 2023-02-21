"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SelfAPI = void 0;
class SelfAPI {
    constructor(base) {
        this.base = base;
    }
    load() {
        return this.base.get("/self/load");
    }
    self() {
        return this.base.get("/self/self");
    }
    self_ws_url() {
        return `${this.base.api_base_url}/self/self/ws?token=${this.base.http.headers["Authorization"]}`;
    }
    user_profile(userid) {
        return this.base.get(`/self/user/${userid}`);
    }
    user_message(userid, message) {
        return this.base.post(`/self/user/${userid}`, message);
    }
    self_update(data) {
        return this.base.post("/self/self", data);
    }
    sessions() {
        return this.base.get("/self/session");
    }
    email_change(data) {
        return this.base.post("/self/email/change", data);
    }
    list_message() {
        return this.base.get("/self/message");
    }
    modify_message(data) {
        return this.base.post("/self/message", data);
    }
    issue_data(data) {
        return this.base.post("/self/issue/data", data);
    }
    issue_folder(data) {
        return this.base.post("/self/issue/folder", data);
    }
    issue_ugroup(data) {
        return this.base.post("/self/issue/ugroup", data);
    }
    // system
    list_cabinet_sources() {
        return this.base.get("/self/system/cabinet");
    }
    list_data_sources() {
        return this.base.get("/self/system/datatable");
    }
    list_adapter_providers() {
        return this.base.get("/self/system/adapter");
    }
    list_executors() {
        return this.base.get("/self/system/executor");
    }
    list_modules() {
        return this.base.get("/self/system/module");
    }
    list_repo_sources() {
        return this.base.get("/self/system/repo");
    }
    // device
    list_devices() {
        return this.base.get("/self/device/");
    }
    add_device(opts) {
        return this.base.post("/self/device/", opts);
    }
    get_device(id) {
        return this.base.get(`/self/device/${id}`);
    }
    delete_device(id) {
        return this.base.delete(`/self/device/${id}`);
    }
}
exports.SelfAPI = SelfAPI;
