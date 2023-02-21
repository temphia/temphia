"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdapterEditorAPI = void 0;
const base_1 = require("../base");
class AdapterEditorAPI {
    constructor(base_url, tenant_id, token) {
        this.base = new base_1.ApiBase(base_url, tenant_id, token);
    }
    perform_action(name, data) {
        return this.base.post(`/admin/adapter_editor/action/${name}`, data);
    }
    self_update(data) {
        return this.base.post("/admin/adapter_editor/", data);
    }
    self_reset() {
        return this.base.post("/admin/adapter_editor/reset", {});
    }
    // app
    list_apps() {
        return this.base.get("/admin/adapter_editor/app");
    }
    new_app(data) {
        return this.base.post("/admin/adapter_editor/app", data);
    }
    get_app(id) {
        return this.base.get(`/admin/adapter_editor/app/${id}`);
    }
    update_app(id, data) {
        return this.base.post(`/admin/adapter_editor/app/${id}`, data);
    }
    delete_app(id) {
        return this.base.delete(`/admin/adapter_editor/app/${id}`);
    }
    // hook
    list_hooks() {
        return this.base.get("/admin/adapter_editor/hook");
    }
    new_hook(data) {
        return this.base.post("/admin/adapter_editor/hook", data);
    }
    get_hook(id) {
        return this.base.get(`/admin/adapter_editor/hook/${id}`);
    }
    update_hook(id, data) {
        return this.base.post(`/admin/adapter_editor/hook/${id}`, data);
    }
    delete_hook(id) {
        return this.base.delete(`/admin/adapter_editor/hook/${id}`);
    }
}
exports.AdapterEditorAPI = AdapterEditorAPI;
