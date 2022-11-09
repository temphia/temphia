<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import { params } from "svelte-hash-router";

  export let source = $params.source;
  export let group = $params.group;

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_data_api();

  const load = async () => {
    const resp = await api.list_tables(source, group);
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_edit = (id: string) =>
    app.nav.admin_data_table(source, group, id);
  const action_delete = async (id: string) => {};
  const action_new = () => {};

  const action_goto_columns = (id) => app.nav.admin_data_columns(source, group, id);
  const action_goto_views = (id) => app.nav.admin_data_views(source, group, id);
  const action_goto_data_hooks = (id) =>
    app.nav.admin_data_hooks(source, group, id);
  const action_goto_data_apps = (id) =>
    app.nav.admin_data_apps(source, group, id);
  const action_goto_data = (id) => app.nav.data_table(source, group, id);
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="slug"
    actions={[
      {
        Class: "bg-green-400",
        Name: "Columns",
        Action: action_goto_columns,
      },

      {
        Name: "Views",
        Action: action_goto_views,
      },

      {
        Name: "Hooks",
        Class: "bg-green-400",
        Action: action_goto_data_hooks,
      },

      {
        Name: "Apps",
        Class: "bg-green-400",
        Action: action_goto_data_apps,
      },

      {
        Name: "Data",
        Class: "bg-yellow-400",
        Action: action_goto_data,
      },

      {
        Name: "Edit",
        Action: action_edit,
        icon: "pencil-alt",
      },

      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: action_delete,
        icon: "trash",
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["slug", "Slug"],
      ["description", "Description"],
    ]}
    color={[]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
