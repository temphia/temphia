import type { ApiBase } from "../base";

export class AdminPlugStateTktAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  query(options: any) {
    return this.base.post(`/admin/plug_state/query`, options);
  }

  add(key: string, value: string) {
    return this.base.post(`/admin/plug_state/key`, {
      key,
      value,
    });
  }

  update(key: string, value: string) {
    return this.base.post(`/admin/plug_state/key/${key}`, {
      key,
      value,
    });
  }

  delete(key: string) {
    return this.base.delete(`/admin/plug_state/key/${key}`, {
      key,
    });
  }

  get(key: string) {
    return this.base.delete(`/admin/plug_state/key/${key}`);
  }
}
