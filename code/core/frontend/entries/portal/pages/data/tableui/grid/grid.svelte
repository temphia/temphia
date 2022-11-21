<script lang="ts">
  import type { TableService } from "../../../../services/data/table";
  import Renderer from "../core/renderer/renderer.svelte";
  import RendererLayout from "../core/renderer/layout.svelte";

  export let table_service: TableService;
  export let show_editor;

  export let hooks: object[];
  export let actions: object[];
  export let selected_rows = [];

  const row_service = table_service.get_row_service();

  const data_store = table_service.state.data_store;
  const dirty_store = table_service.state.dirty_store;
  const nav_store = table_service.state.nav_store;
</script>

<RendererLayout
  {actions}
  active_table={table_service.table_slug}
  all_tables={table_service.all_tables}
  loading={false}
  {selected_rows}
  {hooks}
  on:on_change_to_card
  on:on_hook_click
  on:on_new_row
  on:on_table_change
>
  <Renderer
    columns={$data_store.column_order}
    columns_index={$data_store.indexed_column}
    main_column=""
    rows={$data_store.rows}
    rows_index={$data_store.indexed_rows}
    {selected_rows}
    on:on_hook_click={(ev) => {
      console.log(ev);
    }}
    on:on_new_row={(ev) => {
      show_editor = true;
    }}
    on:on_page_buttom={(ev) => {
      console.log(ev);
    }}
    on:on_page_top={(ev) => {
      console.log(ev);
    }}
    on:on_row_click={(ev) => {
      row_service.state.start_row_edit(ev.detail);
      show_editor = true;
    }}
    on:on_row_toggle_select={(ev) => {
      console.log(ev);
    }}
    on:on_table_change
    on:on_change_to_card
  />
</RendererLayout>
