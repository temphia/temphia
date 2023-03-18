<script lang="ts">
  import {
    FilterLessThan,
    KeyPrimary,
    TableService,
  } from "../../../services/data";
  import CardLayout from "./card/card.svelte";
  import GridLayout from "./grid/grid.svelte";

  import RowPanel from "./core/rowpanel/row.svelte";
  import ViewPanel from "./core/view/view.svelte";

  import { createEventDispatcher } from "svelte";
  import type { ViewModal } from "../table/core/view/view";
  import { get } from "svelte/store";
  import Search from "./core/search/search.svelte";

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

  const cloneRow = (rowId: number) => {
    console.log("@clone_row", rowId);

    const ds = get(data_store);
    const rowdata = { ...(ds.indexed_rows[rowId] || {}) };

    delete rowdata[KeyPrimary];

    row_service.state.start_row_edit(0, rowdata);
    show_editor = true;
  };

  const deleteRows = () => {
    // batch row delete support
    row_service.delete_row(selected_rows[0]);
    selected_rows = [];
  };

  const on_search = () => {
    table_service._open_modal(Search, {
      table_service,
      columns: Object.values($data_store.indexed_column),
    });
  };

  const on_goto = () => {
    let rowId = Number(prompt("Enter row id to goto", "20"));
    if (!rowId) {
      return;
    }

    table_service.init([
      {
        column: KeyPrimary,
        cond: FilterLessThan,
        value: `${rowId - 10}`,
      },
    ]);
  };
</script>

{#key layout}
  <RowPanel
    bind:show_editor
    {table_service}
    columns={$data_store.column_order}
    columns_indexded={$data_store.indexed_column}
    reverse_ref_column={table_service.rev_ref_columns}
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
      data_widgets={table_service.data_widgets}
      bind:show_editor
      {selected_rows}
      {table_service}
      needs_refresh={$nav_store.needs_refresh}
      on:on_change_to_grid
      on:on_table_change
      on:on_page_top={on_page_top}
      on:on_page_buttom={on_page_buttom}
      on:on_new_row={on_new_row}
      on:on_row_click={on_row_click}
      on:tb_clear={() => {
        selected_rows = [];
      }}
      on:tb_clone={() => cloneRow(selected_rows[0])}
      on:tb_delete={deleteRows}
      on:tb_execute_widget
      on:tb_history={() => dispatch("goto_history")}
      on:tb_search={on_search}
      on:tb_goto={on_goto}
      on:tb_view={() => {
        show_view_panel = !show_view_panel;
      }}
      on:tb_goto_setting={() => dispatch("admin_data_table")}
    />
  {:else}
    <GridLayout
      data_widgets={table_service.data_widgets}
      {selected_rows}
      {table_service}
      needs_refresh={$nav_store.needs_refresh}
      on:tb_reload={table_service.refresh}
      on:on_table_change
      on:on_change_to_card
      on:on_new_row={on_new_row}
      on:on_page_top={on_page_top}
      on:on_row_click={on_row_click}
      on:on_hook_click={on_hook_click}
      on:on_page_buttom={on_page_buttom}
      on:on_row_toggle_select={on_row_toggle_select}
      on:tb_clear={() => {
        selected_rows = [];
      }}
      on:tb_clone={() => cloneRow(selected_rows[0])}
      on:tb_delete={deleteRows}
      on:tb_execute_widget
      on:tb_history={() => dispatch("goto_history")}
      on:tb_search={on_search}
      on:tb_goto={on_goto}
      on:tb_view={() => {
        show_view_panel = !show_view_panel;
      }}
      on:tb_goto_setting={() => dispatch("admin_data_table", null)}
    />
  {/if}
{/key}
