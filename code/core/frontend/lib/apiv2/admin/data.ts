import type { ApiBase } from "../base";

export class DataAPI {
  apibase: ApiBase;
  constructor(base: ApiBase) {
    this.apibase = base;
  }

  list_group(source: string) {
    return this.apibase.get(`/admin/data/${source}`);
  }

  new_group(source: string, data: any) {
    return this.apibase.post(`/admin/data/${source}`, data);
  }

  get_group(source: string, gid: string) {
    return this.apibase.get(`/admin/data/${source}/${gid}`);
  }

  delete_group(source: string, gid: string) {
    return this.apibase.delete(`/admin/data/${source}/${gid}`);
  }

  // table

  list_tables(source: string, gid: string) {
    return this.apibase.get(`/admin/data/${source}/${gid}/table`);
  }

  add_table(source: string, data: any, gid: string) {
    return this.apibase.post(`/admin/data/${source}/${gid}/table`, data);
  }

  get_table(source: string, gid: string, tid: string) {
    return this.apibase.get(`/admin/data/${source}/${gid}/table/${tid}`);
  }

  edit_table(source: string, gid: string, tid: string, data: any) {
    return this.apibase.post(`/admin/data/${source}/${gid}/table/${tid}`, data);
  }

  delete_table(source: string, gid: string, tid: string) {
    return this.apibase.delete(`/admin/data/${source}/${gid}/table/${tid}`);
  }

  // column

  list_column(source: string, gid: string, tid: string) {
    return this.apibase.get(`/admin/data/${source}/${gid}/table/${tid}/column`);
  }

  add_column(source: string, data: any, gid: string, tid: string) {
    return this.apibase.post(
      `/admin/data/${source}/${gid}/table/${tid}/column`,
      data
    );
  }

  get_column(source: string, gid: string, tid: string, cid: string) {
    return this.apibase.get(
      `/admin/data/${source}/${gid}/table/${tid}/column/${cid}`
    );
  }

  edit_column(
    source: string,
    gid: string,
    tid: string,
    cid: string,
    data: any
  ) {
    return this.apibase.post(
      `/admin/data/${source}/${gid}/table/${tid}/column/${cid}`,
      data
    );
  }

  delete_column(source: string, gid: string, tid: string, cid: string) {
    return this.apibase.delete(
      `/admin/data/${source}/${gid}/table/${tid}/column/${cid}`
    );
  }

  // view

  list_view(source: string, gid: string, tid: string) {
    return this.apibase.get(`/admin/data/${source}/${gid}/table/${tid}/view`);
  }

  add_view(source: string, data: any, gid: string, tid: string) {
    return this.apibase.post(
      `/admin/data/${source}/${gid}/table/${tid}/view`,
      data
    );
  }

  get_view(source: string, gid: string, tid: string, id: string) {
    return this.apibase.get(
      `/admin/data/${source}/${gid}/table/${tid}/view/${id}`
    );
  }

  edit_view(source: string, gid: string, tid: string, id: string, data: any) {
    return this.apibase.post(
      `/admin/data/${source}/${gid}/table/${tid}/view/${id}`,
      data
    );
  }

  delete_view(source: string, gid: string, tid: string, id: string) {
    return this.apibase.delete(
      `/admin/data/${source}/${gid}/table/${tid}/view/${id}`
    );
  }
}
