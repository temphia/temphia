<script lang="ts">
  import type { SheetColumn } from "../../sheets";
  import Cell from "../../field/cell.svelte";
  import type { SheetService } from "../../../../../services/data";

  export let columns: SheetColumn[];
  export let service: SheetService;

  export let current_cells = {};

  let open_column;
</script>

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
