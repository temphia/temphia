<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { createEventDispatcher } from "svelte";
  import type { FolderTktAPI } from "../../../../../lib/apiv2";
  import Point from "./field/_point.svelte";
  import {
    SheetCell,
    SheetColTypeBoolean,
    SheetColTypeDate,
    SheetColTypeFile,
    SheetColTypeLocation,
    SheetColTypeNumber,
    SheetColTypeRatings,
    SheetColumn,
    SheetCtypeIcons,
    SheetRow,
  } from "./sheets";

  export let columns: SheetColumn[];
  export let rows: SheetRow[];
  export let cells: { [_: number]: { [_: string]: SheetCell } };
  export let selected_rows = [];
  export let folder_api: FolderTktAPI;
  export let editable = true;

  const dispatch = createEventDispatcher();
</script>

<table
  class="border-collapse table-auto w-full whitespace-no-wrap bg-white table-striped relative"
>
  <thead class="text-gray-600 border-gray-200 bg-gray-100">
    <tr class="text-left">
      <th class="py-1 px-3 sticky top-0 border-b w-20 bg-gray-100"> # </th>

      {#each columns as col}
        <th
          class="sticky top-0 border-b  px-6 py-1 font-bold tracking-wider uppercase text-base text-gray-700 bg-gray-100"
        >
          <button
            class="inline-flex hover:bg-blue-200 rounded px-1"
            on:click={() => dispatch("edit_column", col)}
          >
            <Icon
              name={SheetCtypeIcons[col.ctype]}
              class="h-5 w-5 mr-1 mt-1 text-gray-500"
              solid
            />
            {col.name || `Column ${col.__id}`}
          </button>
        </th>
      {/each}

      <th class="w-10 sticky top-0 bg-gray-100">
        {#if editable}
          <button
            on:click={() => dispatch("add_column")}
            class="p-1 rounded bg-blue-500 text-white hover:bg-blue-800"
          >
            <Icon name="plus" class="w-4 h-4" />
          </button>
        {/if}
      </th>
    </tr>
  </thead>
  <tbody>
    {#each rows as row}
      {@const rowdata = cells[row.__id] || {}}

      <tr>
        <td class="border-dashed border-t border-gray-200 px-2">
          {#if editable}
            <label
              class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer"
            >
              <input
                type="checkbox"
                checked={selected_rows.includes(row.__id)}
                on:click={() => {
                  if (selected_rows.includes(row.__id)) {
                    selected_rows = selected_rows.filter((r) => r !== row.__id);
                    selected_rows = selected_rows;
                  } else {
                    selected_rows = [...selected_rows, row.__id];
                  }
                }}
                class="form-checkbox rowCheckbox focus:outline-none focus:shadow-outline"
              />
            </label>
          {/if}
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
                {:else if col.ctype === SheetColTypeLocation}
                  <Point {value} />
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
                          src={folder_api && folder_api.getFilePreviewUrl(cd)}
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
          {#if editable}
            <button
              class="underline text-blue-600"
              on:click={() => dispatch("edit_row", row)}>edit</button
            >
          {:else}
            <button
              class="underline text-blue-600"
              on:click={() => dispatch("pick_row", row)}>pick</button
            >
          {/if}
        </td>
      </tr>
    {/each}
  </tbody>
</table>
