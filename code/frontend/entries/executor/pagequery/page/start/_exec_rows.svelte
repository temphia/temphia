<script lang="ts">
  import SheetInner from "../../../../portal/pages/data/sheet/_sheet_inner.svelte";
  import type {
    SheetCell,
    SheetColumn,
    SheetRow,
  } from "../../../../portal/pages/data/sheet/sheets";

  export let columns: SheetColumn[];
  export let rows: SheetRow[];
  export let cells: { [_: number]: { [_: string]: SheetCell } };
</script>

<div class="flex flex-col p-2 rounded border">
  <SheetInner {cells} {columns} {rows} editable={false} pick_label="" />
</div>
