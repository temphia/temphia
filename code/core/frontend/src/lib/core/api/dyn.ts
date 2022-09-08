import { ApiBase } from "./base";

export class DtableAPI extends ApiBase {
  constructor(url: string, user_token: string, source: string, group: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["dtable", source, group],
    });
  }

  async load_group() {
    return this.get(`/dgroup_load`);
  }

  // dtable
  async list_tables() {
    return this.get(`/dtable`);
  }

  async add_table(data: any) {
    return this.post(`/dtable`, data);
  }

  async edit_table(tid: string, data: any) {
    return this.patch(`/dtable/${tid}`, data);
  }

  async get_table(tid: string) {
    return this.get(`/dtable/${tid}`);
  }

  async delete_table(tid: string) {
    return this.delete(`/dtable/${tid}`);
  }

  async list_columns(tid: string) {
    return this.get(`/dtable/${tid}/column`);
  }
  async add_column(tid: string, data: any) {
    return this.post(`/dtable/${tid}/column`, data);
  }

  async get_column(tid: string, cid: string) {
    return this.get(`/dtable/${tid}/column/${cid}`);
  }

  async edit_column(tid: string, cid: string, data: any) {
    return this.patch(`/dtable/${tid}/column/${cid}`, data);
  }

  async delete_column(tid: string, cid: string) {
    return this.delete(`/dtable/${tid}/column/${cid}`);
  }

  // view stuff

  async list_view(tid: string) {
    return this.get(`/dtable/${tid}/view`);
  }

  async new_view(tid: string, data: any) {
    return this.post(`/dtable/${tid}/view`, data);
  }

  async modify_view(tid: string, id: number, data: any) {
    return this.post(`/dtable/${tid}/view/${id}`, data);
  }

  async get_view(tid: string, id: number) {
    return this.get(`/dtable/${tid}/view/${id}`);
  }

  async del_view(tid: string, id: number) {
    return this.delete(`/dtable/${tid}/view/${id}`);
  }

  // hook stuff
  async list_hook(tid: string) {
    return this.get(`/dtable/${tid}/hook`);
  }

  async new_hook(tid: string, data: any) {
    return this.post(`/dtable/${tid}/hook`, data);
  }

  async modify_hook(tid: string, id: number, data: any) {
    return this.post(`/dtable/${tid}/hook/${id}`, data);
  }

  async get_hook(tid: string, id: number) {
    return this.get(`/dtable/${tid}/hook/${id}`);
  }

  async del_hook(tid: string, id: number) {
    return this.delete(`/dtable/${tid}/hook/${id}`);
  }

  // dtable ops

  async new_row(tid: string, data: any) {
    return this.post(`/dtable_ops/${tid}/row`, data);
  }
  async get_row(tid: string, rid: number) {
    return this.get(`/dtable_ops/${tid}/row/${rid}`);
  }
  async update_row(tid: string, rid: number, data: any) {
    return this.post(`/dtable_ops/${tid}/row/${rid}`, data);
  }
  async delete_row(tid: string, rid: number) {
    return this.delete(`/dtable_ops/${tid}/row/${rid}`);
  }
  async simple_query(tid: string, data?: any) {
    if (!data) {
      data = {};
    }
    return this.post(`/dtable_ops/${tid}/simple_query`, data);
  }

  async fts_query(tid: string, str: string) {
    return this.post(`/dtable_ops/${tid}/fts_query`, {
      qstr: str,
    });
  }

  async ref_load(tid: string, data: any) {
    return this.post(`/dtable_ops/${tid}/ref_load`, data);
  }

  async ref_resolve(tid: string, data: any) {
    return this.post(`/dtable_ops/${tid}/ref_resolve`, data);
  }

  async rev_ref_load(tid: string, data) {
    return this.post(`/dtable_ops/${tid}/rev_ref_load`, data);
  }

  async list_activity(tid: string, rowid: number) {
    return this.get(`/dtable_ops/${tid}/activity/${rowid}`);
  }

  async comment_row(tid: string, rowid: number, msg: string) {
    return this.post(`/dtable_ops/${tid}/activity/${rowid}`, {
      message: msg,
    });
  }
}

export class DynAPI extends ApiBase {
  constructor(url: string, user_token: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["admin"],
    });
  }

  async list_group(source: string) {
    return this.get(`/dgroup/${source}`);
  }

  async get_group(source: string, group: string) {
    return this.get(`/dgroup/${source}/${group}`);
  }

  async new_group(source: string, data: any) {
    return this.post(`/dgroup/${source}`, data);
  }

  async edit_group(source: string, gid: string, data: any) {
    return this.patch(`/dgroup/${source}/${gid}`, data);
  }

  async delete_group(source: string, gid: string) {
    return this.delete(`/dgroup/${source}/${gid}`);
  }
}
