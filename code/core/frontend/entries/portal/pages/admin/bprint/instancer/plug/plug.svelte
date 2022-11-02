<script lang="ts">
  import { getContext } from "svelte";

  import Plug from "./_plug.svelte";
  import type { PlugRawSchema } from "../instance";
  import type { PortalService } from "../../../core";

  export let bid: string;
  export let file: string;

  const app: PortalService = getContext("__app__");

  let loading = true;
  let data: PlugRawSchema;

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
  <Plug {data} {app} {bid} {file} />
{/if}
