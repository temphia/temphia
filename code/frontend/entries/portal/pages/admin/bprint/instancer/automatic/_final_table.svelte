<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let installed_items = {};

  const iconTypes = {
    data_group: "collection",
    plug: "view-grid-add",
    resource: "cube",
    data_sheet: "table",
  };
</script>

<table class="min-w-max w-full table-auto border">
  <thead>
    <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
      <th class="py-3 px-6 text-left">#</th>
      <th class="py-3 px-6 text-left">Name</th>
      <th class="py-3 px-6 text-left">Item</th>
      <th class="py-3 px-6 text-left">Type</th>
      <th class="py-3 px-6 text-center">Actions</th>
    </tr>
  </thead>
  <tbody class="text-gray-600 text-sm font-light">
    {#each Object.entries(installed_items || {}) as [itemkey, item]}
      <tr class="border-b border-gray-200 hover:bg-gray-100">
        <td class="py-3 px-6 text-left whitespace-nowrap">
          <Icon
            name={iconTypes[item["type"]] || "hashtag"}
            class="w-6 h-6 text-blue-600"
          />
        </td>
        <td class="py-3 px-6 text-left">
          <div class="flex items-center">
            <span>{itemkey}</span>
          </div>
        </td>

        <td class="py-3 px-6 text-left">
          <span class="flex items-center bg-gray-100 rounded"
            >{item["slug"]}</span
          >
        </td>

        <td class="py-3 px-6 text-center">
          <span class="text-sm bg-gray-200 p-1 rounded">
            {item["type"]}
          </span>
        </td>

        <td class="py-3 px-6 text-center">
          {#if item["message"]}
            <p>
              {item["message"]}
            </p>
          {:else}
            <button
              class="bg-blue-200 text-blue-600 py-1 px-3 rounded-full text-xs"
              >explore</button
            >
          {/if}
        </td>
      </tr>
    {/each}
  </tbody>
</table>
