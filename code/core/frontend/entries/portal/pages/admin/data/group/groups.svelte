<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import { params } from "svelte-hash-router";
  import NewPick from "./_new_pick.svelte";

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

  const action_edit = (id: string) => app.nav.admin_data_group(source, id);
  const action_explore = (id: string) => app.nav.admin_data_tables(source, id);
  const action_delete = async (id: string) => {};
  const action_new = () => {
    app.utils.small_modal_open(NewPick, { app });
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="slug"
    actions={[
      {
        Name: "Edit",
        Action: action_edit,
        icon: "pencil-alt",
      },
      {
        Class: "bg-green-400",
        Name: "explore",
        icon: "book-open",
        Action: action_explore,
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
