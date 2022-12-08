<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import { params } from "svelte-hash-router";

  export let bid = $params.bid;

  let datas = [];
  let loading = true;
  const app = getContext("__app__") as PortalService;

  const load = async () => {
    const api = app.api_manager.get_admin_bprint_api();
    const resp = await api.list_file(bid);
    if (!resp.ok) {
      return;
    }

    datas = Object.entries(resp.data).map(([key, value]) => ({
      name: key,
      value: value,
    }));

    loading = false;
  };

  load();

  // actions

  const action_new = () => {}
  const action_edit_file = async (name: string) =>
    app.nav.admin_bprint_file(bid, name);
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    show_drop={true}
    color={[]}
    action_key="name"
    actions={[
      {
        Name: "Edit",
        Class: "bg-blue-400",
        icon: "pencil-alt",
        Action: action_edit_file,
      },

      {
        Name: "Edit",
        Class: "bg-green-400",
        icon: "download",
        Action: (name) => {},
      },
    ]}
    key_names={[
      ["name", "name"],
      ["value", "value"],
    ]}
    {datas}
  />
{/if}

<FloatingAdd onClick={action_new} />
