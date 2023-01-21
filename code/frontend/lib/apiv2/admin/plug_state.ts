import type { ApiBase } from "../base";

export class AdminPlugStateTktAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list(qparms: { key_cursor?: string; page?: number, page_count?: number }) {
    const u = new URLSearchParams();

    u.set("page", String(qparms.page || 0));
    u.set("page_count", String(qparms.page_count || 0));
    u.set("key_cursor", qparms.key_cursor || "");

    return this.base.get(`/admin/plug_state/`);
  }

  add(key: string, value: string) {
    return this.base.post(`/admin/plug_state/`, {
      key,
      value,
    });
  }

  update(key: string, value: string) {
    return this.base.post(`/admin/plug_state/${key}`, {
      key,
      value,
    });
  }

  delete(key: string) {
    return this.base.delete(`/admin/plug_state/${key}`, {
      key,
    });
  }

  get(key: string) {
    return this.base.delete(`/admin/plug_state/${key}`);
  }
}
