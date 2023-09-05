<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "$lib/core";
  import { params } from "$lib/params";

  export let ugroup = $params["ugroup"];

  const app = getContext("__app__") as PortalService;

  let datas = [];
  let loading = true;
  const api = app.api_manager.get_admin_ugroup_api();

  const load = async () => {
    const resp = await api.listAuth(ugroup);
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_edit = (id: string) =>
    app.nav.admin_ugroup_auth_edit(ugroup, id);
  const action_delete = async (id: string) => {
    const resp = await api.deleteAuth(ugroup, id);
    load();
  };

  const action_new = () => app.nav.admin_ugroup_auth_new(ugroup);
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
      ["id", "Id"],
      ["name", "Name"],
      ["type", "Type"],
      ["provider", "Provider"],
    ]}
    color={[]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
