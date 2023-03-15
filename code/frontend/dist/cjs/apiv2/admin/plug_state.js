"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdminPlugStateTktAPI = void 0;
const http_1 = require("../http");
class AdminPlugStateTktAPI {
    constructor(api_base_url, token) {
        this.http = new http_1.Http(api_base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
        this.token = token;
        this.api_base_url = api_base_url;
    }
    query(options) {
        return this.http.post(`/admin/plug_state/query`, options);
    }
    add(key, value, opts) {
        return this.http.post(`/admin/plug_state/key`, {
            key,
            value,
            options: opts,
        });
    }
    update(key, value, opts) {
        return this.http.post(`/admin/plug_state/key/${key}`, {
            key,
            value,
            options: opts,
        });
    }
    delete(key) {
        return this.http.delete(`/admin/plug_state/key/${key}`, {
            key,
        });
    }
    get(key) {
        return this.http.get(`/admin/plug_state/key/${key}`);
    }
}
exports.AdminPlugStateTktAPI = AdminPlugStateTktAPI;
