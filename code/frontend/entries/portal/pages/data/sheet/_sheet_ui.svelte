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
          class="p-2 block hover:text-blue-500 focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500"
          >{sheet.name}</button
        >
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
    class="overflow-x-auto bg-white border rounded overflow-y-auto relative p-1"
  >
    <table
      class="border-collapse table-auto w-full whitespace-no-wrap bg-white table-striped relative"
    >
      <thead class="text-gray-600 border-gray-200 bg-gray-100">
        <tr class="text-left">
          <th class="py-2 px-3 sticky top-0 border-b w-20"> # </th>

          {#each columns as col}
            <th
              class="sticky top-0 border-b  px-6 py-2 font-bold tracking-wider uppercase text-xs text-gray-700"
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

          <th class="w-10">
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
              {@const celldata = rowdata[col.__id]}
              {@const color = (celldata || {})["color"] || ""}
              <td
                class="border-dashed border-t border-gray-200 bg-{color}-400"
                style="background-color: {color};"
              >
                {#if celldata}
                  <span class="text-gray-700 px-6 py-3 flex items-center">
                    {#if col.ctype === SheetColTypeBoolean}
                      {#if celldata["value"] === "true"}
                        <Icon name="check" class="w-6 h-6 text-green-500" />
                      {:else if celldata["value"] === "false"}
                        <Icon name="x" class="w-6 h-6 text-red-500" />
                      {/if}
                    {:else if col.ctype === SheetColTypeDate}
                      {new Date(celldata.value).toLocaleDateString()}
                    {:else if col.ctype === SheetColTypeFile}
                      {#if celldata["value"]}
                        {#each celldata["value"].split(",") as cd}
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
                    {:else}
                      {celldata["value"] || ""}
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

        <tr>
          <td>
            <button
              on:click={() => dispatch("add_row")}
              class="p-1 rounded bg-blue-500 text-white hover:bg-blue-800"
            >
              <Icon name="plus" class="w-4 h-4" />
            </button>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
