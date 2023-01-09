<script lang="ts">
  import DataGroup from "./_data_group.svelte";
  import { getContext } from "svelte";
  import { PortalService, LoadingSpinner } from "../../../../core";
  import { params } from "svelte-hash-router";

  export let bid = $params.bid;
  export let file = $params._;

  const app: PortalService = getContext("__app__");

  let loading = true;
  let data: any;

  const load = async () => {
    const bapi = await app.api_manager.get_admin_bprint_api();
    const resp = await bapi.get_file(bid, file);

    if (resp.status !== 200) {
      console.log(resp);
      return;
    }
    console.log("@file", resp.data);
    data = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <DataGroup {data} {bid} {file} {app} />
{/if}
