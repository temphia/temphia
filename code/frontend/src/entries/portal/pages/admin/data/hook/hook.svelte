<script lang="ts">
  import Hook from "../../target/hook/hook_new.svelte";
  import { params } from "svelte-hash-router";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.group;

</script>

<Hook ttype="data" target={`${source}/${group}/${table}`} />
