<script lang="ts">
  import type { SheetColumn } from "../../sheets";
  import Layout from "./_layout.svelte";
  import Cell from "../../field/cell.svelte";
  import type { SheetService } from "../../../../../services/data";

  export let columns: SheetColumn[];
  export let onSave = async (data) => {};
  export let service: SheetService;

  export let open_column;
  export let dirty_data = {};

  $: console.log("@dirty_data", dirty_data);
</script>

<Layout title="Add Row" onSave={async () => onSave(dirty_data)}>
  {#each columns as col}
    <Cell
      {service}
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
