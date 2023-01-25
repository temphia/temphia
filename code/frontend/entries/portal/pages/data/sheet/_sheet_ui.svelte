<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import ToolbarAction from "../table/core/renderer/_toolbar_action.svelte";
  import type { SheetCell, SheetColumn, SheetRow, Sheet } from "./sheets";

  export let columns: SheetColumn[];
  export let rows: SheetRow[];
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let sheets: Sheet[];
  export let active_sheet: number;
</script>

<div class="flex flex-col p-2 rounded">
  <nav class="flex flex-row  border">
    {#each sheets as sheet}
      <button
        class="text-gray-600 p-2 block hover:text-blue-500 focus:outline-none {sheet.id ===
        active_sheet
          ? 'text-blue-500 border-b-2 font-medium border-blue-500'
          : ''}">{sheet.name}</button
      >
    {/each}
  </nav>

  <div class="flex p-1 gap-1">
    <ToolbarAction onClick={() => {}} icon="refresh" name="Refresh" />
    <ToolbarAction onClick={() => {}} icon="hashtag" name="Raw" />
    <ToolbarAction onClick={() => {}} icon="calendar" name="History" />
  </div>

  <div
    class="overflow-x-auto bg-white border rounded overflow-y-auto relative p-1"
  >
    <table
      class="border-collapse table-auto w-full whitespace-no-wrap bg-white table-striped relative"
    >
      <thead class="text-gray-600 border-gray-200 bg-gray-100">
        <tr class="text-left">
          <th class="py-2 px-3 sticky top-0 border-b "> # </th>

          {#each columns as col}
            <th
              class=" sticky top-0 border-b  px-6 py-2 font-bold tracking-wider uppercase text-xs userId"
              >{col.name}</th
            >
          {/each}

          <th class="w-10">
            <button
              class="p-1 rounded bg-blue-500 text-white hover:bg-blue-800"
            >
              <Icon name="plus" class="w-4 h-4" />
            </button>
          </th>
        </tr>
      </thead>
      <tbody>
        {#each rows as row}
          <tr>
            <td class="border-dashed border-t border-gray-200 px-3">
              <label
                class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer"
              >
                <input
                  type="checkbox"
                  class="form-checkbox rowCheckbox focus:outline-none focus:shadow-outline"
                />
              </label>
            </td>

            {#each columns as col}
              <td class="border-dashed border-t border-gray-200">
                <span class="text-gray-700 px-6 py-3 flex items-center"
                  >{(cells[row.id][col.id] || {}).value || ""}</span
                >
              </td>
            {/each}

            <td> <button class="underline text-blue-600">Edit</button> </td>
          </tr>
        {/each}

        <tr>
          <td>
            <button class="p-1 rounded bg-blue-500 text-white hover:bg-blue-800">
              <Icon name="plus" class="w-4 h-4" />
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
