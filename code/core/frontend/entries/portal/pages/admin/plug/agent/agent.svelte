<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import AgentEditor from "./agent_edit.svelte";
  import Layout from "../../layout.svelte";

  export let pid;
  export let aid;

  let agent;

  const app: PortalApp = getContext("__app__");

  app
    .get_apm()
    .get_plug_api()
    .then(async (api) => {
      const resp = await api.get_agent(pid, aid);
      agent = resp.data;
    });
</script>

<Layout>
  {#if agent}
    <AgentEditor data={agent} />
  {/if}
</Layout>
