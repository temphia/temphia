<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../core";
  import TopActions from "../core/top_actions.svelte";

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_ugroup_api();

  const load = async () => {
    const resp = await api.list();
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_edit = (id: string) => app.nav.admin_ugroup_edit(id);
  const action_delete = async (id: string) => {
    const resp = await api.delete(id);
    load();
  };

  const action_new = () => app.nav.admin_ugroup_new();
  const action_explore_group_user = (id: string) =>
    app.nav.admin_ugroup_users(id);

  const action_group_auths = (id: string) => app.nav.admin_ugroup_auths(id);
</script>

<TopActions
  actions={{
    "All Users": () => app.nav.admin_users(),
  }}
/>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="slug"
    show_drop={true}
    actions={[
      {
        Name: "Users",
        Class: "bg-green-400",
        Action: action_explore_group_user,
        icon: "users",
      },

      {
        Name: "Edit",
        Action: action_edit,
        icon: "pencil-alt",
      },
      {
        Name: "Auths",
        Action: action_group_auths,
        icon: "shield-exclamation",
        drop: true,
      },

      {
        Name: "Apps",
        Action: (id) => {},
        icon: "view-grid-add",
        drop: true,
      },

      {
        Name: "Datas",
        Action: (id) => {},
        icon: "collection",
        drop: true,
      },
      

      {
        Name: "Delete",
        Action: action_delete,
        icon: "trash",
        drop: true,
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["slug", "Slug"],
      ["scopes", "Scopes"],
      ["mod_version", "Mod Version"]
    ]}
    color={[]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
