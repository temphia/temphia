<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../app";

  const app: PortalApp = getContext("__app__");

  app.get_cabinet_sources().then((resp) => {
    if (!resp) {
      return;
    }

    app.navigator.goto_cabinet_source(resp[0]);
  });
</script>

<div class="w-full h-full bg-gray-100 p-10">
  <div class="bg-white rounded">
    <table class="text-left w-full border-collapse">
      <thead>
        <tr>
          <th
            class="py-3 px-6 font-semibold text-sm text-gray-700 border-b border-grey-light"
          >
            <div class="w-32 h-10 m-1 bg-gray-200 animate-pulse" />
          </th>
          <th
            class="py-3 px-6 font-semibold text-sm text-gray-700 border-b border-grey-light"
          >
            <div class="w-32 h-10 m-1 bg-gray-200 animate-pulse" />
          </th>
          <th
            class="py-3 px-6 font-semibold text-sm text-gray-700 border-b border-grey-light"
          >
            <div class="w-32 h-10 m-1 bg-gray-200 animate-pulse" />
          </th>
        </tr>
      </thead>
      <tbody class="overflow-auto w-full">
        {#each Array(10) as _}
          <tr class="py-2">
            <td class="px-6">
              <div class="w-32 h-10 m-1 bg-gray-200 animate-pulse" />
            </td>
            <td class="py-4 px-6">
              <div class="w-32 h-10 m-1 bg-gray-200 animate-pulse" />
            </td>
            <td class="py-4 px-6">
              <div class="w-32 h-10 m-1 bg-gray-200 animate-pulse" />
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
  </div>
</div>
