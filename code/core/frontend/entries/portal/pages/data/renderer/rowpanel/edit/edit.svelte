<script lang="ts">
  import type { DataService } from "../../../../../services/data";
  import Field from "../../fields/field.svelte";

  export let columns: string[];
  export let columns_indexded: { [_: string]: object };
  export let row: object;
  export let rowid: number;
  export let manager: DataService;

  const _columns_indexded = columns_indexded as { [_: string]: object };


</script>

<div class="flex-grow flex flex-col h-32 p-2 space-y-1 overflow-y-auto">
  {#each columns as col, index}
    <div class="flex-col flex py-3 p-1 {index === 0 ? '' : 'border-t'}">
      <Field
        {row}
        column={_columns_indexded[col]}
        {manager}
        row_editor={null}
        onChange={(value) => {}}
      />
    </div>
  {/each}
</div>

<div class="flex-shrink h-8 border-t pt-1 pr-1 flex justify-end gap-x-2">
  <button
    on:click={() => {}}
    class="bg-red-100 hover:bg-red-600 w-14 text-red-600 text-sm hover:text-white rounded"
    >Delete</button
  >
  <button
    on:click={() => {}}
    class="bg-green-100 hover:bg-green-600 w-14 text-green-600 text-sm hover:text-white rounded"
    >Refresh</button
  >
  <button
    on:click={() => {}}
    class="bg-blue-400 hover:bg-blue-600 w-14 text-white text-sm rounded"
    >Save</button
  >
</div>
