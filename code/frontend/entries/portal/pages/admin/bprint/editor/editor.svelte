<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../core";
  import { params } from "svelte-hash-router";

  export let bid: string = $params.bid;
  export let etype: string = $params.etype;
  export let file: string = $params._;

  const app = getContext("__app__") as PortalService;
  const compo = app.xtmgr.bprint_editors.get(etype);
</script>

<svelte:component this={compo} {bid} {app} {etype} {file} />
