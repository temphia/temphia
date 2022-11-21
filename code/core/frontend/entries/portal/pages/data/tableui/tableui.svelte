<script lang="ts">
  import type { TableService } from "../../../services/data/table";
  import CardLayout from "./card/card.svelte";
  import GridLayout from "./grid/grid.svelte";

  import RowPanel from "./core/rowpanel/row.svelte";

  export let table_service: TableService;
  export let layout: string;

  const data_store = table_service.state.data_store;
  const nav_store = table_service.state.nav_store;

  let show_editor = false;
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

  {#if layout === "card"}
    <CardLayout />
  {:else}
    <GridLayout
      actions={[]}
      hooks={[]}
      bind:show_editor
      selected_rows={[]}
      {table_service}
      on:on_table_change
      on:on_change_to_card
    />
  {/if}
{/key}
