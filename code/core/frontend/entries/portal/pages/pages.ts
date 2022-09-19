import SelfProfile from "./user/self_profile.svelte";
import UserProfile from "./user/user_profile.svelte";
import Start from "./start/start.svelte";

import CabinetLoader from "./cabinet/loader.svelte";
import CabinetSource from "./cabinet/source.svelte";
import CabinetFolder from "./cabinet/folder.svelte";
import CabinetFile from "./cabinet/file.svelte";

import DgroupLoader from "./dtable/loader.svelte";
import Dgroups from "./dtable/dgroups.svelte";
import Dgroup from "./dtable/dgroup.svelte";
import Dtable from "./dtable/dtable.svelte";

import Store from "./store/store.svelte";
import StoreItem from "./store/item.svelte";

import AdminPlugs from "./admin/plug/plugs.svelte";
import AdminPlug from "./admin/plug/plug.svelte";
import AdminPlugAgents from "./admin/plug/agent/agents.svelte";
import AdminPlugAgent from "./admin/plug/agent/agent.svelte";
import AdminPlugResources from "./admin/plug/plug_resources.svelte";
import AdminAgentDevShell from "./admin/plug/dev/shell.svelte";
import AdminAgentIfaceDoc from "./admin/plug/dev/ifacedoc.svelte";

import AdminRepos from "./admin/repo/repos.svelte"
import AdminRepoNew from "./admin/repo/repo_new.svelte"
import AdminRepoEdit from "./admin/repo/repo_edit.svelte"

import InstancePlug from "./admin/blueprint/instancer/instance_plug.svelte";
import InstanceDataGroup from "./admin/blueprint/instancer/instance_datagroup.svelte";
import InstanceDataTable from "./admin/blueprint/instancer/instance_datatable.svelte";

import AdminAgentNewLink from "./admin/plug/link/link_new.svelte";
import AdminAgentEditLink from "./admin/plug/link/link_edit.svelte";
import AdminAgentLinks from "./admin/plug/link/links.svelte";
import AdminAgentNewExt from "./admin/plug/extension/extension_new.svelte";
import AdminAgentEditExt from "./admin/plug/extension/extension_edit.svelte";
import AdminAgentExts from "./admin/plug/extension/extensions.svelte";

import AdminAgentNewResource from "./admin/plug/resource/resource_new.svelte";
import AdminAgentEditResource from "./admin/plug/resource/resource_edit.svelte";
import AdminAgentResources from "./admin/plug/resource/resources.svelte";

import AdminResources from "./admin/resource/resources.svelte";
import AdminNewResource from "./admin/resource/new_resource.svelte";
import AdminEditResource from "./admin/resource/edit_resource.svelte";

import AdminDtableBuilder from "./admin/blueprint/builder/builder.svelte";

import AdminUsers from "./admin/user/users.svelte";
import AdminUser from "./admin/user/user.svelte";
import AdminNewUser from "./admin/user/new_user.svelte";
import AdminUserByGroup from "./admin/user/users_by_group.svelte";
import AdminUserGroups from "./admin/usergroup/groups.svelte";
import AdminUserGroup from "./admin/usergroup/group.svelte";
import AdminNewUserGroup from "./admin/usergroup/new_group.svelte";

import AdminBrint from "./admin/blueprint/blueprint.svelte";
import AdminBrints from "./admin/blueprint/blueprints.svelte";

import AdminTenant from "./admin/tenant/tenant.svelte";
import AdminTenantEdit from "./admin/tenant/tenant_edit.svelte";
import AdminDomainNew from "./admin/tenant/domain/domain_new.svelte";
import AdminDomainEdit from "./admin/tenant/domain/domain_edit.svelte";
import AdminWidgetNew from "./admin/tenant/domain/widget_new.svelte";
import AdminWidgetEdit from "./admin/tenant/domain/widget_edit.svelte";
import AdminWidgets from "./admin/tenant/domain/widgets.svelte";
import AdminLens from "./admin/lens/lens.svelte"

