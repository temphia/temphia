<script lang="ts">
  import type { TableService } from "../../../services/data/table";
  import CardLayout from "./card/card.svelte";
  import GridLayout from "./grid/grid.svelte";

  import RowPanel from "./core/rowpanel/row.svelte";
  import ViewPanel from "./core/view/view.svelte";

  import { createEventDispatcher } from "svelte";
  import type { ViewModal } from "./core/view/view";

  export let table_service: TableService;
  export let layout: string;
  export let view_modal: ViewModal;

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
      name: "History",
      type: "normal",
      active: false,
      icon: `<path d="M8.515 1.019A7 7 0 0 0 8 1V0a8 8 0 0 1 .589.022l-.074.997zm2.004.45a7.003 7.003 0 0 0-.985-.299l.219-.976c.383.086.76.2 1.126.342l-.36.933zm1.37.71a7.01 7.01 0 0 0-.439-.27l.493-.87a8.025 8.025 0 0 1 .979.654l-.615.789a6.996 6.996 0 0 0-.418-.302zm1.834 1.79a6.99 6.99 0 0 0-.653-.796l.724-.69c.27.285.52.59.747.91l-.818.576zm.744 1.352a7.08 7.08 0 0 0-.214-.468l.893-.45a7.976 7.976 0 0 1 .45 1.088l-.95.313a7.023 7.023 0 0 0-.179-.483zm.53 2.507a6.991 6.991 0 0 0-.1-1.025l.985-.17c.067.386.106.778.116 1.17l-1 .025zm-.131 1.538c.033-.17.06-.339.081-.51l.993.123a7.957 7.957 0 0 1-.23 1.155l-.964-.267c.046-.165.086-.332.12-.501zm-.952 2.379c.184-.29.346-.594.486-.908l.914.405c-.16.36-.345.706-.555 1.038l-.845-.535zm-.964 1.205c.122-.122.239-.248.35-.378l.758.653a8.073 8.073 0 0 1-.401.432l-.707-.707z"/><path d="M8 1a7 7 0 1 0 4.95 11.95l.707.707A8.001 8.001 0 1 1 8 0v1z"/><path d="M7.5 3a.5.5 0 0 1 .5.5v5.21l3.248 1.856a.5.5 0 0 1-.496.868l-3.5-2A.5.5 0 0 1 7 9V3.5a.5.5 0 0 1 .5-.5z"/>`,
      action: () => dispatch("goto_history"),
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
    reverse_ref_column={$data_store.reverse_ref_column}
    rows_indexed={$data_store.indexed_rows}
    onReverseFollow={(stable, scolumn, filter_opts) => {
      console.log("@reverseFollow", stable, scolumn, filter_opts);
    }}
  />

  <ViewPanel
    columns={Object.values($data_store.indexed_column)}
    manager={table_service}
    bind:show={show_view_panel}
    {view_modal}
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
