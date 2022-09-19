<script lang="ts">
  import ViewFilter from "./_view_filter.svelte";
  import type { DynAdminAPI } from "../dtable2";

  export let table = "";
  export let group = "";
  export let source = "";
  export let id = 0;
  export let columns = [];
  export let data = {};
  export let dynapi: DynAdminAPI;

  let filter_conds = data["filter_conds"] || [];

  $: _mod_data = {};
  let filter_modified = false;
  let modified = false;

  const save = async () => {
    const _data = { ..._mod_data };
    if (filter_modified) {
      _data["filter_conds"] = filter_conds;
    }

    dynapi.modify_view(source, group, table, id, _data);
  };
</script>

<div class="mx-auto p-2 bg-white mt-10 rounded" style="max-width: 70rem;">
  <div class="flex flex-col p-2 mt-2 shadow border">
    <h2 class="inline-block text-xl text-slate-800">Name</h2>

    <div class="">
      <input
        type="text"
        bind:value={data["name"]}
        on:change={(ev) => {
          modified = true;
          _mod_data["name"] = ev.target["value"];
        }}
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
      <ViewFilter {columns} bind:filter_conds bind:filter_modified />
    </div>
  </div>

  <div class="flex flex-col p-2 mt-2 shadow border">
    <h2 class="inline-block text-xl text-slate-800">Search Text</h2>

    <div class="">
      <input
        type="text"
        value={data["search_term"] || ""}
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
    <select
      class="p-2"
      value={data["main_column"] || ""}
      on:change={(ev) => {
        modified = true;
        _mod_data["main_column"] = ev.target["value"];
      }}
    >
      <option />
      {#each columns as col}
        <option value={col.slug}>{col.name}</option>
      {/each}
    </select>
  </div>

  <div class="flex flex-col p-2 mt-2 shadow border">
    <h2 class="inline-block text-xl text-slate-800">Column Order</h2>
    <input type="text" class="border" />
  </div>

  <div class="flex flex-col p-2 mt-2 shadow border">
    <h2 class="inline-block text-xl text-slate-800">Fetch Row Count</h2>
    <input
      type="number"
      class="border"
      value={data["count"] || 0}
      on:click={(ev) => {
        modified = true;
        _mod_data["count"] = ev.target["value"];
      }}
    />
  </div>

  <div class="p-2">
    {#if filter_modified || modified}
      <button
        on:click={save}
        class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Update</button
      >
    {/if}
  </div>
</div>
