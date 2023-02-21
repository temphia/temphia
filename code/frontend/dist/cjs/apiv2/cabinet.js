"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.FolderTktAPI = exports.CabinetAPI = void 0;
const http_1 = require("./http");
class CabinetAPI {
    constructor(source, base) {
        this.source = source;
        this.base = base;
    }
    listRoot() {
        return this.base.get(`/cabinet/${this.source}/`);
    }
    listFolder(folder) {
        return this.base.get(`/cabinet/${this.source}/${folder}`);
    }
    newFolder(folder) {
        return this.base.post(`/cabinet/${this.source}/${folder}`, {});
    }
    getFile(folder, fname) {
        return this.base.get(`/cabinet/${this.source}/${folder}/file/${fname}`);
    }
    uploadFile(folder, fname, data) {
        return this.base.postForm(`/cabinet/${this.source}/${folder}/file/${fname}`, true, data);
    }
    deleteFile(folder, fname) {
        return this.base.delete(`/cabinet/${this.source}/${folder}/file/${fname}`);
    }
    getFilePreview(folder, fname) {
        return `${this.base.api_base_url}/cabinet/${this.source}/${folder}/preview/${fname}?token=${this.base.user_token}`;
    }
}
exports.CabinetAPI = CabinetAPI;
class FolderTktAPI {
    constructor(baseUrl, token) {
        this.http = new http_1.Http(baseUrl, {});
        this.ticket = token;
        this.base_url = baseUrl;
    }
    //  /folder/:ticket/
    list() {
        return this.http.get(`/folder/${this.ticket}`);
    }
    getFile(file) {
        return this.http.get(`/folder/${this.ticket}/${file}`);
    }
    getFileUrl(file) {
        return `${this.base_url}/folder/${this.ticket}/${file}`;
    }
    getFilePreviewUrl(file) {
        return `${this.base_url}/folder/${this.ticket}/${file}/preview`;
    }
    uploadFile(file, data) {
        return this.http.post(`/folder/${this.ticket}/${file}`, data);
    }
    // downgraded_ticket() {}
    deleteFile(file) {
        return this.http.delete(`/folder/${this.ticket}/${file}`);
    }
}
exports.FolderTktAPI = FolderTktAPI;
