"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.ExecAPI = void 0;
class ExecAPI {
    constructor(base_url, exec_token) {
        this.base_url = base_url;
        this.exec_token = exec_token;
    }
    agent_file_url(pid, aid, file) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/serve/${file}`;
    }
    executor_file_url(eid, pid, aid, file) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/executor/${eid}/${file}`;
    }
    ws_url(room_token) {
        return `${this.base_url}/engine/ws?room_token=${room_token}`;
    }
    async preform_action(method, data) {
        const url = `${this.base_url}/engine/fixme`;
        const response = await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });
    }
}
exports.ExecAPI = ExecAPI;
