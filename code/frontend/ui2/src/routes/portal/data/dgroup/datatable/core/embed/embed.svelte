<script lang="ts">
  export let onRowSelect: (row: object) => void;
  export let data = {};
  export let selectable = false;

  let selected = null;
</script>

<div class="w-full overflow-auto">
  <table class="w-full border">
    <thead>
      <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
        <th class="p-1 text-xs text-left" />
        {#each Object.entries(data["columns"]) as [colslug, col]}
          <th class="p-1 text-xs text-left">{col["name"] || colslug}</th>
        {/each}
      </tr>
    </thead>

    <tbody class="text-gray-600 text-sm font-light">
      {#each data["rows"] as row}
        <tr class="border-b border-gray-200 hover:bg-gray-100">
          <td class="p-1">
            {#if selectable}
              <input
                type="checkbox"
                checked={row === selected}
                class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                on:click={() => {
                  if (selected === row) {
                    selected = null;
                  } else {
                    selected = row;
                    onRowSelect(selected);
                  }
                }}
              />
            {/if}
          </td>

          {#each Object.entries(data["columns"]) as [colslug, col]}
            <td class="p-1 text-left">
              {#if colslug in row && row[colslug] !== null}
                {row[colslug]}
              {/if}
            </td>
          {/each}
        </tr>
      {/each}
    </tbody>
  </table>
</div>
