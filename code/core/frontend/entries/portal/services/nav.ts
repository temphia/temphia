export class Navigator {
  constructor() {}

  start() {
    location.hash = "#/";
  }

  launch_app(name: string) {
    location.hash = `#/launch/${name}`;
  }

  data_loader() {
    location.hash = `#/data/`;
  }

  data_groups(source: string) {
    location.hash = `#/data/${source}`;
  }

  data_table(source: string, dgroup: string) {
    location.hash = `#/data/${source}/${dgroup}`;
  }

  cab_loader(source: string) {
    location.hash = `#/cabinet/${source}`;
  }

  cab_folder(source: string, folder: string) {
    location.hash = `#/data/${source}/${folder}`;
  }

  cab_file(source: string, folder: string, file: string) {
    location.hash = `#/data/${source}/${folder}/${file}`;
  }

  repo_loader() {
    location.hash = `#/repo/`;
  }

  repo_source(source: string) {
    location.hash = `#/repo/${source}`;
  }

  repo_item(source: string, group: string, item: string) {
    location.hash = `#/repo/${source}/${group}/${item}`;
  }

  self_profile() {
    location.hash = `#/profile/self`;
  }
  user_profile(user: string) {
    location.hash = `#/profile/user/${user}`;
  }

  play() {
    location.hash = `#/play`;
  }

  // ADMIN

  admin_bprints() {
    location.hash = `#/admin/bprint/`;
  }

  admin_bprint(bid: string) {
    location.hash = `#/admin/bprint/${bid}`;
  }

  admin_bprint_editor(bid: string) {
    location.hash = `#/admin/bprint/${bid}/editor`;
  }

  admin_bprint_instancer(bid: string) {
    location.hash = `#/admin/bprint/${bid}/instancer`;
  }

  admin_repos() {
    location.hash = `#/admin/repo/`;
  }

  admin_repo_edit(rid: string) {
    location.hash = `#/admin/repo/${rid}/edit`;
  }

  admin_repo_new() {
    location.hash = `#/admin/repo/new`;
  }

  admin_data_loader() {
    location.hash = `#/admin/data/`;
  }

  admin_data_groups() {
    location.hash = `#/admin/data/group`;
  }

  admin_data_group(group: string) {
    location.hash = `#/admin/data/group/${group}`;
  }

  admin_data_tables(group: string) {
    location.hash = `#/admin/data/table/${group}`;
  }

  admin_data_table(group: string, table: string) {
    location.hash = `#/admin/data/table/${group}/${table}`;
  }

  admin_data_columns(group: string, table: string) {
    location.hash = `#/admin/data/column/${group}/${table}`;
  }

  admin_data_column(group: string, table: string, column: string) {
    location.hash = `#/admin/data/column/${group}/${table}/${column}`;
  }

  admin_data_hooks(group: string, table: string) {
    location.hash = `#/admin/data/hook/${group}/${table}`;
  }

  admin_data_hook(group: string, table: string, id: string) {
    location.hash = `#/admin/data/hook/${group}/${table}/${id}`;
  }

  admin_data_views(group: string, table: string) {
    location.hash = `#/admin/data/view/${group}/${table}`;
  }

  admin_data_view(group: string, table: string, id: string) {
    location.hash = `#/admin/data/view/${group}/${table}/${id}`;
  }

  admin_resources() {
    location.hash = `#/admin/resource/`;
  }

  admin_resource_edit(rid: string) {
    location.hash = `#/admin/resource/${rid}/edit`;
  }

  admin_resource_new() {
    location.hash = `#/admin/resource/new`;
  }

  // apps

  admin_target_apps() {
    location.hash = `#/admin/target/app/`;
  }
  admin_target_app_edit(id: number) {
    location.hash = `#/admin/target/app/${id}/edit`;
  }

  admin_target_app_new() {
    location.hash = `#/admin/target/app/new`;
  }

  // hooks

  admin_target_hooks() {
    location.hash = `#/admin/target/hook/`;
  }
  admin_target_hook_edit(id: number) {
    location.hash = `#/admin/target/hook/${id}/edit`;
  }
  admin_target_hook_new() {
    location.hash = `#/admin/target/hook/new`;
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
