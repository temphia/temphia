<script lang="ts">
  import type { TableService } from "../../../../services/data/table";
  import RendererLayout from "../core/renderer/layout.svelte";
  import CardItem from "./card_item.svelte";

  export let table_service: TableService;
  export let show_editor;

  export let hooks: object[];
  export let actions: object[];
  export let selected_rows = [];

  const row_service = table_service.get_row_service();

  const data_store = table_service.state.data_store;
  const dirty_store = table_service.state.dirty_store;
  const nav_store = table_service.state.nav_store;

  const rows = [
    1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20,
  ];
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
      {#each rows as row}
        <CardItem />
      {/each}
    </div>
  </div>
</RendererLayout>
