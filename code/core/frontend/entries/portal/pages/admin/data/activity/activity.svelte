<script lang="ts">
  import { getContext } from "svelte";
  import { AutoForm, LoadingSpinner, PortalService } from "../../core";
  import { params } from "svelte-hash-router";
  import ActivityInner from "./_activity_inner.svelte";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.table;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let data = {};
  let loading = true;

  const load = async () => {
    loading = false;
  };

  load();

  const save = async (_data) => {
    const resp = await api.edit_table(source, group, table, _data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }
    app.nav.admin_data_groups(source);
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <ActivityInner />
{/if}