import ListDgroup from "./admin/dtable/list_dgroup.svelte";
import ListDtable from "./admin/dtable/list_dtable.svelte";
import ListColumn from "./admin/dtable/list_column.svelte";
import EditColumn from "./admin/dtable/edit_column.svelte";
import EditDgroup from "./admin/dtable/edit_dgroup.svelte";
import EditDtable from "./admin/dtable/edit_dtable.svelte";

import ListViews from "./admin/dtable/view/list_views.svelte";
import ViewEdit from "./admin/dtable/view/view_edit.svelte";
import ViewNew from "./admin/dtable/view/view_new.svelte";

import ListHooks from "./admin/dtable/hook/list_hooks.svelte";
import HookEdit from "./admin/dtable/hook/hook_edit.svelte";
import HookNew from "./admin/dtable/hook/hook_new.svelte";

import UserGroupAuthNew from "./admin/usergroup/auth/new_auth.svelte";
import UserGroupAuthEdit from "./admin/usergroup/auth/edit_auth.svelte";
import UserGroupPlugNew from "./admin/usergroup/plug/new_plug.svelte";
import UserGroupPlugEdit from "./admin/usergroup/plug/edit_plug.svelte";
import UserGroupHookNew from "./admin/usergroup/hook/new_hook.svelte";
import UserGroupHookEdit from "./admin/usergroup/hook/edit_hook.svelte";
import UserGroupDataNew from "./admin/usergroup/data/new_data.svelte";
import UserGroupDataEdit from "./admin/usergroup/data/edit_data.svelte";

import Launcher from "./launcher/launcher.svelte";
import AppsLauncher from "./launcher/apps.svelte";
import AboutTenant from "./tenant/about_tenant.svelte";

import Login from "./auth/login.svelte";

export {
  Start,
  Login,
  CabinetLoader,
  CabinetSource,
  CabinetFolder,
  CabinetFile,
  DgroupLoader,
  Dgroups,
  Dgroup,
  Dtable,
  Store,
  StoreItem,
  AdminPlugs,
  AdminPlug,
  AdminPlugAgents,
  AdminPlugResources,
  AdminPlugAgent,

  AdminRepos,
  AdminRepoNew,
  AdminRepoEdit,

  InstancePlug,
  InstanceDataGroup,
  InstanceDataTable,
  AdminLens,

  AdminAgentDevShell,
  AdminAgentIfaceDoc,
  AdminAgentNewLink,
  AdminAgentEditLink,
  AdminAgentLinks,
  AdminAgentNewExt,
  AdminAgentEditExt,
  AdminAgentExts,
  AdminAgentNewResource,
  AdminAgentEditResource,
  AdminAgentResources,
  AdminResources,
  AdminNewResource,
  AdminEditResource,
  AdminDtableBuilder,
  AdminBrint,
  AdminBrints,
  AdminUsers,
  AdminUser,
  AdminNewUser,
  AdminUserGroups,
  AdminUserGroup,
  AdminNewUserGroup,
  AdminUserByGroup,
  ListDgroup,
  ListDtable,
  ListColumn,
  EditColumn,
  EditDgroup,
  EditDtable,
  ListViews,
  ViewEdit,
  ViewNew,
  ListHooks,
  HookEdit,
  HookNew,
  UserGroupAuthNew,
  UserGroupAuthEdit,
  UserGroupPlugNew,
  UserGroupPlugEdit,
  UserGroupHookNew,
  UserGroupHookEdit,
  UserGroupDataNew,
  UserGroupDataEdit,
  SelfProfile,
  UserProfile,
  Launcher,
  AppsLauncher,
  AdminTenant,
  AdminTenantEdit,
  AboutTenant,
  AdminDomainNew,
  AdminDomainEdit,
  AdminWidgetNew,
  AdminWidgetEdit,
  AdminWidgets,
};
