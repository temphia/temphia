<script lang="ts">
  import { DataTable } from "../../../../../components";

  import View from "./view/view.svelte";

  import Row from "./rowpanel/row.svelte";
  import Preview from "./preview.svelte";
  import DataPlugin from "./data_plugin/data_plugin.svelte";
  import type { DataTableService } from "../../../../../lib/service/dyn";
  import type { Navigator } from "../../../../../lib/app/portal/navigator";


  export let source: string;
  export let dgroup: string;
  export let dtable: string;
  export let manager: DataTableService;
  export let navigator: Navigator;

  let states = manager.store.states;
  $: _state = $states[dtable];
  $: _columns = _state.column_order;
  $: _columns_indexed = _state.indexed_column;
  $: _rows = _state.rows;
  $: _rows_indexed = _state.indexed_rows;
  $: _reverse_ref_column = _state.reverse_ref_column;

  $: _show_editor = false;
  $: _show_data_plugin = false;

  $: _nav_data = manager.navStore;

  const newRowClick = () => {
    manager.row_editor.startNewRow();
    _show_editor = true;
  };

  $: _selected_rows = [];

  const rowToggleSelect = (rowid) => () => {
    if (_selected_rows.includes(rowid)) {
      _selected_rows = _selected_rows.filter((v) => v !== rowid);
    } else {
      _selected_rows = [..._selected_rows, rowid];
    }
  };

  $: _show_view = false;

  $: _actions = [
    {
      name: "Refresh",
      type: "normal",
      active: false,
      icon: `<svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd" /></svg>`,
      action: manager.init,
    },
    {
      name: "Setting",
      type: "normal",
      icon: `<svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path d="M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z" /></svg>`,
      action: () => {
        navigator.goto_admin_dtable_page(source, dgroup, dtable);
      },
    },
    {
      name: "Share",
      type: "normal",
      active: false,
      icon: `<svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z" /></svg>`,
      action: () => {},
    },
    {
      name: "View",
      type: "normal",
      active: false,
      icon: `<svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z" clip-rule="evenodd" /></svg>`,
      action: () => {
        _show_view = !_show_view;
      },
    },

    {
      name: "Clone",
      type: "contextual",
      active: false,
      icon: `<svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"> <path d="M7 9a2 2 0 012-2h6a2 2 0 012 2v6a2 2 0 01-2 2H9a2 2 0 01-2-2V9z" /> <path d="M5 3a2 2 0 00-2 2v6a2 2 0 002 2V5h8a2 2 0 00-2-2H5z" /></svg>`,
      action: () => {},
    },
    {
      name: "Delete",
      type: "contextual",
      active: false,
      icon: `<svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" /></svg>`,
      action: () => {},
    },
    {
      name: "Clear",
      type: "contextual",
      active: false,
      icon: `<svg class="h-5 w-5" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M13.477 14.89A6 6 0 015.11 6.524l8.367 8.368zm1.414-1.414L6.524 5.11a6 6 0 018.367 8.367zM18 10a8 8 0 11-16 0 8 8 0 0116 0z" clip-rule="evenodd" /></svg>`,
      action: () => (_selected_rows = []),
    },
  ];

  let plugin_target;

  const get_ref = () => {
    return plugin_target;
  };

  manager.set_ref_callback(get_ref);
</script>

<Row
  bind:show_editor={_show_editor}
  {manager}
  columns={_columns}
  columns_indexded={_columns_indexed}
  rows_indexed={_rows_indexed}
  reverse_ref_column={_reverse_ref_column}
/>

<DataPlugin bind:show={_show_data_plugin} bind:plugin_target />

<View
  columns={Object.values(_columns_indexed)}
  {manager}
  bind:show={_show_view}
/>

<DataTable
  columns={_columns}
  columns_index={_columns_indexed}
  hooks={_state.hooks}
  rows={_rows}
  rows_index={_rows_indexed}
  onPageButtom={manager.reachedButtom}
  onPageTop={manager.reachedTop}
  onHookClick={(hook) => {
    _show_data_plugin = true;
    manager.hook_executor.execute_hook(hook);
  }}
  loading={$_nav_data.loading}
  actions={_actions}
  all_tables={manager.groupOpts.tables}
  active_table={dtable}
  onChangeDtable={(nexttbl) => {
    navigator.goto_dtable(source, dgroup, nexttbl);
  }}
  main_column=""
  {newRowClick}
  rowClick={(rowid) => {
    console.log("@@@@@@@ START EDITING ", rowid);
    manager.row_editor.startModifyRow(rowid);
    _show_editor = true;
  }}
  selectedRows={_selected_rows}
  {rowToggleSelect}
  let:row
  let:column
>
  <svelte:fragment slot="cell">
    <Preview
      column={_columns_indexed[column] || {}}
      row={_rows_indexed[row] || {}}
      folder_api={manager.FolderTktAPI}
    />
  </svelte:fragment>
</DataTable>
