<script lang="ts">
  import Renderer from "./renderer/renderer.svelte";
  import Skeleton from "./renderer/skeleton.svelte";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../lib/app/portal";
  import type { DataTableService } from "../../../../lib/service/dyn";

  export let source;
  export let dgroup;
  export let dtable;

  const app: PortalApp = getContext("__app__");

  let dman: DataTableService;

  $: {
    dman = null;

    app.get_data_service(source, dgroup).then(async (svc) => {
      dman = await svc.get_table_service(dtable, app.navigator.nav_options);
    });
  }
</script>

{#key dtable}
  {#if dman}
    <Renderer
      {source}
      {dgroup}
      {dtable}
      manager={dman}
      navigator={app.navigator}
    />
  {:else}
    <Skeleton />
  {/if}
{/key}
