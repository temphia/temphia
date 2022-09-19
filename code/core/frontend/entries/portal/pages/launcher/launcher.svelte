<script lang="ts">
  import { getContext, onMount } from "svelte";
  import type { PortalApp } from "../../../../lib/app/portal";

  export let plugid = "";
  export let agentid = "";

  let launchRef;

  const app: PortalApp = getContext("__app__");
  let exec;

  onMount(async () => {
    exec = await app.elauncher.instance({
      agent: agentid,
      handler: null,
      plug: plugid,
      target: launchRef,
      data: {},
    });
    console.log(exec);
  });
</script>

<div class="h-full w-full" bind:this={launchRef} />
