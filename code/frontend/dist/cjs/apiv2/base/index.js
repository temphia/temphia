"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ApiBase = void 0;
const site_1 = require("../../utils/site");
const http_1 = require("../http");
class ApiBase {
    constructor(api_base_url, tenant_id, token) {
        this.api_base_url = api_base_url;
        this.tenant_id = tenant_id;
        this.user_token = token;
        console.log("@api_base =>", this);
        this.http = new http_1.Http(api_base_url, {
            "Content-Type": "application/json",
            Authorization: token,
        });
    }
    async init() {
        // fixme => it should not create apiurl instead it should be passed 
        // base_url or sth and build api_base_url from it
        const resp = await fetch(`${site_1.apiURL(this.tenant_id)}/auth/refresh`, {
            method: "POST",
            body: JSON.stringify({
                path: ["basic"],
                user_token: this.user_token,
            }),
        });
        const rdata = await resp.json();
    }
    async get(path) {
        return this.http.get(path);
    }
    async post(path, data) {
        return this.http.post(path, data);
    }
    async postForm(path, auth, data) {
        return this.http.postForm(path, auth, data);
    }
    async patchForm(path, auth, data) {
        return this.http.patchForm(path, auth, data);
    }
    async put(path, data) {
        return this.http.put(path, data);
    }
    async patch(path, data) {
        return this.http.patch(path, data);
    }
    async delete(path, data) {
        return this.http.delete(path, data);
    }
}
exports.ApiBase = ApiBase;
