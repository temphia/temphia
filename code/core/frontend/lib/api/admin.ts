import { ApiBase } from "./base";

export class TenantAPI extends ApiBase {
  constructor(url: string, user_token: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["admin", "tenant"],
    });
  }
  async update_tenant(data: any) {
    return this.post("/tenant", data);
  }

  async get_tenant() {
    return this.get("/tenant");
  }

  async list_tenant_domain() {
    return this.get("/tenant/domain");
  }

  async add_tenant_domain(data: any) {
    return this.post("/tenant/domain/", data);
  }

  async get_tenant_domain(id: number) {
    return this.get(`/tenant/domain/${id}`);
  }

  async update_tenant_domain(id: number, data: any) {
    return this.post(`/tenant/domain/${id}`, data);
  }

  async remove_tenant_domain(id: number) {
    return this.delete(`/tenant/domain/${id}`);
  }

  // widget

  async list_domain_widget(did: number) {
    return this.get(`/tenant/domain/${did}/widget`);
  }
  async add_domain_widget(did: number, data: any) {
    return this.post(`/tenant/domain/${did}/widget`, data);
  }
  async get_domain_widget(did: number, wid: number) {
    return this.get(`/tenant/domain/${did}/widget${wid}`);
  }
  async update_domain_widget(did: number, wid: number, data: any) {
    return this.get(`/tenant/domain/${did}/widget${wid}`, data);
  }
  async remove_domain_widget(did: number, wid: number) {
    return this.delete(`/tenant/domain/${did}/widget${wid}`);
  }

  async list_repo() {
    return this.get(`/tenant_repo`);
  }
  async new_repo(data: any) {
    return this.post(`/tenant_repo`, data);
  }

  async get_repo(rid: string) {
    return this.get(`/tenant_repo/${rid}`);
  }

  async update_repo(rid: string, data: any) {
    return this.post(`/tenant_repo/${rid}`, data);
  }

  async del_repo(rid: string) {
    return this.delete(`/tenant_repo/${rid}`);
  }

  async list_renderer() {
    return this.get(`/tenant/renderer`);
  }
}

export class BprintAPI extends ApiBase {
  constructor(url: string, user_token: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["admin", "bprint"],
    });
  }

  async bprint_list() {
    return this.get("/bprint");
  }
  async bprint_create(data: any) {
    return this.post("/bprint", data);
  }
  async bprint_get(id: string) {
    return this.get(`/bprint/${id}`);
  }
  async bprint_update(id: string, data: any) {
    return this.post(`/bprint/${id}`, data);
  }
  async bprint_remove(id: string) {
    return this.delete(`/bprint/${id}`);
  }
  async bprint_install(id: string, opts: any) {
    return this.post(`/bprint/${id}/install`, opts);
  }

  async bprint_instance(id: string, opts: any) {
    return this.post(`/bprint/${id}/instance`, opts);
  }


  async bprint_list_files(id: string) {
    return this.get(`/bprint/${id}/file`);
  }
  async bprint_get_file(id: string, file: string) {
    return this.get(`/bprint/${id}/file/${file}`);
  }
  async bprint_new_file(id: string, file: string, data: any) {
    return this.post(`/bprint/${id}/file/${file}`, data);
  }
  async bprint_update_file(id: string, file: string, data: any) {
    return this.patch(`/bprint/${id}/file/${file}`, data);
  }

  async bprint_del_file(id: string, file: string) {
    return this.delete(`/bprint/${id}/file/${file}`);
  }
  async bprint_import(data: any) {
    return this.post(`/import_bprint`, data);
  }

  async repo_sources() {
    return this.get(`/repo`);
  }

  async repo_list(source: string) {
    return this.get(`/repo/${source}`);
  }
  async repo_get(source: string, group: string, slug: string) {
    return this.get(`/repo/${source}/${group}/${slug}`);
  }
  async repo_get_file(source: string, slug: string, file: string) {
    return this.get(`/repo/${source}/${slug}/${file}`);
  }

  async issue_tkt(
    bid: string,
    all_plugs: boolean,
    encoded: boolean,
    pids?: string[]
  ) {
    return this.post("/dev_plug_issue_tkt", {
      plug_ids: pids,
      bprint_id: bid,
      all_plugs,
      encoded,
    });
  }

  async check_bprint(bid: string) {
    return this.get(`/check_slug/bprint/${bid}`);
  }

  async check_plug(pid: string) {
    return this.get(`/check_slug/plug/${pid}`);
  }

  async check_data_group(source: string, gid: string) {
    return this.get(`/check_slug/data_group/${source}/${gid}`);
  }

  async check_data_table(source: string, gid: string, tid: string) {
    return this.get(`/check_slug/data_table/${source}/${gid}/${tid}`);
  }

}

