<script lang="ts">
  import type { AxiosResponse } from "axios";
  export let loader: (cursor: number) => Promise<AxiosResponse<any>>;
  export let onRowSelect: (row: object) => void;

  let loaded = false;
  let data = {};
  let selected = null;

  loader(0).then((resp) => {
    data = resp.data;
    loaded = true;
  });
</script>

<div class="w-full flex flex-col" style="height: 80vh;">
  <div class="flex-grow flex flex-col h-32 p-2 space-y-1 overflow-y-auto">
    {#if loaded}
      <table class="w-full border overflow-scroll">
        <thead>
          <tr
            class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal"
          >
            <th class="text-left" />
            {#each Object.entries(data["columns"]) as [colslug, col]}
              <th class="text-left">{col["name"] || colslug}</th>
            {/each}
          </tr>
        </thead>

        <tbody class="text-gray-600 text-sm font-light">
          {#each data["rows"] as row}
            <tr class="border-b border-gray-200 hover:bg-gray-100">
              <td class="p-1">
                <input
                  type="checkbox"
                  checked={row === selected}
                  class="focus:ring-blue-500 h-4 w-4 text-blue-600 border-gray-300 rounded"
                  on:click={() => {
                    if (selected === row) {
                      selected = null;
                    } else {
                      selected = row;
                    }
                  }}
                />
              </td>

              {#each Object.entries(data["columns"]) as [colslug, col]}
                <td class="text-left">
                  {#if colslug in row && row[colslug] !== null}
                    {row[colslug]}
                  {/if}
                </td>
              {/each}
            </tr>
          {/each}
        </tbody>
      </table>
    {:else}
      <div>Loading...</div>
    {/if}
  </div>

  <div class="flex-shrink h-10 w-full pb-1 border-b flex justify-between">
    {#if selected}
      <div class="flex mt-2 justify-end">
        <button
          class="px-2 py-1 shadow bg-green-500 rounded text-white font-semibold"
          on:click={() => {
            onRowSelect(selected);
          }}
        >
          Select
        </button>
      </div>
    {/if}
  </div>
</div>
