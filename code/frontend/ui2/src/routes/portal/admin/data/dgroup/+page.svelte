<script lang="ts">
  import { params } from "$lib/params";
  import Groups from "./_groups.svelte";
</script>

{#key $params["source"]}
  <Groups source={$params["source"]} />
{/key}