export class UserAPI extends ApiBase {
  constructor(url: string, user_token: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["admin", "user"],
    });
  }

  async list_users(group?: string) {
    return this.get(`/user${group ? `?user_group=` + group : ""}`);
  }
  async add_user(data: any) {
    return this.post(`/user`, data);
  }

  async get_user_by_id(id: string) {
    return this.get(`/user/${id}`);
  }
  async update_user(id: string, data: any) {
    return this.post(`/user/${id}`, data);
  }
  async remove_user(id: string) {
    return this.delete(`/user/${id}`);
  }

  async list_user_group() {
    return this.get(`/user_group`);
  }
  async add_user_group(data: any) {
    return this.post(`/user_group`, data);
  }
  async get_user_group(gid: string) {
    return this.get(`/user_group/${gid}`);
  }

  async update_user_group(gid: string, data: any) {
    return this.post(`/user_group/${gid}`, data);
  }
  async remove_user_group(gid: string) {
    return this.delete(`/user_group/${gid}`);
  }

  // auth

  async user_group_list_auth(gid: string) {
    return this.get(`/user_auth/${gid}`);
  }

  async user_group_add_auth(gid: string, data: any) {
    return this.post(`/user_auth/${gid}`, data);
  }

  async user_group_get_auth(gid: string, id: number) {
    return this.get(`/user_auth/${gid}/${id}`);
  }

  async user_group_update_auth(gid: string, id: number, data: any) {
    return this.post(`/user_auth/${gid}/${id}`, data);
  }
  async user_group_remove_auth(gid: string, id: number) {
    return this.delete(`/user_auth/${gid}/${id}`);
  }

  // hook

  async user_group_list_hook(gid: string) {
    return this.get(`/user_hook/${gid}`);
  }

  async user_group_add_hook(gid: string, data: any) {
    return this.post(`/user_hook/${gid}`, data);
  }

  async user_group_get_hook(gid: string, id: number) {
    return this.get(`/user_hook/${gid}/${id}`);
  }

  async user_group_update_hook(gid: string, id: number, data: any) {
    return this.post(`/user_hook/${gid}/${id}`, data);
  }

  async user_group_remove_hook(gid: string, id: number) {
    return this.get(`/user_hook/${gid}/${id}`);
  }

  // plug

  async user_group_list_plug(gid: string) {
    return this.get(`/user_plug/${gid}`);
  }

  async user_group_add_plug(gid: string, data: any) {
    return this.post(`/user_plug/${gid}`, data);
  }

  async user_group_get_plug(gid: string, id: number) {
    return this.get(`/user_plug/${gid}/${id}`);
  }

  async user_group_update_plug(gid: string, id: number, data: any) {
    return this.post(`/user_plug/${gid}/${id}`, data);
  }

  async user_group_remove_plug(gid: string, id: number) {
    return this.get(`/user_plug/${gid}/${id}`);
  }

  // data

  async user_group_list_data(gid: string) {
    return this.get(`/user_data/${gid}`);
  }

  async user_group_add_data(gid: string, data: any) {
    return this.post(`/user_data/${gid}`, data);
  }

  async user_group_get_data(gid: string, id: number) {
    return this.get(`/user_data/${gid}/${id}`);
  }

  async user_group_update_data(gid: string, id: number, data: any) {
    return this.post(`/user_data/${gid}/${id}`, data);
  }

  async user_group_remove_data(gid: string, id: number) {
    return this.get(`/user_data/${gid}/${id}`);
  }

  // fixme => user perm stuff
}

