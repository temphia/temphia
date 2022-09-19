<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";
  import Kveditor from "../../../../_shared/common/kveditor.svelte";
  import { DynAdminAPI } from "./dtable2";

  export let data = {};
  export let source;

  const app: PortalApp = getContext("__app__");
  let getExtraMeta;
  let extra_meta_modified;

  const dynapi = new DynAdminAPI(app);

  let mod_data = {};

  const set = (name: string) => (ev) => {
    mod_data[name] = ev.target.value;
  };
</script>

<div class="h-full w-full overflow-auto">
  <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
    >
      <div class="text-2xl text-indigo-900">Data Group</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          value={data["name"] || ""}
          on:change={set("name")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Description</label>
        <textarea
          type="text"
          value={data["description"] || ""}
          on:change={set("description")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          placeholder="About this data group..."
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Cabinet source</label>
        <input
          type="text"
          disabled
          value={data["cabinet_source"] || ""}
          on:change={set("cabinet_source")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Cabinet folder</label>
        <input
          type="text"
          disabled
          value={data["cabinet_folder"] || ""}
          on:change={set("cabinet_folder")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
        <Kveditor
          data={data["extra_meta"] || {}}
          bind:getData={getExtraMeta}
          bind:modified={extra_meta_modified}
        />
      </div>

      <div
        class="flex justify-between space-x-1"
        on:click={async () => {
          const _data = { ...mod_data };
          if (extra_meta_modified) {
            _data["extra_meta"] = getExtraMeta();
          }
          await dynapi.edit_group(source, data["slug"], _data);
          app.navigator.goto_admin_dgroup_page(source, data["slug"]);
        }}
      >
        <button class="p-2 bg-blue-400 text-white rounded">Save</button>
      </div>
    </div>
  </div>
</div>
