import { goto } from '$app/navigation';

export class Navigator {
  options: any;
  constructor() {
    this.options = null;
  }

  set = (new_url: string, opts?: any) => {
    this.options = opts;
    goto(`/z/pages/portal/${new_url}`)
  };

  start = () => {
    this.set("");
  };



  launch_target(target: string, opts?: { name?: string; target_type: string }) {

    const __opts = opts || {};
    const __name = __opts["name"] ? window.btoa(__opts["name"]) : ""


    this.set(`launch?target=${target}&name=${__name}`,opts);
  }

  launcher() {
    this.set(`launch/`);
  }

  notifications() {
    this.set(`notification/`);
  }

  // data

  data_page = (source: string) => {
    this.set(`data?dsource=${source}`);
  };

  data_group_page = (source: string, dgroup: string) => {
    this.set(`data/table?dsource=${source}&dgroup=${dgroup}`);
  }

  data_table_render_page = (source: string, dgroup: string, table: string, layout = "") => {
    if (layout === "vcard") {
      this.set(`data/table/render?dsource=${source}&dgroup=${dgroup}&dtable=${table}`);
    } else {
      this.set(`data/table/render?dsource=${source}&dgroup=${dgroup}&dtable=${table}`);
    }
  }

  data_sheets_page(source: string, dgroup: string) {
    this.set(`data/sheet/loader?dsource=${source}&dgroup=${dgroup}`);
  }


  data_sheet_render_page(source: string, dgroup: string, sheetid?: string, opts?: any) {
    const sheetparam = sheetid ? `&sheetid=${sheetid}` : ""

    this.set(`data/sheet?dsource=${source}&dgroup=${dgroup}${sheetparam}`, opts);
  }


  data_sheets_new() {
    this.set(`data/sheet/new`);
  }

  // cab

  cab_loader = () => {
    this.set(`cabinet/listings/`);
  };

  cab_uploader = (source: string, folder: string) => {
    this.set(`cabinet/upload?source=${source}&folder=${folder}`);
  };

  cab_folders = (source: string) => {
    this.set(`cabinet/listings?source=${source}`);
  };

  cab_folder = (source: string, folder: string) => {
    this.set(`cabinet/listings/folder?source=${source}&folder=${folder}`);
  };

  cab_file = (source: string, folder: string, file: string) => {
    this.set(`cabinet/listings/file?source=${source}&folder=${folder}&file=${file}`);
  };

  cab_text_file = (source: string, folder: string, file: string) => {
    this.set(`cabinet/listings/file/text?source=${source}&folder=${folder}&file=${file}`);
  };

  cab_image_file = (source: string, folder: string, file: string) => {
    this.set(`cabinet/listings/file/image?source=${source}&folder=${folder}&file=${file}`);
  };


  // repo

  repo_loader = () => {
    this.set(`repo`);
  };

  repo_source = (source: string) => {
    this.set(`repo/source?source=${source}`);
  };

  repo_item = (source: string, group: string, item: string) => {
    const p = new URLSearchParams({ source, group, item })

    this.set(`repo/item?${p.toString()}`);
  };

  self_profile = () => {
    this.set(`profile/self`);
  };

  user_profile = (user: string) => {
    this.set(`profile/user?id=${user}`);
  };

  self_devices = () => {
    this.set(`profile/device`);
  };

  self_device_new = () => {
    this.set(`profile/device/new`);
  };




  // ADMIN

  admin_bprints = () => {
    this.set(`admin/bprint`);
  };

  admin_bprint = (bid: string) => {
    this.set(`admin/bprint/edit?bid=${bid}`);
  };

  admin_bprint_export_zip = (bid: string) => {
    this.set(`admin/bprint/zipit?bid=${bid}`);
  };

  admin_bprint_new = () => {
    this.set(`admin/bprint/new`);
  };

  admin_bprint_new_zip = () => {
    this.set(`admin/bprint/zip`);
  };


  admin_bprint_instancer = (bid: string, opts?: any) => {
    this.set(`admin/bprint/instancer?bid=${bid}`, opts);
  };



  // plugs

  admin_plugs = () => {
    this.set(`admin/plug/`);
  };

  admin_plug_new = () => {
    this.set(`admin/plug/new`);
  };

  admin_plug_edit = (pid: string) => {
    this.set(`admin/plug/edit?pid=${pid}`);
  };

  // states

  admin_plug_states = (pid: string) => {
    this.set(`admin/plug/states?pid=${pid}`);
  };

  admin_plug_state_new = (pid: string) => {
    this.set(`admin/plug/states/new?pid=${pid}`);
  };

  admin_plug_state_edit = (pid: string, skey: string) => {
    this.set(`admin/plug/states/edit?pid=${pid}&skey=${skey}`);
  };

  // agents

  admin_agents = (pid: string) => {
    this.set(`admin/plug/agent?pid=${pid}`);
  };

  admin_agent_new = (pid: string) => {
    this.set(`admin/plug/agent/new?pid=${pid}`);
  };

  admin_agent_edit = (pid: string, aid: string) => {
    this.set(`admin/plug/agent/edit?pid=${pid}&aid=${aid}`);
  };

  admin_plug_dev_execute = (pid: string, aid: string) => {
    this.set(`admin/plug/agent/execute?pid=${pid}&aid=${aid}`);
  };

