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

<Layout
  title="Edit Row"
  onSave={() => onSave(current_cells)}
  new_record={false}
>
  <svelte:fragment slot="edit">
    <EditRow {cells} {columns} {row} {service} {onSave} bind:current_cells />
  </svelte:fragment>

  <svelte:fragment slot="relation">
    <div>relation</div>
  </svelte:fragment>

  <svelte:fragment slot="history">
    <div>history</div>
  </svelte:fragment>
</Layout>
