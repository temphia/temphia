import type { ApiBase } from "../base";

export class AdminPlugAPI {
  base: ApiBase;
  constructor(base: ApiBase) {
    this.base = base;
  }

  list_plug() {
    return this.base.get("/admin/plug/");
  }

  new_plug(data: any) {
    return this.base.post("/admin/plug/", data);
  }

  get_plug(pid: string) {
    return this.base.get(`/admin/plug/${pid}`);
  }

  update_plug(pid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}`, data);
  }

  delete_plug(pid: string) {
    return this.base.delete(`/admin/plug/${pid}`);
  }

  list_plug_resource(pid: string) {
    return this.base.get(`/admin/plug/${pid}/resource`);
  }

  // agent

  list_agent(pid: string) {
    return this.base.get(`/admin/plug/${pid}/agent`);
  }

  new_agent(pid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/agent`, data);
  }

  get_agent(pid: string, aid: string) {
    return this.base.get(`/admin/plug/${pid}/agent/${aid}`);
  }

  update_agent(pid: string, aid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/agent/${aid}`, data);
  }

  delete_agent(pid: string, aid: string) {
    return this.base.delete(`/admin/plug/${pid}/agent/${aid}`);
  }

  // link

  list_agent_link(pid: string, aid: string) {
    return this.base.get(`/admin/plug/${pid}/agent/${aid}/link`);
  }

  new_agent_link(pid: string, aid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/agent/${aid}/link`, data);
  }

  update_agent_link(pid: string, aid: string, lid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/agent/${aid}/link/${lid}`, data);
  }

  get_agent_link(pid: string, aid: string, lid: string) {
    return this.base.get(`/admin/plug/${pid}/agent/${aid}/link/${lid}`);
  }
  delete_agent_link(pid: string, aid: string, lid: string) {
    return this.base.delete(`/admin/plug/${pid}/agent/${aid}/link/${lid}`);
  }

  // extension

  list_agent_ext(pid: string, aid: string) {
    return this.base.get(`/admin/plug/${pid}/agent/${aid}/extension`);
  }

  new_agent_ext(pid: string, aid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/agent/${aid}/extension`, data);
  }

  update_agent_ext(pid: string, aid: string, eid: string, data: any) {
    return this.base.post(
      `/admin/plug/${pid}/agent/${aid}/extension/${eid}`,
      data
    );
  }

  get_agent_ext(pid: string, aid: string, eid: string) {
    return this.base.get(`/admin/plug/${pid}/agent/${aid}/extension/${eid}`);
  }

  delete_agent_ext(pid: string, aid: string, eid: string) {
    return this.base.delete(`/admin/plug/${pid}/agent/${aid}/extension/${eid}`);
  }

  // resource

  list_agent_resource(pid: string, aid: string) {
    return this.base.get(`/admin/plug/${pid}/agent/${aid}/resource`);
  }
  new_agent_resource(pid: string, aid: string, data: any) {
    return this.base.post(`/admin/plug/${pid}/agent/${aid}/resource`, data);
  }

  update_agent_resource(pid: string, aid: string, rid: string, data: any) {
    return this.base.post(
      `/admin/plug/${pid}/agent/${aid}/resource/${rid}`,
      data
    );
  }
  get_agent_resource(pid: string, aid: string, rid: string) {
    return this.base.get(`/admin/plug/${pid}/agent/${aid}/resource/${rid}`);
  }

  delete_agent_resource(pid: string, aid: string, rid: string) {
    return this.base.delete(`/admin/plug/${pid}/agent/${aid}/resource/${rid}`);
  }
}


