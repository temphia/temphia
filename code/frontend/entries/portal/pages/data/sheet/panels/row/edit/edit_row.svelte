<script lang="ts">
  import type { SheetCell, SheetColumn, SheetRow } from "../../../sheets";
  import Layout from "../_layout.svelte";
  import type { SheetService } from "../../../../../../services/data";
  import EditRow from "./_edit_row.svelte";
  import Relations from "../relations/_relations.svelte";
  import History from "../history/_history.svelte";

  export let columns: SheetColumn[];
  export let row: SheetRow;
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let service: SheetService;

  export let onSave: (data) => Promise<any>;
  export let gotoSiblingSheet = (ssid, rowid) => {};

  let current_cells = cells[row.__id] || {};
  let modified_cells = {};

  $: console.log("@current_cells", current_cells);

  let message = "";
</script>

<Layout
  title="Edit Row"
  onSave={async () => {
    const finalData = { ...current_cells };

    Object.keys(finalData).forEach((colkey) => {
      if (!modified_cells[colkey]) {
        delete finalData[colkey];
      }
    });

    const resp = await onSave(finalData);
    if (!resp["ok"]) {
      message = resp["data"];
    } else {
      modified_cells = {};
    }
  }}
  onDelete={async () => {
    await service.remove_row_cell(String(row.__id));
    service.close_big_modal();
  }}
  new_record={false}
>
  <svelte:fragment slot="edit">
    <EditRow {columns} {service} bind:current_cells bind:modified_cells />
  </svelte:fragment>

  <svelte:fragment slot="relation">
    <Relations {service} rid={row.__id} {gotoSiblingSheet} />
  </svelte:fragment>

  <svelte:fragment slot="history">
    <History {service} rid={row.__id} />
  </svelte:fragment>
</Layout>
