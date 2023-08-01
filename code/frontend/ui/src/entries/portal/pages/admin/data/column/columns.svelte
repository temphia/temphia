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
  export let table = $params.table;


  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_data_api();

  const load = async () => {
    const resp = await api.list_column(source, group, table);
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_edit = (id: string) =>
    app.nav.admin_data_column(source, group, table, id);
  const action_delete = async (id: string) => {};
  const action_new = () => {};

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
        Name: "Delete",
        Class: "bg-red-400",
        Action: action_delete,
        icon: "trash",
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["slug", "Slug"],
      ["ctype", "Column type"],
      ["description", "Description"],
    ]}
    color={[]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
