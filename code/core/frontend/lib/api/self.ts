import { ApiBase } from "./base";

export class SelfAPI extends ApiBase {
  constructor(url: string, user_token: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["admin"],
    });
  }

  async list_cabinet_sources() {
    return this.get(`/cabinet_sources`);
  }
  async list_dgroup_sources() {
    return this.get(`/dgroup`);
  }

  async message_user(data: any) {
    return this.post("/self/message_user", data);
  }

  async get_user_info(userid: string) {
    return this.get(`/self/get_user_info/${userid}`);
  }

  async get_self_info() {
    return this.get("/self/get_self_info");
  }

  async update_self_info(data: any) {
    return this.post("/self/get_self_info", data);
  }

  async self_change_email(data: any) {
    return this.post("/self/change_email", data);
  }

  async self_change_auth(data: any) {
    return this.post("/self/change_auth", data);
  }

  async list_messages(data: any) {
    return this.post("/self/list_messages", data);
  }

  async modify_messages(data: any) {
    return this.post("/self/modify_messages", data);
  }

  async dtable_change(data: any) {
    return this.post("/self/dtable_change", data);
  }

  get_session_token() {
    return this._session_token;
  }
}