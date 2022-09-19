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

  export let id = "";
  export let name = "";
  export let type = "";
  export let sub_type = "";
  export let payload = "";
  export let policy = "";
  export let plug_id = "";
  export let target = "";

  export let saveFn;
  export let app: PortalApp;

  let getData;
</script>

<div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
  <div class="p-5 bg-white w-full ">
    <div class="text-2xl text-indigo-900">Resource</div>

    <div class="flex-col flex py-3">
      <label class="pb-2 text-gray-700 font-semibold">Id</label>
      <input
        type="text"
        bind:value={id}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        placeholder="res1"
      />
    </div>

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
      <select
        bind:value={type}
        class="form-select block w-full  p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      >
        {#each types as _type}
          <option value={_type.slug}>{_type.name}</option>
        {/each}
      </select>
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
      <label class="pb-2 text-gray-700 font-semibold">Target</label>
      <input
        type="text"
        bind:value={target}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Value</label>
      <textarea
        type="text"
        bind:value={payload}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Policy</label>
      <textarea
        bind:value={policy}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Plug</label>
      <input
        type="text"
        bind:value={plug_id}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
      <Kveditor data={{}} bind:getData />
    </div>

    <div class="flex py-3">
      <button
        on:click={() => {
          saveFn({
            id,
            name,
            type,
            sub_type,
            payload,
            target,
            policy,
            plug_id,
            extra_meta: getData(),
          });
        }}
        class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
      >
    </div>
  </div>
</div>
