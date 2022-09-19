<script lang="ts">
  import Layout from "../../layout.svelte";
  import HookNew from "./_hook_new.svelte";
  export let table = "";
  export let group = "";
  export let source = "";
</script>

<Layout current_item={"dtable"}>
  <HookNew group_id={group} {source} table_id={table} />
</Layout>
