import AdminBprint from "./bprint/bprint.svelte";
import AdminBprints from "./bprint/bprints.svelte";
import AdminBprintEditor from "./bprint/editor/editor.svelte";
import AdminBprintDataInstancer from "./bprint/instancer/data/data.svelte";
import AdminBprintPlugInstancer from "./bprint/instancer/plug/plug.svelte";

// plug
import AdminPlugs from "./plug/plugs.svelte";
import AdminPlugNew from "./plug/plug_new.svelte";
import AdminPlugEdit from "./plug/plug_edit.svelte";
import AdminAgents from "./plug/agent/agents.svelte";
import AdminAgentNew from "./plug/agent/agent_new.svelte";
import AdminAgentEdit from "./plug/agent/agent_edit.svelte";

import AdminPlugDevDocs from "./plug/dev/docs/docs.svelte";
import AdminPlugDevExec from "./plug/dev/execute/execute.svelte";
import AdminPlugDevShell from "./plug/dev/shell/shell.svelte";
import AdminPlugDevFlowmap from "./plug/dev/flowmap/flowmap.svelte";

import AdminAgentLinks from "./plug/agent/link/links.svelte";
import AdminAgentLinkNew from "./plug/agent/link/link_new.svelte";
import AdminAgentLinkEdit from "./plug/agent/link/link_edit.svelte";

import AdminAgentExtensions from "./plug/agent/extension/extensions.svelte";
import AdminAgentExtensionNew from "./plug/agent/extension/extension_new.svelte";
import AdminAgentExtensionEdit from "./plug/agent/extension/extension_edit.svelte";

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
import AdminDataApp from "./data/app/app.svelte";
import AdminDataApps from "./data/app/apps.svelte";

import AdminDataViews from "./data/view/views.svelte";
import AdminDataViewNew from "./data/view/view_new.svelte";
import AdminDataViewEdit from "./data/view/view_edit.svelte";
import Admin from "./admin.svelte";

// resource
import AdminResources from "./resource/resources.svelte";
import AdminResourceEdit from "./resource/resource_edit.svelte";
import AdminResourceNew from "./resource/resource_new.svelte";

import AdminResourceDatagroupNew from "./resource/datagroup/datagroup_new.svelte";
import AdminResourceDatagroupEdit from "./resource/datagroup/datagroup_edit.svelte";
import AdminResourceFolderNew from "./resource/folder/folder_new.svelte";
import AdminResourceFolderEdit from "./resource/folder/folder_edit.svelte";
import AdminResourceRoomNew from "./resource/room/room_new.svelte";
import AdminResourceRoomEdit from "./resource/room/room_edit.svelte";

// user
import AdminUsers from "./user/users.svelte";
import AdminUser from "./user/user.svelte";

// ugroup
import AdminUgroups from "./ugroup/ugroups.svelte";
import AdminUgroupNew from "./ugroup/ugroup_new.svelte";
import AdminUgroupEdit from "./ugroup/ugroup_edit.svelte";
import AdminUgroupUsers from "./ugroup/user/users.svelte";
import AdminUgroupUser from "./ugroup/user/user.svelte";

// lens
import AdminLens from "./lens/lens.svelte";
import AdminLensApp from "./lens/app.svelte";
import AdminLensEngine from "./lens/engine.svelte";
import AdminLensSite from "./lens/site.svelte";

// tenant
import AdminTenant from "./tenant/tenant.svelte";
import AdminTenantEdit from "./tenant/tenant_edit.svelte";
import AdminTenantDomainEdit from "./tenant/domain/domain_edit.svelte";
import AdminTenantDomainAdd from "./tenant/domain/domain_add.svelte";

import AdminTargetApps from "./target/app/apps.svelte";
import AdminTargetAppNew from "./target/app/app_new.svelte";
import AdminTargetAppEdit from "./target/app/app_edit.svelte";

import AdminTargetHooks from "./target/hook/hooks.svelte";
import AdminTargetHookNew from "./target/hook/hook_new.svelte";
import AdminTargetHookEdit from "./target/hook/hook_edit.svelte";

