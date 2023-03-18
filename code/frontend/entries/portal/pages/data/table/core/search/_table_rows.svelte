<script lang="ts">
  import type { Column } from "../../../../../services/data";

  export let onRowSelect: (row: object) => void;
  export let data = {};
  export let columns: Column[] = Object.values(data["columns"] || {});
</script>

<div class="w-full overflow-auto">
  <table class="w-full border">
    <thead>
      <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
        {#each columns as col}
          <th class="px-1 py-3 text-xs text-left">{col["name"] || col.slug}</th>
        {/each}
        <th class="p-1 text-xs text-left"> # </th>
      </tr>
    </thead>

    <tbody class="text-gray-600 text-sm font-light">
      {#each data["rows"] as row}
        <tr class="border-b border-gray-200 hover:bg-gray-100">
          {#each columns as col}
            <td class="px-1 py-3 text-left">
              {#if col.slug in row && row[col.slug] !== null}
                {row[col.slug]}
              {/if}
            </td>
          {/each}
          <th class="p-1 text-xs text-left">
            <button
              on:click={() => onRowSelect(row)}
              class="text-blue-600 underline text-sm">goto</button
            >
          </th>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
