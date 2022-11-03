import { routes } from "svelte-hash-router";

export class Navigator {
  options: any;
  constructor() {
    this.options = null;
  }

  set(new_url: string, opts?: any) {
    this.options = opts;
    location.hash = new_url;
  }

  start() {
    this.set("#/");
  }

  launch_app(name: string) {
    this.set(`#/launch/${name}`);
  }

  data_loader() {
    this.set(`#/data/`);
  }

  data_groups(source: string) {
    this.set(`#/data/${source}`);
  }

  data_table(source: string, dgroup: string) {
    this.set(`#/data/${source}/${dgroup}`);
  }

  cab_loader(source: string) {
    this.set(`#/cabinet/${source}`);
  }

  cab_folder(source: string, folder: string) {
    this.set(`#/data/${source}/${folder}`);
  }

  cab_file(source: string, folder: string, file: string) {
    this.set(`#/data/${source}/${folder}/${file}`);
  }

  repo_loader() {
    this.set(`#/repo/`);
  }

  repo_source(source: string) {
    this.set(`#/repo/${source}`);
  }

  repo_item(source: string, group: string, item: string) {
    this.set(`#/repo/${source}/${group}/${item}`);
  }

  self_profile() {
    this.set(`#/profile/self`);
  }
  user_profile(user: string) {
    this.set(`#/profile/user/${user}`);
  }

  play() {
    this.set(`#/play`);
  }

  // ADMIN

  admin_bprints() {
    this.set(`#/admin/bprint/`);
  }

  admin_bprint(bid: string) {
    this.set(`#/admin/bprint/${bid}`);
  }

  admin_bprint_editor(bid: string) {
    this.set(`#/admin/bprint/${bid}/editor`);
  }

  admin_bprint_data_instancer(bid: string, file: string, opts?: any) {
    this.set(`#/admin/bprint/${bid}/instance/data/${file}`, opts);
  }

  admin_bprint_plug_instancer(bid: string, file: string, opts?: any) {
    this.set(`#/admin/bprint/${bid}/instance/plug/${file}`, opts);
  }

  admin_repos() {
    this.set(`#/admin/repo/`);
  }

  admin_repo_edit(rid: string) {
    this.set(`#/admin/repo/${rid}/edit`);
  }

  admin_repo_new() {
    this.set(`#/admin/repo/new`);
  }

  admin_data_loader() {
    this.set(`#/admin/data/`);
  }

  admin_data_groups() {
    this.set(`#/admin/data/group`);
  }

  admin_data_group(group: string) {
    this.set(`#/admin/data/group/${group}`);
  }

  admin_data_tables(group: string) {
    this.set(`#/admin/data/table/${group}`);
  }

  admin_data_table(group: string, table: string) {
    this.set(`#/admin/data/table/${group}/${table}`);
  }

  admin_data_columns(group: string, table: string) {
    this.set(`#/admin/data/column/${group}/${table}`);
  }

  admin_data_column(group: string, table: string, column: string) {
    this.set(`#/admin/data/column/${group}/${table}/${column}`);
  }

  admin_data_hooks(group: string, table: string) {
    this.set(`#/admin/data/hook/${group}/${table}`);
  }

  admin_data_hook(group: string, table: string, id: string) {
    this.set(`#/admin/data/hook/${group}/${table}/${id}`);
  }

  admin_data_views(group: string, table: string) {
    this.set(`#/admin/data/view/${group}/${table}`);
  }

  admin_data_view(group: string, table: string, id: string) {
    this.set(`#/admin/data/view/${group}/${table}/${id}`);
  }

  admin_resources() {
    this.set(`#/admin/resource/`);
  }

  admin_resource_edit(rid: string) {
    this.set(`#/admin/resource/${rid}/edit`);
  }

  admin_resource_new() {
    this.set(`#/admin/resource/new`);
  }

  // apps

  admin_target_apps() {
    this.set(`#/admin/target/app/`);
  }
  admin_target_app_edit(id: number) {
    this.set(`#/admin/target/app/${id}/edit`);
  }

  admin_target_app_new() {
    this.set(`#/admin/target/app/new`);
  }

  // hooks

  admin_target_hooks() {
    this.set(`#/admin/target/hook/`);
  }
  admin_target_hook_edit(id: number) {
    this.set(`#/admin/target/hook/${id}/edit`);
  }
  admin_target_hook_new() {
    this.set(`#/admin/target/hook/new`);
  }

  /*
  

  "/plug": {
    "/": AdminPlugs,
    "/:pid": AdminPlug,

  },

  "/user": {
    "/": AdminUsers,
    "/:userid": AdminUser,
  },

  "/ugroup": {
    "/": AdminUgroups,
    "/:ugroup": AdminUgroup,
  },
  "/lens": {
    $$component: AdminLens,
    "/app": AdminLensApp,
    "/engine": AdminLensEngine,
    "/site": AdminLensSite,
  },
  "/tenant": {
    "/": AdminTenant,
    "/edit": AdminTenantEdit,
    "/domain": AdminTenantDomains,
    "/domain/:did": AdminTenantDomain,
  },
  
  */
}
