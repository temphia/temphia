<script lang="ts">
  import type { SheetCell, SheetColumn, SheetRow } from "../sheets";
  import Layout from "./_layout.svelte";
  import Cell from "../field/cell.svelte";

  export let columns: SheetColumn[];
  export let row: SheetRow;
  export let cells: { [_: number]: { [_: string]: SheetCell } };

  let open_column;

  const current_cells = cells[row.__id];
</script>

<Layout title="Edit Row" onClick={() => {}}>
  {#each columns as col}
    <Cell column={col} bind:open_column celldata={current_cells[col.__id]} />
  {/each}
</Layout>
