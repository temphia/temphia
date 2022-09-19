<script lang="ts">
  import Activity from "./activity/activity.svelte";

  import Edit from "./edit/edit.svelte";
  import Properties from "./properties/properties.svelte";
  import Rowlayout from "./_layout.svelte";
  import Relations from "./relations/relations.svelte";
  import type { DataTableService } from "../../../../app";

  export let show_editor = false;
  export let manager: DataTableService;
  export let rows_indexed: { [_: number]: object };
  export let columns: string[];
  export let columns_indexded: { [_: string]: object };
  export let reverse_ref_column: object[];

  $: _dirty_store = manager.dirtyStore;
  $: _rowid = $_dirty_store.rowid || 0;
  $: _row = rows_indexed[_rowid] || {};
</script>

{#key _rowid}
  <Rowlayout bind:show_editor row_id={_rowid}>
    <svelte:fragment slot="edit">
      <Edit rowid={_rowid} row={_row} {columns} {columns_indexded} {manager} />
    </svelte:fragment>

    <svelte:fragment slot="activity">
      <Activity rowid={_rowid} {manager} />
    </svelte:fragment>

    <svelte:fragment slot="meta">
      <Properties row={_row} />
    </svelte:fragment>

    <svelte:fragment slot="relations">
      <Relations {reverse_ref_column} {manager} row={_row} />
    </svelte:fragment>
  </Rowlayout>
{/key}
