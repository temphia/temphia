<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../../../services";
  import type { TableService } from "../../../../services/data/table";
  import { action_builder } from "../../renderer/actions";
  import Row from "../../renderer/rowpanel/row.svelte";
  import Datatable from "./datatable/datatable.svelte";

  const app: PortalService = getContext("__app__");

  export let table_service: TableService;

  const row_service = table_service.get_row_service();

  const data_store = table_service.state.data_store;
  const nav_store = table_service.state.nav_store;

  let abuilder = action_builder(table_service, app, []);

  let show_editor = false;
</script>

<Row
  bind:show_editor
  {table_service}
  columns={$data_store.column_order}
  columns_indexded={$data_store.indexed_column}
  reverse_ref_column={[]}
  rows_indexed={$data_store.indexed_rows}
/>

<Datatable
  actions={abuilder.actions}
  active_table={table_service.table_slug}
  all_tables={table_service.all_tables}
  columns={$data_store.column_order}
  columns_index={$data_store.indexed_column}
  hooks={[]}
  main_column=""
  loading={$nav_store.loading}
  rows={$data_store.rows}
  rows_index={$data_store.indexed_rows}
  selectedRows={[]}
  on:on_hook_click={(ev) => {}}
  on:on_new_row={(ev) => {
    show_editor = true;
  }}
  on:on_page_buttom={(ev) => {}}
  on:on_page_top={(ev) => {}}
  on:on_row_click={(ev) => {
    row_service.state.start_row_edit(ev.detail);
    show_editor = true;
  }}
  on:on_row_toggle_select={(ev) => {}}
  on:on_table_change={(ev) => {}}
/>
