<script lang="ts">
  import type { Panel } from "../service";
  export let panel: Panel;
  export let data;
</script>

<h4 class="text-xl leading-none font-bold text-gray-600 mb-5">
  {panel.name || ""}
</h4>

<div class="block w-full overflow-x-auto">
  <table class="items-center w-full bg-transparent border-collapse">
    <thead>
      <tr>
        <th
          class="px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap"
          >{data["key_name"]}</th
        >
        <th
          class="px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap"
          >{data["value_name"]}</th
        >
        <th
          class="px-4 bg-gray-50 text-gray-700 align-middle py-3 text-xs font-semibold text-left uppercase border-l-0 border-r-0 whitespace-nowrap min-w-140-px"
        />
      </tr>
    </thead>
    <tbody class="divide-y divide-gray-100">
      {#each data["items"] || [] as item}
        <tr class="text-gray-500">
          <th
            class="border-t-0 px-4 align-middle text-sm font-normal whitespace-nowrap p-4 text-left"
            >{item[0]}</th
          >
          <td
            class="border-t-0 px-4 align-middle text-xs font-medium text-gray-900 whitespace-nowrap p-4"
            >{item[1]}</td
          >
          <td
            class="border-t-0 px-4 align-middle text-xs whitespace-nowrap p-4"
          >
            <div class="flex items-center">
              <span class="mr-2 text-xs font-medium">30%</span>
              <div class="relative w-full">
                <div class="w-full bg-gray-200 rounded-sm h-2">
                  <div class="bg-cyan-600 h-2 rounded-sm" style="width: 30%" />
                </div>
              </div>
            </div>
          </td>
        </tr>
      {/each}
    </tbody>
  </table>
</div>
