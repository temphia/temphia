<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import {
    calculate_card_order,
    TableService,
  } from "../../../../services/data/table";

  import RendererLayout from "../core/layout.svelte";
  import CardItem from "./_item.svelte";

  export let table_service: TableService;
  export let show_editor;
  export let data_widgets: object[];
  export let selected_rows = [];
  export let needs_refresh = false;

  const dispatch = createEventDispatcher();
  const onPageButtom = () => dispatch("on_page_buttom");
  const onPageTop = () => dispatch("on_page_top");
  const rowClick = (payload) => dispatch("on_row_click", payload);

  const folder_api = table_service.get_row_service().folder_api();
  const data_store = table_service.state.data_store;
  const nav_store = table_service.state.nav_store;

  $: _data = $data_store;
  $: _nav = $nav_store;
  $: _order = calculate_card_order(_data.indexed_column, {});

  let container;

  const on_scroll = (ev) => {
    const { scrollTop, scrollTopMax } = container;
    if (scrollTop == 0) {
      onPageTop();
    } else if (scrollTopMax === scrollTop) {
      onPageButtom();
    }
  };
</script>

<RendererLayout
  on:tb_clear
  on:tb_clone
  on:tb_delete
  on:tb_execute_widget
  on:tb_history
  on:tb_view
  on:tb_goto
  on:tb_search
  on:tb_goto_setting
  active_table={table_service.table_slug}
  all_tables={table_service.all_tables}
  loading={false}
  {selected_rows}
  {data_widgets}
  viewmode={_nav.view_mode}
  on:on_hook_click
  on:on_new_row
  on:on_table_change
  on:on_change_to_grid
  rows_total_no={0}
  rows_loaded_no={_data.rows.length || 0}
  {needs_refresh}
  layout={"card"}
>
  <div class="flex w-full" style="height:calc(100vh - 7em);">
    <div
      bind:this={container}
      class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 gap-6 h-full overflow-auto w-full p-2"
      on:scroll={on_scroll}
    >
      {#each _data.rows as row}
        <CardItem
          columns={_data.indexed_column}
          order={_order}
          row={_data.indexed_rows[row] || {}}
          onEdit={() => rowClick(row)}
          {folder_api}
        />
      {/each}
    </div>
  </div>
</RendererLayout>
