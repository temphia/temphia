<script lang="ts">
  import { CtypeFilterConds } from "../../../../../../lib/core/dyntypes";

  import {
    CtypeShortText,
    CtypePhone,
    CtypeSelect,
    CtypeRFormula,
    CtypeFile,
    CtypeMultiFile,
    CtypeCheckBox,
    CtypeCurrency,
    CtypeNumber,
    CtypeLocation,
    CtypeDateTime,
    CtypeMultSelect,
    CtypeLongText,
    CtypeSingleUser,
    CtypeMultiUser,
    CtypeEmail,
    CtypeJSON,
    CtypeRangeNumber,
    CtypeColor,
  } from "../../../dtable/renderer/fields/field";

  export let columns = [];
  export let filter_conds: {
    column: string;
    cond: string;
    value: string;
  }[] = [];

  export let filter_modified = false;

  const colindexed = columns.reduce((acc, curr) => {
    acc[curr.slug] = curr;
    return acc;
  }, {});

  $: _new_column_slug = "";
  $: _new_column_cond = "";
  $: _new_filter_type =
    (colindexed[_new_column_slug] || {}).ctype || CtypeShortText;
  $: _possible_cond = CtypeFilterConds[_new_filter_type] || [];
  $: _new_filter_value = "";

  export let onAdd = () => {
    filter_conds = [
      ...filter_conds,
      {
        column: _new_column_slug,
        cond: _new_column_cond,
        value: _new_filter_value,
      },
    ];
    filter_modified = true;
  };
  export let onRemove = (ftcond) => {
    const newf = filter_conds.filter(
      (c) =>
        c.column !== ftcond.column &&
        c.cond !== ftcond.cond &&
        c.value !== ftcond.value
    );

    filter_conds = [...newf];
    filter_modified = true;
  };
</script>

<table class="w-full border">
  <thead>
    <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
      <th class="py-3 px-6 text-left">Column</th>
      <th class="py-3 px-6 text-left">Condition</th>
      <th class="py-3 px-6 text-center">Value</th>
      <th class="py-3 px-6 text-center">Action</th>
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
            <svg
              xmlns="http://www.w3.org/2000/svg"
              class="h-5 w-5"
              viewBox="0 0 20 20"
              fill="currentColor"
            >
              <path
                fill-rule="evenodd"
                d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                clip-rule="evenodd"
              />
            </svg>
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
          {#each _possible_cond as cond}
            <option value={cond}>{cond}</option>
          {/each}
        </select>
      </td>

      <td class="py-3 px-6 text-center">
        <!-- FIXME => implement all types and their cond -->
        {#if _new_filter_type === CtypeCurrency || _new_filter_type === CtypeNumber}
          <input
            type="number"
            class="w-full border"
            bind:value={_new_filter_value}
          />
        {:else if _new_filter_type === CtypeCheckBox}
          <select class="rounded p-1" bind:value={_new_filter_value}>
            <option>yes</option>
            <option>no</option>
          </select>
        {:else if _new_filter_type === CtypeDateTime}
          <input
            type="datetime-local"
            class="w-full border"
            bind:value={_new_filter_value}
          />
        {:else if _new_filter_type === CtypeLocation}
          <div>not implemented</div>
        {:else}
          <input
            type="text"
            class="w-full border"
            bind:value={_new_filter_value}
          />
        {/if}
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
