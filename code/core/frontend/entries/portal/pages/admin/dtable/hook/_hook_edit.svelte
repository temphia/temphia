<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Kveditor from "../../../../../common/kveditor.svelte";
  import { DynAdminAPI } from "../dtable2";

  export let data = {};

  export let id = 0;
  export let table_id = "";
  export let group_id = "";
  export let source = "";

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let getData;
  let extra_meta_modified = false;

  $: _mod_data = {};
  $: _modified = false;

  const get = (name) => data[name] || "";
  const set = (name) => (ev) => {
    _mod_data[name] = ev.target.value;
    _modified = true;
  };

  const save = async () => {
    const reqdata = {
      ..._mod_data,
    };

    if (extra_meta_modified) {
      reqdata["extra_meta"] = getData();
    }

    const resp = await dynapi.modify_hook(
      source,
      group_id,
      table_id,
      id,
      reqdata
    );
    _mod_data = {};
    _modified = false;
  };
</script>

<div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
  <div class="p-5 bg-white w-full ">
    <div class="text-2xl text-indigo-900">Edit Hook</div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Id</label>
      <input
        type="text"
        value={get("id")}
        disabled
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Name</label>
      <input
        type="text"
        value={get("name")}
        on:change={set("name")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Type</label>
      <input
        type="text"
        value={get("type")}
        on:change={set("type")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Sub Type</label>
      <input
        type="text"
        value={get("sub_type")}
        on:change={set("sub_type")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Payload</label>

      <textarea
        value={get("payload")}
        on:change={set("payload")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Plug</label>
      <input
        type="text"
        value={get("plug_id")}
        on:change={set("plug_id")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="plugxyz"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Agent</label>
      <input
        type="text"
        value={get("agent_id")}
        on:change={set("agent_id")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="agentxyz"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
      <Kveditor
        data={data["extra_meta"] || {}}
        bind:getData
        bind:modified={extra_meta_modified}
      />
    </div>

    {#if _modified || extra_meta_modified}
      <div class="flex py-3">
        <button
          on:click={save}
          class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
        >
      </div>
    {/if}
  </div>
</div>
