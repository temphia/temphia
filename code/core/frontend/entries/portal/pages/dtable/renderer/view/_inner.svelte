<script lang="ts">
  import ViewFilter from "../../../admin/dtable/view/_view_filter.svelte";

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

<div class="flex flex-col p-2 mt-2 shadow border">
  <h2 class="inline-block text-xl text-slate-800">Name</h2>
  <div class="">
    <input
      type="text"
      bind:value={name}
      placeholder="a word.."
      class="w-full h-12 px-4 text-sm text-gray-700 bg-white border border-gray-300 rounded-lg duration-300 focus:border-teal-500 focus:outline-none focus:ring focus:ring-primary focus:ring-opacity-40"
    />
  </div>
</div>

<div class="flex flex-col border shadow text-slate-600">
  <div class="flex justify-between p-1">
    <h2 class="inline-block text-xl font-thin text-slate-800">
      Filter Conditions
    </h2>
  </div>

  <div class="w-full p-2">
    <ViewFilter {columns} bind:filter_conds />
  </div>
</div>

<div class="flex flex-col p-2 mt-2 shadow border">
  <h2 class="inline-block text-xl text-slate-800">Search Text</h2>

  <div class="">
    <input
      type="text"
      bind:value={search_term}
      placeholder="a word.."
      class="w-full h-12 px-4 text-sm text-gray-700 bg-white border border-gray-300 rounded-lg duration-300 focus:border-teal-500 focus:outline-none focus:ring focus:ring-primary focus:ring-opacity-40"
    />
  </div>
</div>

<div class="flex flex-col p-2 mt-2 shadow border">
  <h2 class="inline-block text-xl text-slate-800">Select Columns</h2>

  <div class="flex flex-wrap text-gray-700 gap-1">
    {#each columns as col}
      <label class="p-1 border bg-red-50 rounded">
        {col.name}
        <input type="checkbox" />
      </label>
    {/each}
  </div>
</div>

<div class="flex flex-col p-2 mt-2 shadow border">
  <h2 class="inline-block text-xl text-slate-800">Main Column</h2>
  <select class="p-2" bind:value={main_column}>
    <option />
    {#each columns as col}
      <option value={col.slug}>{col.name}</option>
    {/each}
  </select>
</div>

<div class="flex flex-col p-2 mt-2 shadow border">
  <h2 class="inline-block text-xl text-slate-800">Fetch Row Count</h2>
  <input type="number" class="border" bind:value={count} />
</div>
