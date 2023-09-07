<script lang="ts">
  import type { SheetColumn } from "../../../sheets";
  import Layout from "../_layout.svelte";
  import Cell from "../../../field/cell.svelte";
  import type { SheetService } from "$lib/services/data";

  export let columns: SheetColumn[];
  export let onSave: (data) => Promise<any>;
  export let service: SheetService;

  export let open_column;
  export let dirty_data = {};

  $: console.log("@dirty_data", dirty_data);

  let message = "";
</script>

<Layout
  title="Add Row"
  {message}
  onSave={async () => {
    const resp = await onSave(dirty_data);
    if (!resp["ok"]) {
      message = resp["data"];
    } else {
      message = "";
    }
  }}
>
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
