<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import MapPanel from "./_map_panel.svelte";
  import FilePanel from "./_file_panel.svelte";

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

  export let column: SheetColumn;
  export let open_column;
  export let celldata = {};

  const id = `cell-${column.sheetid}`;
  const value = celldata["value"] || "";
  const value_num = celldata["numval"] || 0;
  const color = celldata["color"];

  $: _is_open = open_column === column.__id;

  const toggle = () => {
    if (_is_open) {
      open_column = null;
    } else {
      open_column = column.__id;
    }
  };

  const close = () => {
    if (_is_open) {
      open_column = null;
    }
  };
</script>

<div
  class="py-2 border-b pl-2 border-l-4 border-l-white rounded   border-l-{color}-400"
>
  <label
    class="block mb-2 text-sm font-bold text-gray-700 uppercase relative"
    for={id}
    >{column.name}

    <div class="absolute right-0 bottom-1 cursor-pointer">
      <Icon name="color-swatch" class="h-4 w-4" />
    </div>
  </label>

  {#if column.ctype === SheetColTypeLongText}
    <textarea
      {id}
      {value}
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
    />
  {:else if column.ctype === SheetColTypeBoolean}
    <label
      class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer hover:border-blue-400"
    >
      <input {id} class="shadow" type="checkbox" />
    </label>
  {:else if column.ctype === SheetColTypeNumber}
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
      {id}
      value={Number(value_num)}
      type="number"
    />
  {:else if column.ctype === SheetColTypeDate}
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
      {id}
      {value}
      type="datetime-local"
    />
  {:else if column.ctype === SheetColTypeFile}
    <div
      class="p-1 rounded bg-gray-50 hover:border-blue-400 flex gap-1 flex-wrap"
      style="min-height: 2rem;"
    >
      {#each ["aa", "bb", "cc", "dd", "ee"] as f}
        <div class="w-10 h-10 bg-gray-500 rounded cursor-pointer" />
      {/each}
    </div>

    <div class="flex justify-end gap-1">
      {#if _is_open}
        <button on:click={close}>
          <Icon
            class="w-6 h-6 p-0.5 rounded border hover:bg-yellow-100"
            name="x"
          />
        </button>
      {:else}
        <button on:click={toggle}>
          <Icon
            class="w-6 h-6 p-0.5 rounded border hover:bg-yellow-100"
            name="photograph"
          />
        </button>
      {/if}
    </div>

    {#if _is_open}
      <div
        class="p-1 border rounded shadow h-64 mt-2 border-green-500 overflow-auto"
      >
        <FilePanel />
      </div>
    {/if}
  {:else if column.ctype === SheetColTypeRatings}
    <div class="flex p-1 gap-1">
      {#each [1, 2, 3, 4, 5] as rt}
        <button>
          <Icon name="star" class="h-5 w-5 text-gray-800" />
        </button>
      {/each}
    </div>
  {:else if column.ctype === SheetColTypeLocation}
    <div class="flex cursor-pointer justify-between">
      <div class="flex gap-1" on:click={toggle}>
        <span class="bg-yellow-100 rounded p-0.5 text-gray-600">
          Lat
          <strong class="font-semibold text-gray-700">{"13.4"}</strong>
        </span>

        <span class="bg-yellow-100 rounded p-0.5 text-gray-600">
          Long
          <strong class="font-semibold text-gray-700">{"78.71"}</strong>
        </span>
      </div>

      {#if _is_open}
        <div class="flex gap-1">
          <button>
            <Icon
              class="w-6 h-6 p-0.5 rounded border hover:bg-yellow-100"
              name="location-marker"
            />
          </button>
          <button on:click={close}>
            <Icon
              class="w-6 h-6 p-0.5 rounded border hover:bg-yellow-100"
              name="x"
            />
          </button>
        </div>
      {/if}
    </div>

    {#if _is_open}
      <div class="p-1 border rounded shadow h-64 mt-2 border-green-500">
        <MapPanel />
      </div>
    {/if}
  {:else}
    <input
      {id}
      {value}
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline hover:border-blue-400"
      type="text"
    />
  {/if}
</div>
