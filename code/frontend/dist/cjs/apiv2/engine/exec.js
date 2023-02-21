"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExecAPI = void 0;
const http_1 = require("../http");
class ExecAPI {
    constructor(base_url, exec_token) {
        this.http = new http_1.Http(base_url, {
            Authorization: exec_token,
        });
    }
    agent_file_url(pid, aid, file) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/serve/${file}`;
    }
    executor_file_url(eid, pid, aid, file) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/executor/${eid}/${file}`;
    }
    ws_url(room_token) {
        return this.http.get(`/engine/ws?room_token=${room_token}`);
    }
    ws_update(room_token, data) {
        this.http.post(`/engine/ws?room_token=${room_token}`, data);
    }
    preform_action(method, data) {
        this.http.headers = {
            ...this.http.headers,
            "Access-Control-Allow-Origin": "*",
        };
        return this.http.post(`/engine/execute/${method}`, data);
    }
}
exports.ExecAPI = ExecAPI;
