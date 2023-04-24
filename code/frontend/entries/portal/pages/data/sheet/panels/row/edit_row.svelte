<script lang="ts">
  import type { SheetCell, SheetColumn, SheetRow } from "../../sheets";
  import Layout from "./_layout.svelte";
  import type { SheetService } from "../../../../../services/data";
  import EditRow from "./_edit_row.svelte";
  import Relations from "./_relations.svelte";
  import History from "./_history.svelte";

  export let columns: SheetColumn[];
  export let row: SheetRow;
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let service: SheetService;
  
  export let onSave = async (data) => {};
  export let gotoSiblingSheet = (ssid, rowid) => {}

  let current_cells = cells[row.__id] || {};

  $: console.log("@current_cells", current_cells);


</script>

<Layout
  title="Edit Row"
  onSave={() => onSave(current_cells)}
  onDelete={async () => {
    await service.remove_row_cell(String(row.__id));
    service.close_big_modal();
  }}
  new_record={false}
>
  <svelte:fragment slot="edit">
    <EditRow {columns} {service} bind:current_cells />
  </svelte:fragment>

  <svelte:fragment slot="relation">
    <Relations {service} rid={row.__id} {gotoSiblingSheet} />
  </svelte:fragment>

  <svelte:fragment slot="history">
    <History {service} rid={row.__id} />
  </svelte:fragment>
</Layout>
