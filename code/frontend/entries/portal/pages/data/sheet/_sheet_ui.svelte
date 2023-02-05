<script lang="ts">
  import type { FolderTktAPI } from "../../../../../lib/apiv2";

  import SheetLayout from "./_sheet_layout.svelte";
  import SheetInner from "./_sheet_inner.svelte";
  import type { SheetColumn, SheetRow, Sheet, SheetCell } from "./sheets";

  export let columns: SheetColumn[];
  export let rows: SheetRow[];
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let sheets: Sheet[];
  export let active_sheet: number;
  export let selected_rows = [];
  export let folder_api: FolderTktAPI;
</script>

<SheetLayout
  {sheets}
  {active_sheet}
  {selected_rows}
  on:action_delete_trash
  on:action_goto_history
  on:action_goto_rawtable
  on:action_refresh
  on:add_row
  on:add_sheet
  on:change_sheet
  on:remove_sheet
>
  <SheetInner
    on:add_column
    on:edit_column
    on:edit_row
    {columns}
    {rows}
    {cells}
    {selected_rows}
    {folder_api}
  />
</SheetLayout>
