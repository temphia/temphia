import { routes } from "svelte-hash-router";

export class Navigator {
  options: any;
  constructor() {
    this.options = null;
  }

  set = (new_url: string, opts?: any) => {
    this.options = opts;
    location.hash = new_url;
  };

  start = () => {
    this.set("#/");
  };

  launch_target(target: string, opts?: { name?: string; target_type: string }) {
    this.set(
      `#/launch/${target}${
        (opts || {}).name ? "/" + window.btoa(opts.name) : ""
      }`,
      opts
    );
  }

  launcher() {
    this.set(`#/launch/`);
  }

  // data

  data_loader = () => {
    this.set(`#/data/`);
  };

  data_groups = (source: string) => {
    this.set(`#/data/${source}`);
  };

  data_group(source: string, dgroup: string) {
    this.set(`#/data/${source}/${dgroup}`);
  }

  data_render_custom_loader(source: string, dgroup: string) {
    this.set(`#/data/${source}/${dgroup}/custom`);
  }

  data_render_table(source: string, dgroup: string, table: string, layout = "") {
    this.set(`#/data/${source}/${dgroup}/table/${table}${layout}`);
  }

  data_render_table_loader(source: string, dgroup: string) {
    this.set(`#/data/${source}/${dgroup}/table`);
  }

  data_render_sheet(source: string, dgroup: string, sheetid: string) {
    this.set(`#/data/${source}/${dgroup}/sheet/${sheetid}`);
  }

  data_render_sheet_loader(source: string, dgroup: string) {
    this.set(`#/data/${source}/${dgroup}/sheet`);
  }

  // cab

  cab_loader = () => {
    this.set(`#/cabinet/`);
  };

  cab_folders = (source: string) => {
    this.set(`#/cabinet/${source}`);
  };

  cab_folder = (source: string, folder: string) => {
    this.set(`#/cabinet/${source}/${folder}`);
  };

  cab_file = (source: string, folder: string, file: string) => {
    this.set(`#/cabinet/${source}/${folder}/${file}`);
  };

  // repo

  repo_loader = () => {
    this.set(`#/repo/`);
  };

  repo_source = (source: string) => {
    this.set(`#/repo/${source}`);
  };

  repo_item = (source: string, group: string, item: string) => {
    this.set(`#/repo/${source}/${group}/${item}`);
  };

  self_profile = () => {
    this.set(`#/profile/self`);
  };
  user_profile = (user: string) => {
    this.set(`#/profile/user/${user}`);
  };

  self_devices = () => {
    this.set(`#/profile/device`);
  };

  self_device_new = () => {
    this.set(`#/profile/device/new`);
  };

  play = () => {
    this.set(`#/play`);
  };

  // ADMIN

  admin_bprints = () => {
    this.set(`#/admin/bprint/`);
  };

  admin_bprint = (bid: string) => {
    this.set(`#/admin/bprint/${bid}/edit`);
  };

  admin_bprint_export_zip = (bid: string) => {
    this.set(`#/admin/bprint/${bid}/zipit`);
  };

  admin_bprint_new = () => {
    this.set(`#/admin/bprint/new`);
  };

  admin_bprint_new_zip = () => {
    this.set(`#/admin/bprint/new/zip`);
  };

  admin_bprint_files = (bid: string) => {
    this.set(`#/admin/bprint/${bid}/file`);
  };

  admin_bprint_file = (bid: string, file: string) => {
    this.set(`#/admin/bprint/${bid}/file/${file}`);
  };

  admin_bprint_data_instancer = (bid: string, file: string, opts?: any) => {
    this.set(`#/admin/bprint/${bid}/instance/data/${file}`, opts);
  };

  admin_bprint_plug_instancer = (bid: string, file: string, opts?: any) => {
    this.set(`#/admin/bprint/${bid}/instance/plug/${file}`, opts);
  };

  admin_bprint_auto_instancer = (bid: string, opts?: any) => {
    this.set(`#/admin/bprint/${bid}/instance/automatic`, opts);
  };

  admin_bprint_editor = (bid: string, etype: string, file: string) => {
    this.set(`#/admin/bprint/${bid}/editor/${etype}/${file}`);
  };

  // plugs

  admin_plugs = () => {
    this.set(`#/admin/plug/`);
  };

  admin_plug_new = () => {
    this.set(`#/admin/plug/new`);
  };

  admin_plug_edit = (bid: string) => {
    this.set(`#/admin/plug/${bid}/edit`);
  };