  // agent link

  admin_agent_links = (pid: string, aid: string) => {
    this.set(`admin/plug/agent/link?pid=${pid}&aid=${aid}`);
  };

  admin_agent_link_new = (pid: string, aid: string, opts?: any) => {
    this.set(`admin/plug/agent/link/new?pid=${pid}&aid=${aid}`, opts);
  };

  admin_agent_link_edit = (pid: string, aid: string, lid: string) => {
    this.set(`admin/plug/agent/link/edit?pid=${pid}&aid=${aid}&lid=${lid}`);
  };

  // agent ext

  admin_agent_ext = (pid: string, aid: string) => {
    this.set(`admin/plug/agent/ext?pid=${pid}&aid=${aid}`);
  };

  admin_agent_ext_new = (pid: string, aid: string) => {
    this.set(`admin/plug/agent/ext/new?pid=${pid}&aid=${aid}`);
  };

  admin_agent_ext_edit = (pid: string, aid: string, eid: string) => {
    this.set(`admin/plug/agent/ext/edit?pid=${pid}&aid=${aid}&eid=${eid}`);
  };

  admin_agent_res = (pid: string, aid: string) => {
    this.set(`admin/plug/agent/res?pid=${pid}&aid=${aid}`);
  };


  admin_agent_res_new = (pid: string, aid: string, opts?: any) => {
    this.set(`admin/plug/agent/res/new?pid=${pid}&aid=${aid}`, opts);
  };

  admin_agent_res_edit = (pid: string, aid: string, rid: string) => {
    this.set(`admin/plug/agent/res/edit?pid=${pid}&aid=${aid}&rid=${rid}`);
  };

  // dev

  admin_plug_dev_flowmap = (pid: string) => {
    this.set(`#/admin/plug/dev/flowmap?pid=${pid}`);
  };


  admin_plug_dev_live_shell = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/liveshell/${aid}`);
  };

  admin_plug_dev_ipc_shell = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/ipcshell/${aid}`);
  };

  admin_plug_dev_docs = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/docs/${aid}`);
  };

  admin_plug_dev_logs = (pid: string, aid: string) => {
    this.set(`#/admin/plug/${pid}/dev/logs/${aid}`);
  };



  // repo

  admin_repos = () => {
    this.set(`admin/repo/`);
  };

  admin_repo_edit = (rid: string) => {
    this.set(`admin/repo/edit/${rid}`);
  };

  admin_repo_new = () => {
    this.set(`admin/repo/new`);
  };

  // data

  admin_data_loader = () => {
    this.set(`admin/data/`);
  };

  admin_data_groups = (source: string) => {
    this.set(`admin/data/dgroup?source=${source}`);
  };

  admin_data_group = (source: string, group: string) => {
    this.set(`admin/data/dgroup/edit?source=${source}&group=${group}`);
  };

  admin_data_tables = (source: string, group: string) => {
    this.set(`admin/data/dtable?source=${source}&group=${group}`);
  };

  admin_data_table = (source: string, group: string, table: string) => {
    this.set(`admin/data/dtable/edit?source=${source}&group=${group}&table=${table}`);
  };

  admin_data_columns = (source: string, group: string, table: string) => {
    this.set(`admin/data/dcolumn?source=${source}&group=${group}&table=${table}`);
  };

  admin_data_column = (
    source: string,
    group: string,
    table: string,
    column: string
  ) => {
    this.set(`admin/data/dcolumn/edit?source=${source}&group=${group}&table=${table}&column=${column}`);
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
    this.set(`admin/resource/`);
  };

  admin_resource_edit = (rid: string) => {
    this.set(`admin/resource/edit/${rid}`);
  };

  admin_resource_new = (opts?: any) => {
    this.set(`admin/resource/new`, opts);
  };

  // apps

  admin_target_apps = () => {
    this.set(`admin/target/app`);
  };
  admin_target_app_edit = (ttype: string, id: number) => {
    this.set(`admin/target/app/edit?ttype=${ttype}&id=${id}`);
  };

  admin_target_app_new = (opts?: any) => {
    this.set(`admin/target/app/new`, opts);
  };

  // hooks

  admin_target_hooks = () => {
    this.set(`admin/target/hook/`);
  };
  admin_target_hook_edit = (ttype: string, id: number) => {
    this.set(`admin/target/hook/edit?ttype=${ttype}&id=${id}`);
  };
  admin_target_hook_new = () => {
    this.set(`admin/target/hook/new`);
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

  admin_ugroups() {
    this.set(`admin/ugroup/`);
  }
  admin_ugroup_new() {
    this.set(`admin/ugroup/new`);
  }
  admin_ugroup_edit(ugroup: string) {
    this.set(`admin/ugroup/edit/${ugroup}`);
  }

  admin_ugroup_users(ugroup: string) {
    this.set(`admin/ugroup/subusers/${ugroup}`);
  }

  ////

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
    this.set(`admin/tenant/system/kv`);
  }

  admin_tenant_system_events() {
    this.set(`admin/tenant/system/event`);
  }

  admin_lens_logs(opts?: any) {
    this.set(`admin/lens/logs`, opts);
  }

  admin_lens_watcher(opts?: any) {
    this.set(`admin/lens/watcher`, opts);
  }
}
