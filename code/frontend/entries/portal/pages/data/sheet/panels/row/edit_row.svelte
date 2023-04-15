<script lang="ts">
  import type { SheetCell, SheetColumn, SheetRow } from "../../sheets";
  import Layout from "./_layout.svelte";
  import type { SheetService } from "../../../../../services/data";
  import EditRow from "./_edit_row.svelte";

  export let columns: SheetColumn[];
  export let row: SheetRow;
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let service: SheetService;
  export let onSave = async (data) => {};

  let open_column;

  let current_cells = cells[row.__id] || {};
</script>

<Layout title="Edit Row" onClick={() => onSave(current_cells)}>
  <EditRow {cells} {columns} {row} {service} {onSave} bind:current_cells />
</Layout>
