<script lang="ts">
  import Builder from "./_builder.svelte";
  import { LoadingSpinner, PortalService } from "../../../core";
  import * as b from "./builder";

  export let bid: string;
  export let app: PortalService;
  export let etype: string;
  export let file: string;

  let loading = false;
  let data: b.State;

  const load = async () => {
    const bapi = app.api_manager.get_admin_bprint_api();
    const resp = await bapi.get_file(bid, file);
    if (!resp.ok) {
      console.log("Err", resp.data);
      return;
    }

    data = resp.data;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <Builder
    builder={b.Builder.from_batch(data)}
    open_modal={(comp, opts) => app.utils.small_modal_open(comp, opts)}
    close_modal={() => app.utils.small_modal_close()}
  />
{/if}
