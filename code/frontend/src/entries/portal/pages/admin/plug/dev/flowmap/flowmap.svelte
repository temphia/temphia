<script lang="ts">
  import { params } from "svelte-hash-router";
  import FlowmapInner from "./_flowmap_inner.svelte";
  import WithScript from "./_with_script.svelte";
  export let pid = $params.pid;
</script>

<WithScript script="/z/assets/lib/leader-line.min.js" key="LeaderLine">
  {#key pid}
    <FlowmapInner {pid} />
  {/key}
</WithScript>
