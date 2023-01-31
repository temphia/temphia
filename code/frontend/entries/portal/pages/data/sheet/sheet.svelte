<script lang="ts">
  import { params } from "svelte-hash-router";
  import SheetInner from "./_sheet_inner.svelte";
</script>

{#key $params.sheet}
  <SheetInner
    group={$params.dgroup}
    sheetid={$params.sheet}
    source={$params.source}
  />
{/key}
