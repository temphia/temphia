<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { CEditor, LoadingSpinner, PortalService } from "../../core";

  export let bid = $params.bid;
  export let file = $params._;

  let loading = true;
  let text = "";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_bprint_api();

  const load = async () => {

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

  const save = async () => {
    // api.update_file(bid, file, "")
  }



</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="w-full h-full p-2">
    <div class="flex flex-col bg-white rounded-xl">
      <div class="flex justify-between border border-gray-100">
        <div
          class="grow h-10 flex flex-row flex-nowrap overflow-hidden justify-between"
        >
          <div
            class=" py-2 px-2 flex hover:text-blue-500 cursor-pointer focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500"
          >
            {file}
          </div>

          <div class="flex gap-1 p-1">
            <button class="hover:bg-gray-300 rounded inline-flex border p-1">
              <Icon name="beaker" class="h-5 w-5" />
              Editor
            </button>

            <button class="hover:bg-gray-300 rounded inline-flex border p-1">
              <Icon name="save" class="h-5 w-5" />
              Save
            </button>

            <button
              on:click={() => app.nav.admin_bprint_files(bid)}
              class="hover:bg-gray-300 rounded inline-flex border p-1"
            >
              <Icon name="arrow-up" class="h-5 w-5" />
              Back
            </button>
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