export default {
  $$component: Admin,
  "/bprint": {
    "/": AdminBprints,
    "/:bid": AdminBprint,
    "/:bid/editor": AdminBprintEditor,
    "/:bid/instance/data/*": AdminBprintDataInstancer,
    "/:bid/instance/plug/*": AdminBprintPlugInstancer,
  },

  "/plug": {
    "/": AdminPlugs,
    "/:pid/edit": AdminPlugEdit,
    "/:pid/dev": {
      "/flowmap": AdminPlugDevFlowmap,
      "/execute/:aid": AdminPlugDevExec,
      "/shell/:aid": AdminPlugDevShell,
      "/docs/:aid": AdminPlugDevDocs,
    },

    "/new": AdminPlugNew,
    "/:pid/agent": {
      "/": AdminAgents,
      "/:aid/edit": AdminAgentEdit,
      "/new": AdminAgentNew,
      "/:aid/link": {
        "/": AdminAgentLinks,
        "/new": AdminAgentLinkNew,
        "/:eid/edit": AdminAgentLinkEdit,
      },
      "/:aid/ext": {
        "/": AdminAgentExtensions,
        "/new": AdminAgentExtensionNew,
        "/:eid/edit": AdminAgentExtensionEdit,
      },
    },
  },

  "/repo": {
    "/": AdminRepos,
    "/:rid/edit": AdminRepoEdit,
    "/new": AdminRepoNew,
  },

  "/resource": {
    "/": AdminResources,
    "/:rid/edit": AdminResourceEdit,
    "/new": AdminResourceNew,

    "/rtype/data_group/new": AdminResourceDatagroupNew,
    "/rtype/data_group/:rid/edit": AdminResourceDatagroupEdit,
    "/rtype/room/new": AdminResourceFolderNew,
    "/rtype/room/:rid/edit": AdminResourceFolderEdit,
    "/rtype/folder/new": AdminResourceRoomNew,
    "/rtype/folder/:rid/edit": AdminResourceRoomEdit,
  },

  "/data": {
    "/": AdminDataLoader,
    "/:source/group": AdminDataGroups,
    "/:source/group/:group": AdminDataGroup,
    "/:source/table/:group": AdminDataTables,
    "/:source/table/:group/:table": AdminDataTable,
    "/:source/column/:group/:table": AdminDataColumns,
    "/:source/column/:group/:table/:column": AdminDataColumn,
    "/:source/hook/:group/:table": AdminDataHooks,
    "/:source/hook/:group/:table/:id": AdminDataHook,
    "/:source/app/:group/:table": AdminDataApps,
    "/:source/app/:group/:table/:id": AdminDataApp,
    "/:source/view/:group/:table": AdminDataViews,
    "/:source/view/:group/:table/new": AdminDataViewNew,
    "/:source/view/:group/:table/:id/edit": AdminDataViewEdit,
  },

  "/ugroup": {
    "/": AdminUgroups,
    "/new": AdminUgroupNew,
    "/:ugroup/edit": AdminUgroupEdit,
    "/:ugroup/user": {
      "/": AdminUgroupUsers,
      "/:userid/edit": AdminUgroupUser,
    },
  },

  "/user": {
    "/": AdminUsers,
    "/:userid": AdminUser,
  },

  "/lens": {
    $$component: AdminLens,
    "/app": AdminLensApp,
    "/engine": AdminLensEngine,
    "/site": AdminLensSite,
  },

  "/target": {
    "/app": {
      "/": AdminTargetApps,
      "/new": AdminTargetAppNew,
      "/:ttype/:id/edit": AdminTargetAppEdit,
    },
    "/hook": {
      "/": AdminTargetHooks,
      "/new": AdminTargetHookNew,
      "/:ttype/:id/edit": AdminTargetHookEdit,
    },
  },

  "/tenant": {
    "/": AdminTenant,
    "/edit": AdminTenantEdit,
    "/domain/:did/edit": AdminTenantDomainEdit,
    "/domain/new": AdminTenantDomainAdd,
  },
};
