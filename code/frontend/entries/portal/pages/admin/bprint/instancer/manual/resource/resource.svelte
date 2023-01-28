<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";

  import type { PortalService } from "../../../../core";

  export let bid: string = $params.bid;
  export let file: string = $params._ || "schema.json"

  const app: PortalService = getContext("__app__");

  let loading = true;
  let data: object;

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

{#if !loading}
  <div>Resource</div>
{/if}
