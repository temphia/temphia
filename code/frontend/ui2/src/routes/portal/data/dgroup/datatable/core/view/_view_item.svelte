<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { Column } from "../../../../../services/data";
  import * as f from "../fields/field";
  import { CtypeFilterConds } from "./filters";
  
  export let columns: Column[] = [];
  export let filter_conds: {
    column: string;
    cond: string;
    value: string;
  }[] = [];

  export let onModify  = (conds) => {}

  export let filter_modified = false;
  const colindexed = columns.reduce((acc, curr) => {
    acc[curr.slug] = curr;
    return acc;
  }, {});

  $: _new_column_slug = "";
  $: _new_column_cond = "";
  $: _current_column = colindexed[_new_column_slug] || {};
  $: _new_column_type = _current_column.ctype || f.CtypeShortText;

  $: _possible_cond = CtypeFilterConds[_new_column_type] || {};
  $: _new_filter_value = undefined;
  $: _filer_component = _possible_cond[_new_column_cond];

  $: console.log("@filter_component", _filer_component);
  $: console.log("@possible_cond", _possible_cond);
  $: console.log("@new_column_type", _new_column_type);
  $: console.log("@new_filter_value", _new_column_type);

  let onAdd = () => {
    filter_conds = [
      ...filter_conds,
      {
        column: _new_column_slug,
        cond: _new_column_cond,
        value:
          typeof _new_filter_value === "undefined" ? "" : _new_filter_value,
      },
    ];

    _new_column_slug = "";
    _new_column_cond = "";
    _new_filter_value = undefined;
    filter_modified = true;
    onModify(filter_conds)
  };
  let onRemove = (ftcond) => {
    const newf = filter_conds.filter(
      (c) =>
        c.column === ftcond.column &&
        c.cond === ftcond.cond &&
        c.value === ftcond.value
    );

    filter_conds = [...newf];
    filter_modified = true;
    onModify(filter_conds)
  };
</script>

<table class="w-full border">
  <thead>
    <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
      <th class="py-1 px-2 text-left">Column</th>
      <th class="py-1 px-2 text-left">Condition</th>
      <th class="py-1 px-2 text-center">Value</th>
      <th class="py-1 px-2 text-center">Action</th>
    </tr>
  </thead>
  <tbody class="text-gray-600 text-sm font-light">
    {#each filter_conds as ft}
      <tr class="border-b border-gray-200 hover:bg-gray-100">
        <td class="py-3 px-6 text-left whitespace-nowrap">
          <div class="flex items-center">
            <span class="font-medium">{ft.column}</span>
          </div>
        </td>
        <td class="py-3 px-6 text-left">
          <div class="flex items-center">
            <span
              class="bg-purple-200 text-purple-600 py-1 px-3 rounded-full text-xs"
              >{ft.cond}</span
            >
          </div>
        </td>

        <td class="py-3 px-6 text-center">
          {ft.value}
        </td>

        <td class="py-3 px-6 text-center">
          <button on:click={() => onRemove(ft)}>
            <Icon name="x-circle" class="h-5 w-5" />
          </button>
        </td>
      </tr>
    {/each}

    <tr class="border-b border-gray-200 hover:bg-gray-100">
      <td class="py-3 px-6 text-left whitespace-nowrap">
        <select class="p-1 rounded" bind:value={_new_column_slug}>
          {#each columns as col}
            <option value={col.slug}>{col.name}</option>
          {/each}
        </select>
      </td>
      <td class="py-3 px-6 text-left">
        <select class="p-1 rounded" bind:value={_new_column_cond}>
          <option />
          {#each Object.entries(_possible_cond) as [cond, condCompo]}
            <option value={cond}>{cond}</option>
          {/each}
        </select>
      </td>

      <td class="py-3 px-6 text-center">
        {#key _filer_component}
          {#if _filer_component}
            <svelte:component
              this={_filer_component[0]}
              bind:value={_new_filter_value}
              column={_current_column}
            />
          {/if}
        {/key}
      </td>

      <td class="py-3 px-6 text-center">
        <button
          class="bg-blue-300 rounded hover:bg-blue-400 p-1 text-white"
          on:click={onAdd}
        >
          Add
        </button>
      </td>
    </tr>
  </tbody>
</table>
