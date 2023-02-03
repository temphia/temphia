<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { createEventDispatcher } from "svelte";
  import type { FolderTktAPI } from "../../../../../lib/apiv2";
  import ToolbarAction from "../table/core/renderer/_toolbar_action.svelte";

  import {
    SheetCell,
    SheetColumn,
    SheetRow,
    Sheet,
    SheetColTypeBoolean,
    SheetColTypeDate,
    SheetCtypeIcons,
    SheetColTypeFile,
    SheetColTypeRatings,
    SheetColTypeNumber,
  } from "./sheets";

  export let columns: SheetColumn[];
  export let rows: SheetRow[];
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let sheets: Sheet[];
  export let active_sheet: number;
  export let selected_rows = [];
  export let folder_api: FolderTktAPI;

  const dispatch = createEventDispatcher();
</script>

<div class="flex flex-col p-2 rounded">
  <nav class="flex flex-row  border flex-nowrap overflow-auto">
    {#each sheets as sheet}
      {#if sheet.__id === active_sheet}
        <button
          class="p-2 hover:text-blue-500 focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500 inline-flex"
          >{sheet.name}
          &nbsp;&nbsp;
          <span on:click={() => dispatch("remove_sheet")}>
            <Icon
              name="x-circle"
              class="w-5 pt-1 text-gray-500 hover:text-red-500"
            />
          </span>
        </button>
      {:else}
        <button
          on:click={() => dispatch("change_sheet", sheet.__id)}
          class="text-gray-600 p-2 block hover:text-blue-500 focus:outline-none"
          >{sheet.name}</button
        >
      {/if}
    {/each}

    <button
      on:click={() => dispatch("add_sheet")}
      class="m-2 p-1 rounded hover:bg-blue-200 border border-blue-200"
    >
      <Icon name="plus" class="w-4 h-4" />
    </button>
  </nav>

  <div class="flex p-1 gap-1">
    <ToolbarAction
      onClick={() => dispatch("action_refresh")}
      icon="refresh"
      name="Refresh"
    />
    <ToolbarAction
      onClick={() => dispatch("action_goto_rawtable")}
      icon="hashtag"
      name="Raw"
    />
    <ToolbarAction
      onClick={() => dispatch("action_goto_history")}
      icon="calendar"
      name="History"
    />

    {#if selected_rows.length === 1}
      <ToolbarAction
        onClick={() => dispatch("action_delete_trash")}
        icon="trash"
        name="Delete"
      />
    {/if}
  </div>

  <div
    id="sheet-main"
    class="overflow-x-auto bg-white border rounded overflow-y-auto relative"
  >
    <table
      class="border-collapse table-auto w-full whitespace-no-wrap bg-white table-striped relative"
    >
      <thead class="text-gray-600 border-gray-200 bg-gray-100">
        <tr class="text-left">
          <th class="py-2 px-3 sticky top-0 border-b w-20 bg-gray-100"> # </th>

          {#each columns as col}
            <th
              class="sticky top-0 border-b  px-6 py-2 font-bold tracking-wider uppercase text-xs text-gray-700 bg-gray-100"
            >
              <span class="inline-flex">
                <Icon
                  name={SheetCtypeIcons[col.ctype]}
                  class="h-5 w-5 mr-1 text-gray-500"
                  solid
                />
                {col.name || `Column ${col.__id}`}
              </span>
            </th>
          {/each}

          <th class="w-10 sticky top-0 bg-gray-100">
            <button
              on:click={() => dispatch("add_column")}
              class="p-1 rounded bg-blue-500 text-white hover:bg-blue-800"
            >
              <Icon name="plus" class="w-4 h-4" />
            </button>
          </th>
        </tr>
      </thead>
      <tbody>
        {#each rows as row}
          {@const rowdata = cells[row.__id] || {}}

          <tr>
            <td class="border-dashed border-t border-gray-200 px-2">
              <label
                class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer"
              >
                <input
                  type="checkbox"
                  checked={selected_rows.includes(row.__id)}
                  on:click={() => {
                    if (selected_rows.includes(row.__id)) {
                      selected_rows = selected_rows.filter(
                        (r) => r !== row.__id
                      );
                      selected_rows = selected_rows;
                    } else {
                      selected_rows = [...selected_rows, row.__id];
                    }
                  }}
                  class="form-checkbox rowCheckbox focus:outline-none focus:shadow-outline"
                />
              </label>
              <span class="text-xs text-gray-500">{row.__id || ""}</span>
            </td>

            {#each columns as col}
              {@const hasCellData = !!rowdata[col.__id]}
              {@const celldata = rowdata[col.__id] || {}}
              {@const color = celldata["color"] || ""}
              {@const value = celldata["value"] || ""}
              {@const num_value = celldata["numval"] || 0}
              <td
                class="border-dashed border-t border-gray-200 bg-{color}-400"
                style="background-color: {color};"
              >
                {#if hasCellData}
                  <span class="text-gray-700 px-6 py-3 flex items-center">
                    {#if col.ctype === SheetColTypeBoolean}
                      {#if value === "true"}
                        <Icon name="check" class="w-6 h-6 text-green-500" />
                      {:else if value === "false"}
                        <Icon name="x" class="w-6 h-6 text-red-500" />
                      {/if}
                    {:else if col.ctype === SheetColTypeDate}
                      {value && new Date(value).toLocaleDateString()}
                    {:else if col.ctype === SheetColTypeRatings}
                      {#if num_value}
                        {#each [1, 2, 3, 4, 5] as rt}
                          {#if rt <= num_value}
                            <Icon
                              name="star"
                              class="h-5 w-5 text-yellow-400 "
                              solid={true}
                            />
                          {/if}
                        {/each}
                      {/if}
                    {:else if col.ctype === SheetColTypeFile}
                      {#if value}
                        {#each value.split(",") as cd}
                          <div class="flex gap-1">
                            <img
                              class="h-8 w-auto border rounded"
                              src={folder_api &&
                                folder_api.getFilePreviewUrl(cd)}
                              alt=""
                            />
                          </div>
                        {/each}
                      {/if}
                    {:else if col.ctype === SheetColTypeNumber}
                      {num_value}
                    {:else}
                      {value}
                    {/if}
                  </span>
                {/if}
              </td>
            {/each}

            <td>
              <button
                class="underline text-blue-600"
                on:click={() => dispatch("edit_row", row)}>edit</button
              >
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>

<div class="fixed bottom-4 z-5 right-14">
  <button
    on:click={() => dispatch("add_row")}
    class="p-0 w-8 h-8 bg-blue-400 rounded-full hover:bg-blue-700 active:shadow-lg mouse shadow transition ease-in duration-200 focus:outline-none"
  >
    <Icon name="plus" class="w-6 h-6 inline-block text-white" />
  </button>
</div>

<style>
  #sheet-main {
    height: calc(-7rem + 100vh);
  }

  tr {
    max-height: 50rem;
  }
</style>
