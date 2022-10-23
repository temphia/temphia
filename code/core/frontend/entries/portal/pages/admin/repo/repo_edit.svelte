<script lang="ts">
  import { getContext } from "svelte";
  import { KvEditor } from "../../../../xcompo/common";
  import type { PortalService } from "../../../services";

  const app = getContext("__app__") as PortalService;
  const rapi = app.api_manager.get_admin_repo_api();

  export let rid;

  let data = {};
  let message = "";

  let get_meta_opts;
  let meta_opts_modified = false;

  let modified = false;
  let loading = true;

  const set = (name) => (ev) => {
    data = { ...data, [name]: ev.target.value };
    modified = true;
  };

  const load = async () => {
    const resp = await rapi.get(rid);
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();

  const saveHandle = async () => {
    const fdata = { ...data };

    if (meta_opts_modified) {
      fdata["extra_meta"] = get_meta_opts();
    }

    rapi.update(rid, fdata);
  };
</script>

<div class="h-full w-full overflow-auto">
  <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
    >
      <div class="text-2xl text-indigo-900">Edit Repo</div>
      <p class="text-red-500">{message}</p>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Id</label>
        <input
          type="text"
          disabled
          value={data["id"]}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

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
        <label class="pb-2 text-gray-700 font-semibold">Provider</label>
        <input
          type="text"
          value={data["provider"] || ""}
          on:change={set("provider")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">URL</label>
        <input
          type="text"
          value={data["url"] || ""}
          on:change={set("url")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
        <KvEditor
          data={data["extra_meta"] || {}}
          bind:getData={get_meta_opts}
          bind:modified={meta_opts_modified}
        />
      </div>

      <div class="flex py-3">
        {#if modified || meta_opts_modified}
          <button
            class="p-2 bg-blue-400 m-1 w-20 text-white rounded"
            on:click={saveHandle}>Save</button
          >
        {/if}
      </div>
    </div>
  </div>
</div>
