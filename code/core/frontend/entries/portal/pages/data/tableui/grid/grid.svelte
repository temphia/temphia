<script lang="ts">
  import type { TableService } from "../../../../services/data/table";
  import Renderer from "../core/renderer/renderer.svelte";
  import RendererLayout from "../core/renderer/layout.svelte";

  export let table_service: TableService;
  export let hooks: object[];
  export let actions: object[];
  export let selected_rows = [];

  const data_store = table_service.state.data_store;

  $: _data = $data_store;
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
    columns={_data.column_order}
    columns_index={_data.indexed_column}
    main_column=""
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
