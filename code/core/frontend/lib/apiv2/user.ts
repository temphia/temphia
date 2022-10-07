import type { ApiBase } from "./base";

export class UserAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list() {
    return this.base.get("/user/");
  }

  get(id: string) {
    return this.base.get(`/user/${id}`);
  }

  post(id: string, data: any) {
    return this.base.post(`/user/${id}`, data);
  }
}