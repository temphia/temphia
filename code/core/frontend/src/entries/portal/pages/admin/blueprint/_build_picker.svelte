<script lang="ts">
  import type { PortalApp } from "../../../../../lib/app/portal";

  export let bid;
  export let app: PortalApp;

  let files = [];
  let selected = "";

  (async () => {
    const bapi = await app.get_apm().get_bprint_api();
    const resp = await bapi.bprint_list_files(bid);
    if (resp.status !== 200) {
      return;
    }
    files = resp.data;
  })();
</script>

<table class="w-full text-sm text-left text-gray-500 ">
  <thead class="text-xs text-gray-700 uppercase bg-gray-50">
    <tr>
      <th scope="col" class="p-4"> <div class="flex items-center" /></th>
      <th scope="col" class="px-6 py-3"> Name </th>
      <th scope="col" class="px-6 py-3">
        <span class="sr-only">Open</span>
      </th>
    </tr>
  </thead>
  <tbody>
    {#each files as file}
      <tr class="bg-white border-b hover:bg-gray-50 ">
        <td class="w-4 p-4">
          <div class="flex items-center">
            <input
              type="checkbox"
              on:click={() => {
                if (selected === file) {
                  selected = "";
                } else {
                  selected = file;
                }
              }}
              checked={selected === file}
              class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500"
            />
          </div>
        </td>
        <th
          scope="row"
          class="px-6 py-4 font-medium text-gray-900 whitespace-nowrap"
        >
          {file}
        </th>
        <td class="px-6 py-4 text-right">
          <a href="#" class="font-medium text-blue-600 hover:underline">Edit</a>
        </td>
      </tr>
    {/each}
  </tbody>
</table>

<div class="flex justify-between mt-4">
  <div />
  <div>
    {#if selected}
      <button
        class="p-1 bg-green-500 hover:bg-green-700 text-white rounded"
        on:click={() => {
          app.navigator.goto_admin_dtable_builder(bid, selected);
        }}>DTABLE BUILDER</button
      >
    {/if}
  </div>
</div>
