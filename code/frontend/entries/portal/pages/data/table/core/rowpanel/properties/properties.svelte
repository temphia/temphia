<script lang="ts">
  export let row;
</script>

<div class="w-full p-2">
  <table class="w-full border">
    <thead>
      <tr class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal">
        <th class="py-3 px-6 text-left">Properties</th>
        <th class="py-3 px-6 text-left">Value</th>
      </tr>
    </thead>
    <tbody class="text-gray-600 text-sm font-light">
      <tr class="border-b border-gray-200 hover:bg-gray-100">
        <td class="py-3 px-6 text-left">Id</td>
        <td class="py-3 px-6 text-left">{row["__id"] || 0} </td>
      </tr>
      <tr class="border-b border-gray-200 hover:bg-gray-100">
        <td class="py-3 px-6 text-left">Version</td>
        <td class="py-3 px-6 text-left">{row["__version"] || 0}</td>
      </tr>

      {#if "deleted_at" in row}
        <td class="py-3 px-6 text-left">>Deleted At</td>
        <td class="py-3 px-6 text-left">>{row["deleted_at"] || ""}</td>
      {/if}
    </tbody>
  </table>
</div>
