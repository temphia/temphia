<script lang="ts">
  import type { AxiosResponse } from "axios";
  import PanelLayout from "../_panel_layout.svelte";
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

<PanelLayout
  loading={!loaded}
  selected={!!selected}
  onSelect={() => onRowSelect(selected)}
>
  <table class="w-full border overflow-scroll">
    <thead>
      <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
        <th class="text-left" />
        {#each Object.entries(data["columns"]) as [colslug, col]}
          <th class="text-left">{col["name"] || colslug}</th>
        {/each}
      </tr>
    </thead>

    <tbody class="text-gray-800 text-sm font-light">
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
</PanelLayout>
