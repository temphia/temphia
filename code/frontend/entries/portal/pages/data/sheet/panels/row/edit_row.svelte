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

  let open_column;

  let current_cells = cells[row.__id] || {};
</script>

<Layout
  title="Edit Row"
  onSave={() => onSave(current_cells)}
  new_record={false}
>
  <svelte:fragment slot="edit">
    <EditRow {columns} {service} bind:current_cells />
  </svelte:fragment>

  <svelte:fragment slot="relation">
    <Relations {service} rid={row.__id} />
  </svelte:fragment>

  <svelte:fragment slot="history">
    <History {service}  rid={row.__id} />
  </svelte:fragment>
</Layout>
