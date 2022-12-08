<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner, PortalService } from "../../core";

  export let bid = $params.bid;
  export let file = $params._;

  let loading = true;
  let text = "";

  const app = getContext("__app__") as PortalService;

  const load = async () => {
    const api = app.api_manager.get_admin_bprint_api();
    const resp = await api.get_file(bid, file);
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    if (typeof resp.data === "object") {
      text = JSON.stringify(resp.data);
    } else {
      text = resp.data;
    }

    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <pre>{text}</pre>
{/if}
