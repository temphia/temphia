export class ExecAPI {
    exec_token: string;
    base_url: string;

    constructor(base_url: string, exec_token: string) {
        this.base_url = base_url
        this.exec_token = exec_token
    }

    agent_file_url(pid: string, aid: string, file: string) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/serve/${file}`;
    }

    executor_file_url(eid: string, pid: string, aid: string, file: string) {
        return `${this.base_url}/engine/plug/${pid}/agent/${aid}/executor/${eid}/${file}`;
    }

    ws_url(room_token: string) {
        return `${this.base_url}/engine/ws?room_token=${room_token}`
    }


    async preform_action(method: string, data: any) {
        const url = `${this.base_url}/engine/fixme`

        const response = await fetch(url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(data),
        });
    }
}