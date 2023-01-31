<script lang="ts">
  import type { SheetColumn } from "../sheets";
  import Layout from "./_layout.svelte";
  import Cell from "../field/cell.svelte";

  export let columns: SheetColumn[];
  export let open_column;

  let dirty_data = {};
</script>

<Layout title="Add Row" onClick={() => {}}>
  {#each columns as col}
    <Cell column={col} bind:open_column celldata={dirty_data[col.__id]} />
  {/each}
</Layout>
