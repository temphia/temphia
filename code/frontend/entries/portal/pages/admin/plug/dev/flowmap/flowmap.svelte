<script lang="ts">
  import { params } from "svelte-hash-router";
  import FlowmapInner from "./_flowmap_inner.svelte";
  export let pid = $params.pid;
</script>

{#key pid}
  <FlowmapInner {pid} />
{/key}
