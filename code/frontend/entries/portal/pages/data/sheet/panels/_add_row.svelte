<script lang="ts">
  import type { SheetColumn } from "../sheets";
  import Layout from "./_layout.svelte";
  import Row from "./__row.svelte";
  export let columns: SheetColumn[];
</script>

<Layout title="Add Row" onClick={() => {}}>
  <Row {columns} />
</Layout>
