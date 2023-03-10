<script lang="ts">
  import type { PortalService } from "../../core";
  import { ResourceTypeSockRoom } from "../../../../../../lib/entities";
  import { generateId } from "../../../../../../lib/utils";

  export let service: PortalService;

  let slug = generateId();
  let name = "";

  const save = async () => {
    service.nav.admin_resource_new({
      type: ResourceTypeSockRoom,
      slug,
      name,
    });

    service.utils.small_modal_close();
  };
</script>

<div class="p-1 flex flex-col">
  <div class="text-2xl text-indigo-900">New Room</div>
  <p class="text-red-500" />
  <div class="flex-col flex py-3">
    <label for="field-0" class="pb-2 text-gray-700 font-semibold">Id</label>
    <input
      type="text"
      bind:value={slug}
      class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
    />
  </div>
  <div class="flex-col flex py-3">
    <label for="field-1" class="pb-2 text-gray-700 font-semibold">Name</label>
    <input
      bind:value={name}
      type="text"
      class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
    />
  </div>

  <button
    on:click={save}
    class="p-1 bg-blue-600 rounded hover:bg-blue-800 text-white">Next</button
  >
</div>
