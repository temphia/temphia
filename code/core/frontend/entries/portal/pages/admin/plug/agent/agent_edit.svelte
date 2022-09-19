<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import {
    KvEditor,
  } from "../../../../../_shared/common";
  import ResourcesPicker from "./_resources_picker.svelte";

  export let data = {};
  $: _mod_data = {};
  $: _modified = false;

  const app: PortalApp = getContext("__app__");

  let getMetaData;
  let meta_modified = false;

  let getWebFiles;
  let web_files_modified = false;

  let getEnvVars;
  let env_vars_modified = false;

  const get = (name) => data[name] || "";
  const set = (name) => (ev) => {
    _mod_data[name] = ev.target.value;
    _modified = true;
  };

  const save = async () => {
    const rapi = await app.get_apm().get_resource_api();
    if (meta_modified) {
      _mod_data["extra_meta"] = getMetaData();
    }

    // fixme other get like this way ^

    await rapi.resource_update(data["id"], _mod_data);
    _mod_data = {};
    _modified = false;
  };

  const addResource = () =>
    app.simple_modal_open(ResourcesPicker, {
      onSelect: (res) => {
        console.log(res);
      },
    });

  const previewResources = () => {};
</script>

<div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
  <div class="p-5 bg-white w-full ">
    <div class="text-2xl text-indigo-900">Agent</div>

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
      <label class="pb-2 text-gray-700 font-semibold">Type</label>
      <input
        type="text"
        value={get("type")}
        on:change={set("type")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Executor</label>
      <input
        type="text"
        value={get("executor")}
        on:change={set("executor")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Iface File</label>
      <input
        type="text"
        value={get("iface_file")}
        on:change={set("iface_file")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Entry File</label>
      <input
        type="text"
        value={get("entry_file")}
        on:change={set("entry_file")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Web Entry</label>
      <input
        type="text"
        value={get("web_entry")}
        on:change={set("web_entry")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Web Script</label>
      <input
        type="text"
        value={get("web_script")}
        on:change={set("web_script")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Web Style</label>
      <input
        type="text"
        value={get("web_script")}
        on:change={set("web_script")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Web Loader</label>
      <input
        type="text"
        value={get("web_script")}
        on:change={set("web_script")}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Web Files</label>
      <KvEditor
        data={data["web_files"] || {}}
        bind:getData={getWebFiles}
        bind:modified={web_files_modified}
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Env Vars</label>
      <KvEditor
        data={data["env_vars"] || {}}
        bind:getData={getEnvVars}
        bind:modified={env_vars_modified}
      />
    </div>

    <div class="flex-col flex py-3 relative">
      <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
      <KvEditor
        data={data["extra_meta"] || {}}
        bind:getData={getMetaData}
        bind:modified={meta_modified}
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

    {#if _modified || meta_modified}
      <div class="flex py-3">
        <button
          on:click={save}
          class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
        >
      </div>
    {/if}
  </div>
</div>
