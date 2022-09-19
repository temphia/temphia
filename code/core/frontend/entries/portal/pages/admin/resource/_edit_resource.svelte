<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";
  import Kveditor from "../../../../_shared/common/kveditor.svelte";

  const types = [
    { name: "⚡ Socket Room", slug: "socket_room" },
    { name: "💾 Data Table", slug: "dtable" },
    { name: "💾 Data Group", slug: "dgroup" },
    { name: "🗄️ Folder", slug: "cabinet_folder" },
  ];

  export let data = {};
  $: _mod_data = {};
  $: _modified = false;

  const app: PortalApp = getContext("__app__");
  let getData;
  let extra_meta_modified = false;

  const get = (name) => data[name] || "";
  const set = (name) => (ev) => {
    _mod_data[name] = ev.target.value;
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
    <div class="text-2xl text-indigo-900">Resource</div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Id</label>
      <input
        type="text"
        value={get("id")}
        disabled
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="res1"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Name</label>
      <input
        type="text"
        value={get("name")}
        on:change={set("name")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Type</label>
      <input
        type="text"
        value={get("type")}
        on:change={set("type")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Sub Type</label>
      <input
        type="text"
        value={get("sub_type")}
        on:change={set("sub_type")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="signal"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Target</label>
      <input
        type="text"
        value={get("target")}
        on:change={set("target")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Value</label>
      <textarea
        value={get("payload")}
        on:change={set("payload")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Policy</label>
      <textarea
        value={get("policy")}
        on:change={set("policy")}
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
