<script lang="ts">
  import type { SheetCell, SheetColumn, SheetRow } from "../sheets";
  import Layout from "./_layout.svelte";
  import Cell from "../field/cell.svelte";
  import type { SheetService } from "../../../../services/data";

  export let columns: SheetColumn[];
  export let row: SheetRow;
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let service: SheetService;
  export let onSave = async (data) => {};

  let open_column;

  const current_cells = cells[row.__id] || {};
</script>

<Layout title="Edit Row" onClick={() => onSave(current_cells)}>
  {#each columns as col}
    <Cell
      {service}
      column={col}
      bind:open_column
      celldata={current_cells[col.__id]}
      onCellChange={(data) => {
        const old = current_cells[col.__id] || {};
        current_cells[col.__id] = { ...old, ...data };
      }}
    />
  {/each}
</Layout>
