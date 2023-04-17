<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { SheetService } from "../../../../services/data";
  import { formatRefCells } from "../../../../services/data/sheet/format";
  import { LoadingSpinner } from "../../../admin/core";
  import SheetInner from "../_sheet_inner.svelte";

  export let service: SheetService;

  let searchstring = "";
  let loading = false;

  let data;

  const load = async () => {
    loading = true;

    const resp = await service.search(searchstring);
    if (!resp.ok) {
      console.log("@resp", resp);
      return;
    }

    data = formatRefCells({ ...resp.data, sheet_id: service.sheetid });
    loading = false;
  };
</script>

<form class="flex items-center">
  <label for="simple-search" class="sr-only">Search</label>
  <div class="relative w-full">
    <div
      class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none"
    >
      <Icon name="search" class="h-5 w-5" />
    </div>
    <input
      type="text"
      id="simple-search"
      class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full pl-10 p-2.5 dark:bg-gray-700 dark:border-gray-600 dark:placeholder-gray-400 dark:text-white dark:focus:ring-blue-500 dark:focus:border-blue-500"
      placeholder="Search"
      bind:value={searchstring}
      required
    />
  </div>
  <button
    type="submit"
    on:click|preventDefault={() => load()}
    class="p-2.5 ml-2 text-sm font-medium text-white bg-blue-700 rounded-lg border border-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800"
  >
    <Icon name="search" class="h-5 w-5" />
  </button>
</form>

<div class="flex flex-wrap justify-between text-sm">
  <div class="flex flex-wrap text-sm p-1 gap-1">
    <label class="inline-flex items-center">
      <span class="ml-2 text-gray-700">Count</span>
      <select class="font-medium rounded text-sm p-1">
        <option value="10">10</option>
        <option value="50">50</option>
        <option value="100">100</option>
      </select>
    </label>
  </div>
</div>

{#if loading}
  <LoadingSpinner classes="" />
{:else if data}
  <div class="p-1 border rounded shadow overflow-auto">
    <SheetInner
      editable={false}
      cells={data["cells"] || {}}
      columns={data["columns"] || []}
      rows={data["rows"] || []}
      selected_rows={[]}
      pick_label="goto"
      on:pick_row={(ev) => {
        service.goto_row(ev.detail["__id"]);
        service.close_big_modal()
      }}
    />
  </div>
{/if}
