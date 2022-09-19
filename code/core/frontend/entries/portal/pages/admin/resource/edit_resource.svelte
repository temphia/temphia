<script lang="ts">
  import EditResource from "./_edit_resource.svelte";
  import Layout from "../layout.svelte";
  import type { PortalApp } from "../../../app";
  import { getContext } from "svelte";

  export let id = "";

  const app: PortalApp = getContext("__app__");
  let data;

  const load = async () => {
    const rapi = await app.get_apm().get_resource_api();
    const resp = await rapi.resource_get(id);
    data = resp.data;
  };
  
  load()
</script>

<Layout current_item="resources" loading={!!data}>
  {#if data}
    <EditResource {data} />
  {/if}
</Layout>
