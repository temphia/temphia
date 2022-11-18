import { Http } from "../http";

export class ExecAPI {
  http: Http;
  exec_token: string;
  base_url: string;

  constructor(base_url: string, exec_token: string) {
    this.http = new Http(base_url, {
      Authorization: exec_token,
    });
  }

  agent_file_url(pid: string, aid: string, file: string) {
    return `${this.base_url}/engine/plug/${pid}/agent/${aid}/serve/${file}`;
  }

  executor_file_url(eid: string, pid: string, aid: string, file: string) {
    return `${this.base_url}/engine/plug/${pid}/agent/${aid}/executor/${eid}/${file}`;
  }

  ws_url(room_token: string) {
    return this.http.get(`/engine/ws?room_token=${room_token}`);
  }

  ws_update(room_token: string, data: any) {
    this.http.post(`/engine/ws?room_token=${room_token}`, data);
  }

  preform_action(method: string, data: any) {
    return this.http.post(`/engine/execute/${method}`, data);
  }
}
