<!-- <script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";

  let datas = [];
  let loading = true;
  const app = getContext("__app__") as PortalService;

  const load = async () => {
    const api = app.api_manager.get_admin_plug_api();
    const resp = await api.list_plug();
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_execute = (id: string) =>
    app.nav.admin_plug_dev_execute(id, "default");



  const action_new = () => {};


</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="id"
    actions={[]}
    key_names={[]}
    color={["executor"]}
    {datas}
    show_drop={true}
  />
{/if}

<FloatingAdd onClick={action_new} /> -->
