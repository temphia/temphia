<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { CEditor, LoadingSpinner, PortalService } from "../../core";

  export let bid = $params.bid;
  export let file = $params._;

  let loading = true;
  let text = "";

  const app = getContext("__app__") as PortalService;

  const load = async () => {
    const api = app.api_manager.get_admin_bprint_api();
    const resp = await api.get_file(bid, file);
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    if (typeof resp.data === "object") {
      text = JSON.stringify(resp.data);
    } else {
      text = resp.data;
    }

    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="w-full h-full p-2">
    <div class="flex flex-col bg-white rounded-xl">
      <div class="flex justify-between border border-gray-100">
        <div class="grow h-10 flex flex-row flex-nowrap overflow-hidden">
          <div
            class=" py-2 px-2 flex hover:text-blue-500 cursor-pointer focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500"
          >
            {file}
          </div>
        </div>
      </div>

      <div class="toolbar" />

      <div class="h-full bg-white rounded p-2 ">
        <CEditor code={text} container_style="height:85vh;" />
      </div>
    </div>
  </div>
{/if}
