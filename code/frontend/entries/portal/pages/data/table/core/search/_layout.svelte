<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { Column } from "../../../../../services/data";

  export let columns: Column[] = [];
  export let filters = [];

  export let count = "10";
  export let column = columns[0].slug;
  export let pattern = false;
  export let search_string = "";
  export let filter_name = "";

  export let onSubmit;
</script>

<form class="flex items-center pr-6 sm:pr-0">
  <label for="data-table-search" class="sr-only">Search</label>
  <div class="relative w-full">
    <div
      class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none"
    >
      <Icon name="search" class="h-5 w-5" />
    </div>
    <input
      type="text"
      id="data-table-search"
      bind:value={search_string}
      class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5  dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
      placeholder="Search"
      required
    />
  </div>
  <button
    type="submit"
    on:click|preventDefault={() => onSubmit()}
    class="p-2.5 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
  >
    <Icon name="search" class="h-5 w-5" />
  </button>
</form>

<div class="flex flex-wrap text-sm border-b p-1 gap-3 mb-2">
  <label class="inline-flex items-center">
    <span class="text-gray-700">Count</span>
    <select bind:value={count} class="font-medium rounded text-sm p-1 ml-1">
      <option value="10">10</option>
      <option value="50">50</option>
      <option value="100">100</option>
      <option value="200">200</option>
    </select>
  </label>

  <label class="inline-flex items-center">
    <span class="text-gray-700">Use Filter</span>

    <select
      bind:value={filter_name}
      class="font-medium rounded text-sm p-1 ml-1"
    >
      {#each filters as ft}
        <option value={ft["name"]}>{ft["name"]}</option>
      {/each}
    </select>
  </label>

  <label class="inline-flex items-center">
    <span class="text-gray-700">Column</span>
    <select bind:value={column} class="font-medium rounded text-sm p-1 ml-1">
      {#each columns as col}
        <option value={col.slug}>{col.name}</option>
      {/each}
    </select>
  </label>

  <label class="inline-flex items-center">
    <span class="text-gray-700">Pattern</span>
    <input
      class="hover:bg-gray-200 rounded-lg p-1 ml-1"
      bind:checked={pattern}
      type="checkbox"
    />
  </label>
</div>

<slot />
