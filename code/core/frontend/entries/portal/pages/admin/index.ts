import AdminBprint from "./bprint/bprint.svelte";
import AdminBprints from "./bprint/bprints.svelte";
import AdminBprintEditor from "./bprint/editor/editor.svelte";
import AdminBprintInstancer from "./bprint/instancer/instancer.svelte";

// repo
import AdminRepos from "./repo/repos.svelte";
import AdminRepoEdit from "./repo/repo_edit.svelte";
import AdminRepoNew from "./repo/repo_new.svelte";

// data
import AdminDataLoader from "./data/loader.svelte";
import AdminDataGroup from "./data/group/group.svelte";
import AdminDataGroups from "./data/group/groups.svelte";
import AdminDataTable from "./data/table/table.svelte";
import AdminDataTables from "./data/table/tables.svelte";
import AdminDataColumns from "./data/column/columns.svelte";
import AdminDataColumn from "./data/column/column.svelte";
import AdminDataHook from "./data/hook/hook.svelte";
import AdminDataHooks from "./data/hook/hooks.svelte";
import AdminDataView from "./data/view/view.svelte";
import AdminDataViews from "./data/view/views.svelte";
import Admin from "./admin.svelte";

// resource
import AdminResources from "./resource/resources.svelte";
import AdminResourceEdit from "./resource/resource_edit.svelte";
import AdminResourceNew from "./resource/resource_new.svelte";

// user
import AdminUsers from "./user/users.svelte";
import AdminUser from "./user/user.svelte";

// ugroup
import AdminUgroups from "./ugroup/ugroups.svelte";
import AdminUgroup from "./ugroup/ugroup.svelte";

// lens
import AdminLens from "./lens/lens.svelte";
import AdminLensApp from "./lens/app.svelte";
import AdminLensEngine from "./lens/engine.svelte";
import AdminLensSite from "./lens/site.svelte";

// tenant
import AdminTenant from "./tenant/tenant.svelte";
import AdminTenantEdit from "./tenant/tenant_edit.svelte";
import AdminTenantDomains from "./tenant/domain/domains.svelte";
import AdminTenantDomain from "./tenant/domain/domain.svelte";

// plug
import AdminPlugs from "./plug/plugs.svelte";
import AdminPlug from "./plug/plug.svelte";

export default {
  $$component: Admin,
  "/bprint": {
    "/": AdminBprints,
    "/:bid": AdminBprint,
    "/:bid/editor": AdminBprintEditor,
    "/:bid/instancer": AdminBprintInstancer,
  },
  "/repo": {
    "/": AdminRepos,
    "/:rid/edit": AdminRepoEdit,
    "/new": AdminRepoNew,
  },

  "/data": {
    "/": AdminDataLoader,
    "/group": AdminDataGroups,
    "/group/:group": AdminDataGroup,
    "/table/:group": AdminDataTables,
    "/table/:group/:table": AdminDataTable,
    "/column/:group/:table": AdminDataColumns,
    "/column/:group/:table/:column": AdminDataColumn,
    "/hook/:group/:table": AdminDataHooks,
    "/hook/:group/:table/:id": AdminDataHook,
    "/view/:group/:table": AdminDataViews,
    "/view/:group/:table/:id": AdminDataView,
  },

  "/resource": {
    "/": AdminResources,
    "/:rid/edit": AdminResourceEdit,
    "/new": AdminResourceNew,
  },

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
};
