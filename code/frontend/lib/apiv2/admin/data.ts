import type { ApiBase } from "../base";

export class AdminDataAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list_group(source: string) {
    return this.base.get(`/admin/data/${source}/group`);
  }

  new_group(source: string, data: any) {
    return this.base.post(`/admin/data/${source}/group`, data);
  }

  get_group(source: string, gid: string) {
    return this.base.get(`/admin/data/${source}/group/${gid}`);
  }

  edit_group(source: string, gid: string, data: any) {
    return this.base.patch(`/admin/data/${source}/group/${gid}`, data);
  }

  delete_group(source: string, gid: string) {
    return this.base.delete(`/admin/data/${source}/group/${gid}`);
  }

  // table

  list_tables(source: string, gid: string) {
    return this.base.get(`/admin/data/${source}/group/${gid}/table`);
  }



  get_table(source: string, gid: string, tid: string) {
    return this.base.get(`/admin/data/${source}/group/${gid}/table/${tid}`);
  }

  edit_table(source: string, gid: string, tid: string, data: any) {
    return this.base.patch(
      `/admin/data/${source}/group/${gid}/table/${tid}`,
      data
    );
  }

  delete_table(source: string, gid: string, tid: string) {
    return this.base.delete(`/admin/data/${source}/group/${gid}/table/${tid}`);
  }

  // column

  list_column(source: string, gid: string, tid: string) {
    return this.base.get(
      `/admin/data/${source}/group/${gid}/table/${tid}/column`
    );
  }

  get_column(source: string, gid: string, tid: string, cid: string) {
    return this.base.get(
      `/admin/data/${source}/group/${gid}/table/${tid}/column/${cid}`
    );
  }

  edit_column(
    source: string,
    gid: string,
    tid: string,
    cid: string,
    data: any
  ) {
    return this.base.patch(
      `/admin/data/${source}/group/${gid}/table/${tid}/column/${cid}`,
      data
    );
  }

  delete_column(source: string, gid: string, tid: string, cid: string) {
    return this.base.delete(
      `/admin/data/${source}/group/${gid}/table/${tid}/column/${cid}`
    );
  }

  // view

  list_view(source: string, gid: string, tid: string) {
    return this.base.get(
      `/admin/data/${source}/group/${gid}/table/${tid}/view`
    );
  }

  add_view(source: string, gid: string, tid: string, data: any) {
    return this.base.post(
      `/admin/data/${source}/group/${gid}/table/${tid}/view`,
      data
    );
  }

  get_view(source: string, gid: string, tid: string, id: string) {
    return this.base.get(
      `/admin/data/${source}/group/${gid}/table/${tid}/view/${id}`
    );
  }

  edit_view(source: string, gid: string, tid: string, id: string, data: any) {
    return this.base.post(
      `/admin/data/${source}/group/${gid}/table/${tid}/view/${id}`,
      data
    );
  }

  delete_view(source: string, gid: string, tid: string, id: string) {
    return this.base.delete(
      `/admin/data/${source}/group/${gid}/table/${tid}/view/${id}`
    );
  }

  seed_table = (source: string, gid: string, tid: string, max: number) => {
    return this.base.get(
      `/admin/data/${source}/group/${gid}/table/${tid}/seed?max=${max}`
    );
  };

  query = (source: string, gid: string, opts: any) => {
    return this.base.post(`/admin/data/${source}/group/${gid}/query`, opts);
  };

  list_table_activity = (
    source: string,
    gid: string,
    tid: string,
    offset: number
  ) => {
    return this.base.get(
      `/admin/data/${source}/group/${gid}/table/${tid}/activity?offset=${offset}`
    );
  };
}
