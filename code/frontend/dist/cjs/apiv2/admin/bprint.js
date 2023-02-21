"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminBprintAPI = void 0;
class AdminBprintAPI {
    constructor(base) {
        this.base = base;
    }
    list() {
        return this.base.get("/admin/bprint/");
    }
    create(data) {
        return this.base.post("/admin/bprint/", data);
    }
    create_from_zip(data) {
        return this.base.patchForm("/admin/bprint/", true, data);
    }
    import(data) {
        return this.base.put("/admin/bprint/", data);
    }
    get(bid) {
        return this.base.get(`/admin/bprint/${bid}`);
    }
    update(bid, data) {
        return this.base.post(`/admin/bprint/${bid}`, data);
    }
    delete(bid) {
        return this.base.delete(`/admin/bprint/${bid}`);
    }
    list_file(bid) {
        return this.base.get(`/admin/bprint/${bid}/file`);
    }
    get_file(bid, file) {
        return this.base.get(`/admin/bprint/${bid}/file/${file}`);
    }
    add_file(bid, file, data) {
        return this.base.postForm(`/admin/bprint/${bid}/file/${file}`, true, data);
    }
    update_file(bid, file, data) {
        return this.base.patchForm(`/admin/bprint/${bid}/file/${file}`, true, data);
    }
    delete_file(bid, file) {
        return this.base.delete(`/admin/bprint/${bid}/file/${file}`);
    }
    instance(bid, data) {
        return this.base.post(`/admin/bprint/${bid}/instance`, data);
    }
    issue(bid, data) {
        return this.base.post(`/admin/bprint/${bid}/issue`, data);
    }
    issue_encoded(bid, data) {
        return this.base.post(`/admin/bprint/${bid}/issue/encoded`, data);
    }
    list_plugs(bid) {
        return this.base.get(`/admin/bprint/${bid}/plug`);
    }
}
exports.AdminBprintAPI = AdminBprintAPI;
