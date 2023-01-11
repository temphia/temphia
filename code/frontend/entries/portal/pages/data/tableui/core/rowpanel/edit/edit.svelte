<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { TableService, Column } from "../../../../../../services/data";
  import Field from "../../fields/field.svelte";

  export let columns: string[];
  export let columns_indexded: { [_: string]: object };
  export let row: object;
  export let rowid: number;
  export let table_service: TableService;

  const rservice = table_service.get_row_service();
  const dservice = rservice.get_dirty_service();
  const _columns_indexded = columns_indexded as { [_: string]: Column };
</script>

<div class="flex-grow flex flex-col h-32 p-2 space-y-1 overflow-y-auto">
  {#each columns as col, index}
    <div class="flex-col flex py-3 p-1 {index === 0 ? '' : 'border-t'}">
      <Field
        {row}
        column={_columns_indexded[col]}
        row_service={table_service.get_row_service()}
        onChange={(value) => dservice.on_ohange(col, value)}
      />
    </div>
  {/each}
</div>

<div class="flex-shrink h-10 border-t p-1 flex justify-end gap-x-2">
  <button
    on:click={() => rservice.delete_row(rowid)}
    class="bg-red-100 hover:bg-red-600 text-red-600 hover:text-white rounded inline-flex p-1"
  >
    <Icon name="trash" class="h-6 w-6 pt-1" />
    Delete</button
  >
  <button
    on:click={() => rservice.fetch_row_latest(rowid)}
    class="bg-green-100 hover:bg-green-600 text-green-600 hover:text-white rounded inline-flex p-1"
    >
    <Icon name="refresh" class="h-6 w-6 pt-1" />

    Refresh</button
  >
  <button
    on:click={rservice.save_row}
    
    class="bg-blue-400 hover:bg-blue-600 text-white rounded inline-flex p-1"
    >
    <Icon name="save" class="h-6 w-6 pt-1" />
    Save</button
  >
</div>
