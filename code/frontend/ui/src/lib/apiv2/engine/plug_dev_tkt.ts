import { Http } from "../http";

export class PlugDevTktAPI {
  http: Http;
  token: string;
  base_url: string;
  constructor(base_url: string, token: string) {
    this.http = new Http(base_url, {
      Authorization: token,
    });
    this.token = token;
    this.base_url = base_url;
  }

  bprint_list_file() {
    return this.http.get("/dev/bprint/file");
  }

  bprint_push_file(data: FormData) {
    return this.http.postForm("/dev/bprint/file", true, data);
  }

  bprint_get_file(file: string) {
    return this.http.get(`/dev/bprint/file/${file}`);
  }

  bprint_del_file(file: string) {
    return this.http.delete(`/dev/bprint/file/${file}`);
  }

  exec_watch_agents_url(pid: string, aid: string) {
    return `${this.base_url}/dev/exec/watch/plug/${pid}/agent/${aid}`;
  }

  exec_reset_plug(pid: string, aid: string, data: any) {
    return this.http.post(`/dev/exec/reset/plug/${pid}/agent/${aid}`, data);
  }

  exec_run_agent_action(pid: string, aid: string, action: string, data: any) {
    return this.http.post(
      `/dev/exec/run/plug/${pid}/agent/${aid}/${action}`,
      data
    );
  }
}
