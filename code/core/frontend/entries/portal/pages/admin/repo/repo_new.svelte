<script lang="ts">
  import { getContext } from "svelte";
  import { KvEditor } from "../../../../xcompo/common";
  import type { PortalService } from "../../../services";

  const app = getContext("__app__") as PortalService;
  const rapi = app.api_manager.get_admin_repo_api();

  let name = "";
  let url = "";
  let provider = "";
  let extra_meta = {};

  let message = "";

  const saveHandle = async () => {
    const resp = await rapi.new({
      name,
      url,
      provider,
      extra_meta,
    });
    if (resp.status !== 200) {
      message = resp.data;
      return;
    }

    // app.navigator.goto_admin_repo();
  };
</script>

<div class="h-full w-full overflow-auto">
  <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
    <div
      class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
    >
      <div class="text-2xl text-indigo-900">Add Repo</div>

      <p class="text-red-500">{message}</p>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          bind:value={name}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Provider</label>
        <input
          type="text"
          bind:value={provider}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">URL</label>
        <input
          type="text"
          bind:value={url}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
        <KvEditor bind:data={extra_meta} />
      </div>

      <div class="flex py-3">
        <button
          on:click={saveHandle}
          class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
        >
      </div>
    </div>
  </div>
</div>
