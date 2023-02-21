"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PlugDevTktAPI = void 0;
const http_1 = require("../http");
class PlugDevTktAPI {
    constructor(base_url, token) {
        this.http = new http_1.Http(base_url, {
            Authorization: token,
        });
        this.token = token;
        this.base_url = base_url;
    }
    bprint_list_file() {
        return this.http.get("/dev/bprint/file");
    }
    bprint_push_file(data) {
        return this.http.postForm("/dev/bprint/file", true, data);
    }
    bprint_get_file(file) {
        return this.http.get(`/dev/bprint/file/${file}`);
    }
    bprint_del_file(file) {
        return this.http.delete(`/dev/bprint/file/${file}`);
    }
    exec_watch_agents_url(pid, aid) {
        return `${this.base_url}/dev/exec/watch/plug/${pid}/agent/${aid}`;
    }
    exec_reset_plug(pid, aid, data) {
        return this.http.post(`/dev/exec/reset/plug/${pid}/agent/${aid}`, data);
    }
    exec_run_agent_action(pid, aid, action, data) {
        return this.http.post(`/dev/exec/run/plug/${pid}/agent/${aid}/${action}`, data);
    }
}
exports.PlugDevTktAPI = PlugDevTktAPI;