export class PlugAPI extends ApiBase {
  constructor(url: string, user_token: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["admin", "plug"],
    });
  }

  async list_plug() {
    return this.get(`/plug`);
  }

  async list_agent_bprint(bid: string) {
    return this.get(`/plug?bprint_id=${bid}`);
  }

  async new_plug(data: string) {
    return this.post(`/plug`, data);
  }

  async update_plug(id: string, data: any) {
    return this.post(`/plug/${id}`, data);
  }

  async get_plug(pid: string) {
    return this.get(`/plug/${pid}`);
  }
  async del_plug(pid: string) {
    return this.delete(`/plug/${pid}`);
  }

  async list_agent(pid: string) {
    return this.get(`/plug/${pid}/agent`);
  }

  async new_agent(pid: string, data: any) {
    return this.post(`/plug/${pid}/agent`, data);
  }

  async update_agent(pid: string, aid: string, data: any) {
    return this.post(`/plug/${pid}/agent/${aid}`, data);
  }

  async get_agent(pid: string, aid: string) {
    return this.get(`/plug/${pid}/agent/${aid}`);
  }
  async del_agent(pid: string, aid: string) {
    return this.delete(`/plug/${pid}/agent/${aid}`);
  }

  async agent_link_list(pid: string, aid: string) {
    return this.get(`/plug/${pid}/agent/${aid}/link`);
  }

  async agent_link_new(pid: string, aid: string, data: any) {
    return this.post(`/plug/${pid}/agent/${aid}/link`, data);
  }

  async agent_link_get(pid: string, aid: string, lid: number ) {
    return this.get(`/plug/${pid}/agent/${aid}/link/${lid}`);
  }

  async agent_link_update(pid: string, aid: string, lid: number, data: any) {
    return this.post(`/plug/${pid}/agent/${aid}/link/${lid}`, data);
  }

  async agent_link_del(pid: string, aid: string, lid: number) {
    return this.delete(`/plug/${pid}/agent/${aid}/link/${lid}`);
  }

  async agent_extension_list(pid: string, aid: string) {
    return this.get(`/plug/${pid}/agent/${aid}/extension`);
  }

  async agent_extension_new(pid: string, aid: string, data: any) {
    return this.post(`/plug/${pid}/agent/${aid}/extension`, data);
  }

  async agent_extension_get(pid: string, aid: string, eid: number ) {
    return this.get(`/plug/${pid}/agent/${aid}/extension/${eid}`);
  }

  async agent_extension_update(pid: string, aid: string, eid: number, data: any) {
    return this.post(`/plug/${pid}/agent/${aid}/extension/${eid}`, data);
  }

  async agent_extension_del(pid: string, aid: string, eid: number) {
    return this.delete(`/plug/${pid}/agent/${aid}/extension/${eid}`);
  }

  async agent_resource_list(pid: string, aid: string) {
    return this.get(`/plug/${pid}/agent/${aid}/resource`);
  }

  async agent_resource_new(pid: string, aid: string, data: any) {
    return this.post(`/plug/${pid}/agent/${aid}/resource`, data);
  }

  async agent_resource_get(pid: string, aid: string, slug: string ) {
    return this.get(`/plug/${pid}/agent/${aid}/resource/${slug}`);
  }

  async agent_resource_update(pid: string, aid: string, slug: string, data: any) {
    return this.post(`/plug/${pid}/agent/${aid}/resource/${slug}`, data);
  }

  async agent_resource_del(pid: string, aid: string, slug: string) {
    return this.delete(`/plug/${pid}/agent/${aid}/resource/${slug}`);
  }

  async launch_agent(plug: string, agent: string, data: any) {
    return this.post(`/engine/${plug}/${agent}/launcher/json`, data);
  }
}

export class CabinetAPI extends ApiBase {
  constructor(url: string, user_token: string, source: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["cabinet", source],
    });
  }
  async list_root() {
    return this.get(`/cabinet`);
  }
  async list_folder(folder: string) {
    return this.get(`/cabinet/${folder}`);
  }
  async new_folder(folder: string) {
    return this.post(`/cabinet/${folder}`);
  }
  async get_file(folder: string, file: string) {
    return this.get(`/cabinet/${folder}/file/${file}`);
  }
  async upload_file(folder: string, file: string, data) {
    return this.post(`/cabinet/${folder}/file/${file}`, data);
  }
  async delete_file(folder: string, file: string) {
    return this.delete(`/cabinet/${folder}/file/${file}`);
  }

  async get_folder_ticket(folder: string) {
    return this.post(`/cabinet/${folder}/ticket`);
  }
}

export class ResourceAPI extends ApiBase {
  constructor(url: string, user_token: string) {
    super({
      url: url,
      user_token: user_token,
      path: ["resource"],
    });
  }

  async agent_resources_list(plug_id: string, agent_id: string) {
    return this.post("/agent_resources", {
      plug_id,
      agent_id,
    });
  }

  async resource_list(plug_id?: string) {
    return this.get(`/resource?plug_id=${plug_id ? plug_id : ""}`);
  }

  async resource_create(data: any) {
    return this.post("/resource", data);
  }

  async resource_get(slug: string) {
    return this.get(`/resource/${slug}`);
  }

  async resource_update(slug: string, data: any) {
    return this.post(`/resource/${slug}`, data);
  }

  async resource_remove(slug: string) {
    return this.delete(`/resource/${slug}`);
  }
}
