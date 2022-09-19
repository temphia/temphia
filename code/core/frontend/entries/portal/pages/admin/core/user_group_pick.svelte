<script lang="ts">
  import type { PortalApp } from "../../../app";

  export let app: PortalApp;
  export let onSelected;
  export let selected_ugroup = "";

  let user_groups = [];

  (async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.list_user_group();
    user_groups = resp.data;
  })();
</script>

<h3 class="mb-3 text-xl font-medium text-gray-900 dark:text-white">
  User Groups
</h3>
<table
  class="w-full border text-sm text-left text-gray-500 dark:text-gray-400 overflow-auto"
  style="max-width: 40rem;"
>
  <thead
    class="text-xs text-gray-700 uppercase bg-gray-50 dark:bg-gray-700 dark:text-gray-400"
  >
    <tr>
      <th scope="col" class="p-1" />
      <th scope="col" class="px-2 py-1"> Name </th>
      <th scope="col" class="px-2 py-1"> Slug </th>
      <th scope="col" class="px-2 py-1" />
    </tr>
  </thead>
  <tbody>
    {#each user_groups as ugroup}
      <tr
        class="bg-white border-b dark:bg-gray-800 dark:border-gray-700 hover:bg-gray-50 dark:hover:bg-gray-600"
      >
        <td class="w-4 p-4">
          <input
            type="checkbox"
            class="w-4 h-4 text-blue-600 bg-gray-100 border-gray-300 rounded focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
            on:click={() => {
              if (selected_ugroup === ugroup.slug) {
                selected_ugroup = "";
              } else {
                selected_ugroup = ugroup.slug;
              }
            }}
            checked={selected_ugroup === ugroup.slug}
          />
        </td>
        <th
          scope="row"
          class="px-6 py-4 font-medium text-gray-900 dark:text-white whitespace-nowrap"
        >
          {ugroup.name}
        </th>
        <td class="px-6 py-4"> {ugroup.slug} </td>
        <td class="px-6 py-4" />
      </tr>
    {/each}
  </tbody>
</table>

<div class="flex justify-between mt-4">
  <div />
  <div>
    {#if selected_ugroup}
      <button
        class="p-1 bg-green-500 hover:bg-green-700 text-white rounded"
        on:click={() => {
          if (!onSelected) {
            return;
          }

          onSelected(selected_ugroup);
        }}>Select</button
      >
    {/if}
  </div>
</div>
