<script lang="ts">
  import type { SheetColumn } from "../sheets";
  import Layout from "./_layout.svelte";
  import Cell from "./__cell.svelte";

  export let columns: SheetColumn[];
  export let open_column;
</script>

<Layout title="Add Row" onClick={() => {}}>
  {#each columns as col}
    <Cell column={col} bind:open_column />
  {/each}
</Layout>
