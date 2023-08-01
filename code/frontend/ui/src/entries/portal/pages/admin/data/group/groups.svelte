<script lang="ts">
  import { params } from "svelte-hash-router";
  import Groups from "./_groups.svelte";
</script>

{#key $params.source}
  <Groups source={$params.source} />
{/key}
