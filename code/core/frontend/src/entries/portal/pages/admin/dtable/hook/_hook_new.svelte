<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Kveditor from "../../../../../common/kveditor.svelte";
  import { DynAdminAPI } from "../dtable2";

  export let name = "";
  export let type = "";
  export let sub_type = "";
  export let payload = "";
  export let plug_id = "";
  export let agent_id = "";
  export let table_id = "";
  export let group_id = "";
  export let source = "";

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let getData;

  const save = async () => {
    const resp = await dynapi.new_hook(source, group_id, table_id, {
      name,
      type,
      sub_type,
      payload,
      plug_id,
      agent_id,
      table_id,
      group_id,
      extra_meta: getData(),
    });

    app.navigator.goto_admin_dtable_page(source, group_id, table_id);
  };
</script>

<div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
  <div class="p-5 bg-white w-full ">
    <div class="text-2xl text-indigo-900">New Hook</div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Name</label>
      <input
        type="text"
        bind:value={name}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Type</label>
      <input
        type="text"
        bind:value={type}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Sub Type</label>
      <input
        type="text"
        bind:value={sub_type}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Payload</label>

      <textarea
        bind:value={payload}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Plug</label>
      <input
        type="text"
        bind:value={plug_id}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="plugxyz"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Agent</label>
      <input
        type="text"
        bind:value={agent_id}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="agentxyz"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
      <Kveditor data={{}} bind:getData />
    </div>

    <div class="flex py-3">
      <button
        on:click={save}
        class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
      >
    </div>
  </div>
</div>
