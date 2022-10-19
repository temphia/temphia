import type { ApiBase } from "./base";

export class DataAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  newRow(tid: string) {
    return this.base.get(`/data/${tid}/row`);
  }

  getRow(tid: string, rid: string) {
    return this.base.get(`/data/${tid}/row/${rid}`);
  }

  updateRow(tid: string, rid: string, data: any) {
    return this.base.post(`/data/${tid}/row/${rid}`, data);
  }

  deleteRow(tid: string, rid: string) {
    return this.base.delete(`/data/${tid}/row/${rid}`);
  }

  simpleQuery(tid: string, query: any) {
    return this.base.post(`/data/${tid}/simple_query`, query);
  }

  refLoad(tid: string, data: any) {
    return this.base.post(`/data/${tid}/ref_load`, data);
  }

  refResolve(tid: string, data: any) {
    return this.base.post(`/data/${tid}/ref_resolve`, data);
  }

  reverseRefLoad(tid: string, data: any) {
    return this.base.post(`/data/${tid}/rev_ref_load`, data);
  }
  listActivity(tid: string, rid: string) {
    return this.base.get(`/data/${tid}/activity${rid}`);
  }

  commentRow(tid: string, rid: string, data: any) {
    return this.base.post(`/data/${tid}/activity${rid}`, data);
  }
}

