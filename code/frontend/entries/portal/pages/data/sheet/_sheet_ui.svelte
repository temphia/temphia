<script lang="ts">
  import type { FolderTktAPI } from "../../../../../lib/apiv2";

  import SheetLayout from "./_sheet_layout.svelte";
  import SheetInner from "./_sheet_inner.svelte";
  import type {
    SheetColumn,
    SheetRow,
    Sheet,
    SheetCell,
    SheetWidget,
  } from "./sheets";

  export let columns: SheetColumn[];
  export let rows: SheetRow[];
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let sheets: Sheet[];
  export let widgets: SheetWidget[];
  export let active_sheet: number;
  export let selected_rows = [];
  export let folder_api: FolderTktAPI;
  export let profile_genrator: (string) => string;
</script>

<SheetLayout
  {sheets}
  {active_sheet}
  {selected_rows}
  {widgets}
  on:action_delete_trash
  on:action_goto_history
  on:action_goto_rawtable
  on:action_refresh
  on:action_search
  on:add_row
  on:add_sheet
  on:change_sheet
  on:remove_sheet
  on:action_run_widget
>
  <SheetInner
    on:add_column
    on:edit_column
    on:edit_row
    {profile_genrator}
    {columns}
    {rows}
    {cells}
    {selected_rows}
    {folder_api}
  />
</SheetLayout>
