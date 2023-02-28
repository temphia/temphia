<script lang="ts">
  import { KvEditor } from "../../../../admin/core";
  import ViewItem from "./_view_item.svelte";
  export let columns;
  export let data;

  let name = data["name"];
  let filter_conds = data["filter_conds"] || [];
  let search_term = data["search_term"] || "";
  let main_column = data["main_column"] || "";
  let count = data["count"] || 0;

  export const getViewData = () => ({
    name,
    filter_conds,
    search_term,
    main_column,
    count,
  });
</script>

<div class="flex flex-col px-2 py-3 mt-2 border-b">
  <h2 class="inline-block text-lg  text-slate-800 mb-1">Name</h2>

  <input
    type="text"
    bind:value={name}
    placeholder="a word.."
    class="w-full h-10 px-4 text-sm text-gray-700 bg-white border border-gray-300 rounded-lg duration-300 focus:border-teal-500 focus:outline-none focus:ring focus:ring-primary focus:ring-opacity-40"
  />
</div>

<div class="flex flex-col px-2 py-3 text-slate-600 border-b">
  <h2 class="inline-block text-lg  text-slate-800 mb-1">Filter Conditions</h2>

  <div class="w-full p-2">
    <ViewItem {columns} bind:filter_conds />
  </div>
</div>

<div class="flex flex-col px-2 py-3 mt-2 border-b">
  <h2 class="inline-block text-lg text-slate-800 mb-1">Select Columns</h2>

  <div class="flex flex-wrap text-gray-700 gap-1">
    {#each columns as col}
      <label class="p-1 border bg-red-50 rounded">
        {col.name}
        <input type="checkbox" />
      </label>
    {/each}
  </div>
</div>

<div class="flex flex-col px-2 py-3 mt-2 border-b">
  <h2 class="inline-block text-xl text-slate-800 mb-1">Fetch Row Count</h2>
  <input type="number" class="border rounded w-20" bind:value={count} />
</div>

<div class="flex flex-col px-2 py-3 border-b">
  <h2 class="inline-block text-lg text-slate-800 mb-1">View Tags</h2>
  <KvEditor data={{}} />
</div>
