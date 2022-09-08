<script lang="ts">
  import PlugEdit from "./plug_edit.svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";
  import { getContext } from "svelte";
  import Layout from "../layout.svelte";
  export let id;

  const app: PortalApp = getContext("__app__");

  let plug;

  app
    .get_apm()
    .get_plug_api()
    .then(async (papi) => {
      const resp = await papi.get_plug(id);
      plug = resp.data;
    });
</script>

<Layout>
  {#if plug}
    <PlugEdit data={plug} />
  {/if}
</Layout>
