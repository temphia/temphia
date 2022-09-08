<script lang="ts">
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import { ActionAddButton } from "../../../../../common";

  export let onSelect;
  export let app: PortalApp;

  let resources = [];
  let selected = "";

  (async () => {
    const apm = app.get_apm();
    const rapi = await apm.get_resource_api();
    const resp = await rapi.resource_list();
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }

    resources = resp.data;
  })();
</script>

<table class="min-w-full divide-y divide-gray-200 table-fixed">
  <thead class="bg-gray-100">
    <tr>
      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase"
      >
        Id
      </th>
      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase"
      >
        Name
      </th>
      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase"
      >
        Type
      </th>

      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase"
      >
        Sub Type
      </th>

      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase"
      >
        Target
      </th>

      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase"
      >
        Plug Id
      </th>

      <th scope="col" class="p-4" />
    </tr>
  </thead>
  <tbody class="bg-white divide-y divide-gray-200">
    {#each resources as resource}
      <tr class="hover:bg-gray-100">
        <td
          class="py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap"
          >{resource["id"] || ""}</td
        >

        <td
          class="py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap"
          >{resource["name"] || ""}</td
        >

        <td
          class="py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap"
          >{resource["type"] || ""}</td
        >

        <td
          class="py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap"
          >{resource["sub_type"] || ""}</td
        >

        <td
          class="py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap"
          >{resource["target"] || ""}</td
        >

        <td
          class="py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap"
          >{resource["plug_id"] || ""}</td
        >

        <td class="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
          <a href="#" class="text-blue-600 hover:underline">Remove</a>
        </td>
      </tr>
    {/each}
  </tbody>
</table>

<div class="flex justify-between">
  <div />
  <div>
    {#if selected}
      <ActionAddButton
        onClick={() => {
          if (!onSelect) {
            return;
          }
          onSelect(selected);
        }}
      />
    {/if}
  </div>
</div>
