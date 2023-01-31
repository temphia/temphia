<script lang="ts">
  import type { SheetColumn } from "../sheets";
  import Layout from "./_layout.svelte";
  import Cell from "../field/cell.svelte";

  export let columns: SheetColumn[];
  export let onSave = async (data) => {};

  export let open_column;

  let dirty_data = {};
</script>

<Layout title="Add Row" onClick={async () => onSave(dirty_data)}>
  {#each columns as col}
    <Cell
      column={col}
      bind:open_column
      celldata={dirty_data[col.__id]}
      onCellChange={(data) => {
        const old = dirty_data[col.__id] || {};
        dirty_data[col.__id] = { ...old, ...data };
      }}
    />
  {/each}
</Layout>