  // agents

  admin_plug_states = (pid: string) => {
    this.set(`#/admin/plug/${pid}/state/`);
  };

  admin_plug_state_new = (pid: string) => {
    this.set(`#/admin/plug/${pid}/state/new`);
  };

  admin_plug_state_edit = (pid: string, skey) => {
    this.set(`#/admin/plug/${pid}/state/${skey}/edit`);
  };

  admin_agents = (pid: string) => {
    this.set(`#/admin/plug/${pid}/agent/`);
  };

  admin_agent_new = (pid: string) => {
    this.set(`#/admin/plug/${pid}/agent/new`);
  };

  admin_agent_edit = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/edit`);
  };

  // agent link

  admin_agent_links = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/link/`);
  };

  admin_agent_link_new = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/link/new`);
  };

  admin_agent_link_edit = (pid: string, aid: string, lid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/link/${lid}/edit`);
  };

  // agent ext

  admin_agent_ext = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/ext/`);
  };

  admin_agent_ext_new = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/ext/new`);
  };

  admin_agent_ext_edit = (pid: string, aid: string, eid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/ext/${eid}/edit`);
  };

  admin_agent_res = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/res/`);
  };

  admin_agent_res_new = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/res/new`);
  };

  admin_agent_res_edit = (pid: string, aid: string, rid: string) => {
    this.set(`#/admin/plug/${pid}/agent/${aid}/res/${rid}/edit`);
  };

  // dev

  admin_plug_dev_flowmap = (pid: string) => {
    this.set(`#/admin/plug/${pid}/dev/flowmap`);
  };
  admin_plug_dev_execute = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/execute/${aid}`);
  };
  admin_plug_dev_shell = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/shell/${aid}`);
  };
  admin_plug_dev_docs = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/docs/${aid}`);
  };

  admin_plug_dev_logs = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/logs/${aid}`);
  };

  // repo

  admin_repos = () => {
    this.set(`#/admin/repo/`);
  };

  admin_repo_edit = (rid: string) => {
    this.set(`#/admin/repo/${rid}/edit`);
  };

  admin_repo_new = () => {
    this.set(`#/admin/repo/new`);
  };

  // data

  admin_data_loader = () => {
    this.set(`#/admin/data/`);
  };

  admin_data_groups = (source: string) => {
    this.set(`#/admin/data/${source}/group`);
  };

  admin_data_group = (source: string, group: string) => {
    this.set(`#/admin/data/${source}/group/${group}`);
  };

  admin_data_tables = (source: string, group: string) => {
    this.set(`#/admin/data/${source}/table/${group}`);
  };

  admin_data_table = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/table/${group}/${table}`);
  };

  admin_data_columns = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/column/${group}/${table}`);
  };

  admin_data_column = (
    source: string,
    group: string,
    table: string,
    column: string
  ) => {
    this.set(`#/admin/data/${source}/column/${group}/${table}/${column}`);
  };

  // hooks

  admin_data_hooks = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/hook/${group}/${table}`);
  };

  admin_data_hook = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/hook/${group}/${table}/new`);
  };

  admin_data_apps = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/app/${group}/${table}`);
  };

  admin_data_app = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/app/${group}/${table}/new`);
  };
  // views

  admin_data_views = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/view/${group}/${table}`);
  };

  admin_data_view_new = (source: string, group: string, table: string) => {
    this.set(`#/admin/data/${source}/view/${group}/${table}/new`);
  };

  admin_data_view_edit = (
    source: string,
    group: string,
    table: string,
    id: string
  ) => {
    this.set(`#/admin/data/${source}/view/${group}/${table}/${id}/edit`);
  };

  admin_data_activity(source: string, group: string, table: string) {
    this.set(`#/admin/data/${source}/activity/${group}/${table}`);
  }

  admin_data_seed(source: string, group: string, table: string) {
    this.set(`#/admin/data/${source}/seed/${group}/${table}`);
  }

  admin_data_query(source: string, group: string) {
    this.set(`#/admin/data/${source}/query/${group}`);
  }

  // resources

  admin_resources = () => {
    this.set(`#/admin/resource/`);
  };

  admin_resource_edit = (rid: string) => {
    this.set(`#/admin/resource/${rid}/edit`);
  };

  admin_resource_new = (plug_id?: string) => {
    this.set(`#/admin/resource/new`, plug_id ? { plug_id } : undefined);
  };

  admin_resource_datagroup_new = () => {
    this.set(`#/admin/resource/rtype/data_group/new`);
  };
  admin_resource_datagroup_edit = (rid: string) => {
    this.set(`#/admin/resource/rtype/data_group/${rid}/edit`);
  };

  admin_resource_room_new = () => {
    this.set(`#/admin/resource/rtype/room/new`);
  };

  admin_resource_room_edit = (rid: string) => {
    this.set(`#/admin/resource/rtype/room/${rid}/edit`);
  };

  admin_resource_folder_new = () => {
    this.set(`#/admin/resource/rtype/folder/new`);
  };

  admin_resource_folder_edit = (rid: string) => {
    this.set(`#/admin/resource/rtype/folder/${rid}/edit`);
  };

  // apps

  admin_target_apps = () => {
    this.set(`#/admin/target/app/`);
  };
  admin_target_app_edit = (ttype: string, id: number) => {
    this.set(`#/admin/target/app/${ttype}/${id}/edit`);
  };

  admin_target_app_new = (opts?: any) => {
    this.set(`#/admin/target/app/new`, opts);
  };

  // hooks

  admin_target_hooks = () => {
    this.set(`#/admin/target/hook/`);
  };
  admin_target_hook_edit = (ttype: string, id: number) => {
    this.set(`#/admin/target/hook/${ttype}/${id}/edit`);
  };
  admin_target_hook_new = () => {
    this.set(`#/admin/target/hook/new`);
  };

  // user

  admin_users() {
    this.set(`#/admin/user/`);
  }

  admin_user_edit(userid: string) {
    this.set(`#/admin/user/edit/${userid}`);
  }

  admin_user_new() {
    this.set(`#/admin/user/new`);
  }

  admin_perms() {
    this.set(`#/admin/user/perm`);
  }

  admin_perm_new() {
    this.set(`#/admin/user/perm/new`);
  }

  admin_roles() {
    this.set(`#/admin/user/role`);
  }

  admin_role_new() {
    this.set(`#/admin/user/role/new`);
  }

  admin_ugroups() {
    this.set(`#/admin/ugroup/`);
  }
  admin_ugroup_new() {
    this.set(`#/admin/ugroup/new`);
  }
  admin_ugroup_edit(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/edit`);
  }

  admin_ugroup_users(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/user/`);
  }

  admin_ugroup_user_edit(ugroup: string, user_id: string) {
    this.set(`#/admin/ugroup/${ugroup}/user/${user_id}/edit`);
  }

  admin_ugroup_user_new(ugroup: string, opts?: any) {
    this.set(`#/admin/ugroup/${ugroup}/user/new`, opts);
  }

  admin_ugroup_auths(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/auth`);
  }

  admin_ugroup_auth_new(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/auth/new`);
  }

  admin_ugroup_auth_edit(ugroup: string, id: string) {
    this.set(`#/admin/ugroup/${ugroup}/auth/edit/${id}`);
  }

  admin_ugroup_apps(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/app`);
  }

  admin_ugroup_app_new(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/app/new`);
  }

  admin_ugroup_app_edit(ugroup: string, id: string) {
    this.set(`#/admin/ugroup/${ugroup}/app/edit/${id}`);
  }

  admin_ugroup_datas(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/data`);
  }

  admin_ugroup_data_new(ugroup: string) {
    this.set(`#/admin/ugroup/${ugroup}/data/new`);
  }

  admin_ugroup_data_edit(ugroup: string, id: string) {
    this.set(`#/admin/ugroup/${ugroup}/data/edit/${id}`);
  }

  admin_tenant() {
    this.set(`#/admin/tenant/`);
  }

  admin_tenant_edit() {
    this.set(`#/admin/tenant/edit`);
  }

  admin_tenant_domain_edit(did: string) {
    this.set(`#/admin/tenant/domain/${did}/edit`);
  }

  admin_tenant_domain_new() {
    this.set(`#/admin/tenant/domain/new`);
  }

  admin_tenant_domain_adapter_editor(did: string) {
    this.set(`#/admin/tenant/domain/${did}/adapter_editor`);
  }

  admin_tenant_system_kvs() {
    this.set(`#/admin/tenant/system/kv`);
  }

  admin_tenant_system_events() {
    this.set(`#/admin/tenant/system/event`);
  }

  admin_lens(idx: string) {
    this.set(`#/admin/lens/${idx}`);
  }
}
