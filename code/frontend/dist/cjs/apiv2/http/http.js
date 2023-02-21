"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.Http = void 0;
class Http {
    constructor(baseURL, headers) {
        this.baseURL = baseURL;
        this.headers = headers;
    }
    replace_headers(headers) {
        this.headers = headers;
    }
    async get(path) {
        const resp = await fetch(`${this.baseURL}${path}`, {
            method: "GET",
            headers: this.headers,
        });
        if (resp.ok) {
            const text = await resp.text();
            try {
                const data = JSON.parse(text);
                return Promise.resolve({
                    ok: true,
                    data,
                    status: resp.status,
                });
            }
            catch (error) {
                return Promise.resolve({
                    ok: true,
                    data: text,
                    status: resp.status,
                });
            }
        }
        return Promise.resolve({
            ok: false,
            data: await resp.text(),
            status: resp.status,
        });
    }
    async post(path, data) {
        return this.jsonMethod(path, "POST", data);
    }
    async patch(path, data) {
        return this.jsonMethod(path, "PATCH", data);
    }
    async put(path, data) {
        return this.jsonMethod(path, "PUT", data);
    }
    async jsonMethod(path, method, data) {
        const resp = await fetch(`${this.baseURL}${path}`, {
            method: method,
            headers: this.headers,
            body: JSON.stringify(data),
            mode: "cors",
        });
        if (resp.ok) {
            return Promise.resolve({
                ok: true,
                data: await resp.json(),
                status: resp.status,
            });
        }
        return Promise.resolve({
            ok: false,
            data: await resp.text(),
            status: resp.status,
        });
    }
    async postForm(path, auth, data) {
        return await fetch(`${this.baseURL}${path}`, {
            method: "POST",
            headers: auth ? { Authorization: this.headers["Authorization"] } : {},
            body: data,
        });
    }
    async patchForm(path, auth, data) {
        return await fetch(`${this.baseURL}${path}`, {
            method: "PATCH",
            headers: auth ? { Authorization: this.headers["Authorization"] } : {},
            body: data,
        });
    }
    async delete(path, data) {
        const resp = await fetch(`${this.baseURL}${path}`, {
            method: "DELETE",
            headers: this.headers,
            body: data ? JSON.stringify(data) : data,
        });
        if (resp.ok) {
            return Promise.resolve({
                ok: true,
                data: await resp.json(),
                status: resp.status,
            });
        }
        return Promise.resolve({
            ok: false,
            data: await resp.text(),
            status: resp.status,
        });
    }
}
exports.Http = Http;
