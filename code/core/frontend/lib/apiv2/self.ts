import type { ApiBase } from "./base";

export class SelfAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list_cabinet_sources() {
    return this.base.get("/self/system/cabinet");
  }

  list_data_sources() {
    return this.base.get("/self/system/datatable");
  }

  list_adapter_providers() {
    return this.base.get("/self/system/adapter");
  }

  list_repo_sources() {
    return this.base.get("/self/system/repo");
  }

  load() {
    return this.base.get("/self/load");
  }

  self() {
    return this.base.get("/self/self");
  }

  user_profile(userid: string) {
    return this.base.get(`/self/user/${userid}`);
  }

  user_message(userid: string, message: string) {
    return this.base.post(`/self/user/${userid}`, message);
  }

  self_update(data: any) {
    return this.base.post("/self/self", data);
  }

  sessions() {
    return this.base.get("/self/session");
  }

  email_change(data: any) {
    return this.base.post("/self/email/change", data);
  }

  list_message() {
    return this.base.get("/self/message");
  }

  modify_message(data: any) {
    return this.base.post("/self/message", data);
  }

  issue_data(data: any) {
    return this.base.post("/self/issue/data", data);
  }
}
