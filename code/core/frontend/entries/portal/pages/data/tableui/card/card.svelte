<script lang="ts">
  import type { TableService } from "../../../../services/data/table";
  import RendererLayout from "../core/renderer/layout.svelte";
  import CardItem from "./card_item.svelte";
  import { calculate_order } from "./order";

  export let table_service: TableService;
  export let show_editor;

  export let hooks: object[];
  export let actions: object[];
  export let selected_rows = [];

  console.log("@table_service1", table_service);
  const data_store = table_service.state.data_store;
  console.log("@table_service", table_service, data_store);

  $: _data = $data_store;
  $: _order = calculate_order(_data.indexed_column, {});
</script>

<RendererLayout
  {actions}
  active_table={table_service.table_slug}
  all_tables={table_service.all_tables}
  loading={false}
  {selected_rows}
  {hooks}
  on:on_hook_click
  on:on_new_row
  on:on_table_change
  on:on_change_to_grid
  layout={"card"}
>
  <div class="flex w-full" style="height:calc(100vh - 7em);">
    <div
      class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 h-full overflow-auto w-full p-2"
    >
      {#each _data.rows as row}
        <CardItem
          columns={_data.indexed_column}
          order={_order}
          row={_data.indexed_rows[row] || {}}
        />
      {/each}
    </div>
  </div>
</RendererLayout>
