<script>
  import Builder from "./_builder/builder.svelte";
  import { params } from "svelte-hash-router";
  export let pid = $params.pid;
</script>

<Builder />
