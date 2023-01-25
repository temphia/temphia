<script lang="ts">
  import {
    SheetColTypeBoolean,
    SheetColTypeDate,
    SheetColTypeFile,
    SheetColTypeLocation,
    SheetColTypeLongText,
    SheetColTypeNumber,
    SheetColTypeRatings,
    SheetColumn,
  } from "../sheets";
  export let columns: SheetColumn[];
</script>

{#each columns as col}
  {@const id = `cell-${col.sheet_id}`}
  <div class="py-2 border-b">
    <label class="block mb-2 text-sm font-bold text-gray-700 uppercase" for={id}
      >{col.name}</label
    >

    {#if col.ctype === SheetColTypeLongText}
      <textarea
        {id}
        class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
      />
    {:else if col.ctype === SheetColTypeBoolean}
      <label
        class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer hover:border-blue-400"
      >
        <input {id} class="shadow" type="checkbox" />
      </label>
    {:else if col.ctype === SheetColTypeNumber}
      <input
        class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
        {id}
        type="number"
      />
    {:else if col.ctype === SheetColTypeDate}
      <input
        class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
        {id}
        type="datetime-local"
      />
    {:else if col.ctype === SheetColTypeFile}
      <div
        class="p-1 rounded bg-gray-50 hover:border-blue-400"
        style="min-height: 2rem;"
      >
        <span />
      </div>
    {:else if col.ctype === SheetColTypeRatings}
      <div />
    {:else if col.ctype === SheetColTypeLocation}
      <div />
    {:else}
      <input
        {id}
        class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
        type="text"
      />
    {/if}
  </div>
{/each}
