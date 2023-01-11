<script lang="ts">
  import { params } from "svelte-hash-router";
  import LensInner from "./_lens_inner.svelte";
</script>

{#key $params.index}
  <LensInner index={$params.index} />
{/key}
