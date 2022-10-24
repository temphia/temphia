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
}

