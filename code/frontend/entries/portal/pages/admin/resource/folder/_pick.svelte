<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
</script>

<div class="p-2 flex flex-row justify-end">
  <select class="p-1 rounded border">
    <option>default</option>
    <option>default2</option>
  </select>
</div>

<table class="w-full">
  <thead
    ><tr
      class="text-md font-semibold tracking-wide text-left text-gray-900 bg-gray-100 uppercase border-b border-gray-600"
      ><th class="px-2 py-1" /> <th class="px-2 py-1">Name</th>
      <th class="px-4 py-1">File Size</th>
      <th class="px-4 py-1">Last Modified</th>
      <th class="px-4 py-1" />
    </tr></thead
  >
  <tbody class="bg-white"
    ><tr class="text-gray-700 hover:bg-gray-200"
      ><td class="px-2 py-1 border">
        <Icon name="folder" fill class="w-10" />
      </td>
      <td class="px-4 py-1 text-ms border">Public</td>
      <td class="px-4 py-1 text-xs border"
        ><span
          class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-sm"
          >--</span
        ></td
      > <td class="px-4 py-1 text-sm border">2023-01-05 18:01:00</td>

      <td class="px-4 py-1 text-sm border">
        <input type="checkbox" />
      </td>
    </tr>
  </tbody>
</table>
