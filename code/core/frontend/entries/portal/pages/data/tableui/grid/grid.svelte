<script lang="ts">
  import type { TableService } from "../../../../services/data/table";
  import Renderer from "../core/renderer/renderer.svelte";
  import RendererLayout from "../core/renderer/layout.svelte";

  export let table_service: TableService;
  export let hooks: object[];
  export let actions: object[];
  export let selected_rows = [];

  const data_store = table_service.state.data_store;
  const nav_store = table_service.state.nav_store;

  $: _data = $data_store;
  $: _nav_store = $nav_store;
</script>

<RendererLayout
  {actions}
  active_table={table_service.table_slug}
  all_tables={table_service.all_tables}
  loading={_nav_store.loading}
  {selected_rows}
  {hooks}
  on:on_change_to_card
  on:on_hook_click
  on:on_new_row
  on:on_table_change
  rows_total_no={0}
  rows_loaded_no={_data.rows.length || 0}
>
  <Renderer
    columns={_data.column_order}
    columns_index={_data.indexed_column}
    main_column={_nav_store.active_view.main_column}
    rows={_data.rows}
    rows_index={_data.indexed_rows}
    {selected_rows}
    on:on_new_row
    on:on_page_buttom
    on:on_page_top
    on:on_row_click
    on:on_row_toggle_select
    on:on_table_change
    on:on_change_to_card
  />
</RendererLayout>
