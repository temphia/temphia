<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";
  import Kveditor from "../../../../_shared/common/kveditor.svelte";

  export let data = {};
  $: _mod_data = {};
  $: _modified = false;

  const app: PortalApp = getContext("__app__");

  let getMetaData;
  let meta_modified = false;

  const get = (name) => data[name] || "";
  const set = (name) => (ev) => {
    _mod_data[name] = ev.target.value;
    _modified = true;
  };

  const setBool = (name) => (ev) => {
    _mod_data = { ..._mod_data, [name]: ev.target.checked };
    _modified = true;
  };

  const save = async () => {
    const rapi = await app.get_apm().get_resource_api();
    await rapi.resource_update(data["id"], _mod_data);
    _mod_data = {};
    _modified = false;
  };
</script>

<div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
  <div class="p-5 bg-white w-full ">
    <div class="text-2xl text-indigo-900">Plug</div>

    <div class="flex-col flex py-3">
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
      <label class="pb-2 text-gray-700 font-semibold">Live</label>
      <input
        type="checkbox"
        value={get("live") || false}
        on:change={setBool("live")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder=""
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Dev</label>
      <input
        type="checkbox"
        value={get("dev") || false}
        on:change={setBool("dev")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder=""
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Owner</label>
      <input
        type="text"
        value={get("owner")}
        on:change={set("owner")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Policy</label>
      <textarea
        value={get("invoke_policy")}
        on:change={set("invoke_policy")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Bprint</label>
      <input
        type="text"
        value={get("bprint_id")}
        on:change={set("bprint_id")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>


    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
      <Kveditor
        data={data["extra_meta"] || {}}
        bind:getData={getMetaData}
        bind:modified={meta_modified}
      />
    </div>

    {#if _modified  || meta_modified}
      <div class="flex py-3">
        <button
          on:click={save}
          class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
        >
      </div>
    {/if}
  </div>
</div>
