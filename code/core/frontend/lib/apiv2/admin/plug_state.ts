import type { ApiBase } from "../base";

export class AdminPlugStateAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list(pid: string) {
    return this.base.get(`/admin/plug/${pid}/state/`);
  }

  new(pid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/state/`, data);
  }

  get(pid: string, key: string) {
    return this.base.get(`/admin/plug/${pid}/state/${key}`);
  }

  update(pid: string, key: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/state/${key}`, data);
  }

  del(pid: string, key: string) {
    return this.base.delete(`/admin/plug/${pid}/state/${key}`);
  }
}