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

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_data_api();

  const load = async () => {
    const resp = await api.list_group(source);
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
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="id"
    actions={[
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
