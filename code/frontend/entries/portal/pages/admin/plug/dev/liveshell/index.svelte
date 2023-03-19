<script lang="ts">
  import LiveshellInner from "./liveshell_inner.svelte";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner, PortalService } from "../../../core";
  import { getContext } from "svelte";

  const app = getContext("__app__") as PortalService;
  const bapi = app.api_manager.get_admin_bprint_api();
  const papi = app.api_manager.get_admin_plug_api();

  let loading = true;
  let bid;

  let files = [];

  const load = async () => {
    loading = true;
    const resp = await papi.get_plug($params.pid);
    if (!resp.ok) {
      console.log("@resp", resp);
      return;
    }

    bid = resp.data["bprint_id"];

    const bresp = await bapi.list_file(bid);
    if (!bresp.ok) {
      console.log("@resp", bresp);
      return;
    }
    files = Object.keys(bresp.data);
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <LiveshellInner {files} file={files[0] || ""} {bid} />
{/if}
