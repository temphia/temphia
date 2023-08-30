import type { ApiBase } from "../base";

export class AdminTargetAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  // app

  listApp() {
    return this.base.get(`/admin/target/app`);
  }
  listAppByType(ttype: string, target="") {
    return this.base.get(`/admin/target/app/${ttype}?target=${target}`);
  }
  newApp(ttype: string, data: any) {
    return this.base.post(`/admin/target/app/${ttype}`, data);
  }
  getApp(ttype: string, id: number) {
    return this.base.get(`/admin/target/app/${ttype}/${id}`);
  }
  updateApp(ttype: string, id: number, data: any) {
    return this.base.post(`/admin/target/app/${ttype}/${id}`, data);
  }
  deleteApp(ttype: string, id: number) {
    return this.base.delete(`/admin/target/app/${ttype}/${id}`);
  }

  // hooks

  listHook() {
    return this.base.get(`/admin/target/hook`);
  }
  listHookByType(ttype: string, target="") {
    return this.base.get(`/admin/target/hook/${ttype}?target=${target}`);
  }
  newHook(ttype: string, data: any) {
    return this.base.post(`/admin/target/hook/${ttype}`, data);
  }
  getHook(ttype: string, id: number) {
    return this.base.get(`/admin/target/hook/${ttype}/${id}`);
  }
  updateHook(ttype: string, id: number, data: any) {
    return this.base.post(`/admin/target/hook/${ttype}/${id}`, data);
  }
  deleteHook(ttype: string, id: number) {
    return this.base.delete(`/admin/target/hook/${ttype}/${id}`);
  }
}
