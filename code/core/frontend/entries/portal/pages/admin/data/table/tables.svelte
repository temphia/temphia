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

  const action_edit = (id: string) => {};
  const action_delete = async (id: string) => {};
  const action_new = () => {};

  const action_goto_views = (id) => {};
  const action_goto_data_hooks = (id) => {};
  const action_goto_data_apps = (id) => {};
  const action_goto_data = (id) => {};
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="slug"
    actions={[
      {
        Class: "bg-green-400",
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
