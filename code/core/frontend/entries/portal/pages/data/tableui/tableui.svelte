<script lang="ts">
  import type { TableService } from "../../../services/data/table";
  import CardLayout from "./card/card.svelte";
  import GridLayout from "./grid/grid.svelte";

  import RowPanel from "./core/rowpanel/row.svelte";
  import ViewPanel from "./core/view/view.svelte";

  import { createEventDispatcher } from "svelte";

  export let table_service: TableService;
  export let layout: string;

  const row_service = table_service.get_row_service();
  const data_store = table_service.state.data_store;
  const nav_store = table_service.state.nav_store;

  let show_editor = false;
  let show_view_panel = false;
  let selected_rows = [];

  const dispatch = createEventDispatcher();

  const actions = [
    {
      name: "Refresh",
      type: "normal",
      active: false,
      icon: `<path fill-rule="evenodd" d="M4 2a1 1 0 011 1v2.101a7.002 7.002 0 0111.601 2.566 1 1 0 11-1.885.666A5.002 5.002 0 005.999 7H9a1 1 0 010 2H4a1 1 0 01-1-1V3a1 1 0 011-1zm.008 9.057a1 1 0 011.276.61A5.002 5.002 0 0014.001 13H11a1 1 0 110-2h5a1 1 0 011 1v5a1 1 0 11-2 0v-2.101a7.002 7.002 0 01-11.601-2.566 1 1 0 01.61-1.276z" clip-rule="evenodd" />`,
      action: () => {},
    },
    {
      name: "Setting",
      type: "normal",
      icon: `<path d="M5 4a1 1 0 00-2 0v7.268a2 2 0 000 3.464V16a1 1 0 102 0v-1.268a2 2 0 000-3.464V4zM11 4a1 1 0 10-2 0v1.268a2 2 0 000 3.464V16a1 1 0 102 0V8.732a2 2 0 000-3.464V4zM16 3a1 1 0 011 1v7.268a2 2 0 010 3.464V16a1 1 0 11-2 0v-1.268a2 2 0 010-3.464V4a1 1 0 011-1z" />`,
      action: () => dispatch("admin_data_table", null),
    },
    {
      name: "Share",
      type: "normal",
      active: false,
      icon: `<path d="M15 8a3 3 0 10-2.977-2.63l-4.94 2.47a3 3 0 100 4.319l4.94 2.47a3 3 0 10.895-1.789l-4.94-2.47a3.027 3.027 0 000-.74l4.94-2.47C13.456 7.68 14.19 8 15 8z" />`,
      action: () => {},
    },
    {
      name: "View",
      type: "normal",
      active: false,
      icon: `<path fill-rule="evenodd" d="M3 3a1 1 0 011-1h12a1 1 0 011 1v3a1 1 0 01-.293.707L12 11.414V15a1 1 0 01-.293.707l-2 2A1 1 0 018 17v-5.586L3.293 6.707A1 1 0 013 6V3z" clip-rule="evenodd" />`,
      action: () => {
        show_view_panel = !show_view_panel;
      },
    },

    {
      name: "Clone",
      type: "contextual",
      active: false,
      icon: ` <path d="M7 9a2 2 0 012-2h6a2 2 0 012 2v6a2 2 0 01-2 2H9a2 2 0 01-2-2V9z" /> <path d="M5 3a2 2 0 00-2 2v6a2 2 0 002 2V5h8a2 2 0 00-2-2H5z" />`,
      action: () => {},
    },
    {
      name: "Delete",
      type: "contextual",
      active: false,
      icon: `<path fill-rule="evenodd" d="M9 2a1 1 0 00-.894.553L7.382 4H4a1 1 0 000 2v10a2 2 0 002 2h8a2 2 0 002-2V6a1 1 0 100-2h-3.382l-.724-1.447A1 1 0 0011 2H9zM7 8a1 1 0 012 0v6a1 1 0 11-2 0V8zm5-1a1 1 0 00-1 1v6a1 1 0 102 0V8a1 1 0 00-1-1z" clip-rule="evenodd" />`,
      action: () => {},
    },
    {
      name: "Clear",
      type: "contextual",
      active: false,
      icon: `<path fill-rule="evenodd" d="M13.477 14.89A6 6 0 015.11 6.524l8.367 8.368zm1.414-1.414L6.524 5.11a6 6 0 018.367 8.367zM18 10a8 8 0 11-16 0 8 8 0 0116 0z" clip-rule="evenodd" />`,
      action: () => {
        selected_rows = [];
      },
    },
  ];

  // event handlers

  const on_new_row = (ev) => {
    row_service.state.start_row_edit(0);
    show_editor = true;
  };

  const on_hook_click = (ev) => {};
  const on_page_buttom = (ev) => table_service.reached_buttom();
  const on_page_top = (ev) => table_service.reached_top();
  const on_row_click = (ev) => {
    row_service.state.start_row_edit(ev.detail);
    show_editor = true;
  };

  const on_row_toggle_select = (ev) => {
    const rowid: number = ev.detail;
    if (selected_rows.includes(rowid)) {
      selected_rows = selected_rows.filter((v) => v !== rowid);
    } else {
      selected_rows = [...selected_rows, rowid];
    }
  };
</script>

{#key layout}
  <RowPanel
    bind:show_editor
    {table_service}
    columns={$data_store.column_order}
    columns_indexded={$data_store.indexed_column}
    reverse_ref_column={[]}
    rows_indexed={$data_store.indexed_rows}
  />

  <ViewPanel
    columns={Object.values($data_store.indexed_column)}
    manager={table_service}
    bind:show={show_view_panel}
  />

  {#if layout === "card"}
    <CardLayout
      {actions}
      hooks={[]}
      bind:show_editor
      {selected_rows}
      {table_service}
      on:on_change_to_grid
      on:on_table_change
      on:on_page_top={on_page_top}
      on:on_page_buttom={on_page_buttom}
      on:on_new_row={on_new_row}
      on:on_row_click={on_row_click}
    />
  {:else}
    <GridLayout
      {actions}
      hooks={[]}
      {selected_rows}
      {table_service}
      on:on_table_change
      on:on_change_to_card
      on:on_new_row={on_new_row}
      on:on_page_top={on_page_top}
      on:on_row_click={on_row_click}
      on:on_hook_click={on_hook_click}
      on:on_page_buttom={on_page_buttom}
      on:on_row_toggle_select={on_row_toggle_select}
    />
  {/if}
{/key}
