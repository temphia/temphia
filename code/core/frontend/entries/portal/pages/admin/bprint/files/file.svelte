<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { CEditor, LoadingSpinner, PortalService } from "../../core";
  import PickEditor from "./_pick_editor.svelte";

  export let bid = $params.bid;
  export let file = $params._;

  let loading = true;
  let code = "";
  let editor;
  let modified = false;
  let ext = (file || "a.js").split(".").pop();

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_bprint_api();

  const load = async () => {
    const resp = await api.get_file(bid, file);
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    if (typeof resp.data === "object") {
      code = JSON.stringify(resp.data, undefined, 2);
    } else {
      code = resp.data;
    }

    loading = false;
  };

  load();

  let message = "";
  let save_loading = false;
  const save = async () => {
    if (loading || save_loading) {
      return;
    }

    save_loading = true;

    const formdata = new FormData();
    formdata.append("file", editor.getValue());
    const resp = await api.update_file(bid, file, formdata);
    save_loading = false;
    if (resp.ok) {
      modified = false;
      message = "";
      return;
    }
    message = await resp.text();
  };
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
            <button
              class="hover:bg-gray-300 rounded inline-flex border p-1"
              on:click={() => app.utils.small_modal_open(PickEditor, { app })}
            >
              <Icon name="beaker" class="h-5 w-5" />
              Editor
            </button>

            <button
              class="hover:bg-gray-300 rounded inline-flex border p-1"
              on:click={save}
            >
              <Icon
                name="save"
                class="h-5 w-5 {save_loading ? 'animate-spin' : ''}"
              />
              {modified ? "*" : ""} Save
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

      <div class="toolbar">
        <p class="text-red-600">
          {message}
        </p>
      </div>

      <div class="h-full bg-white rounded p-2 ">
        <CEditor
          {code}
          mode={ext}
          bind:editor
          container_style="height:85vh;"
          on:change={(ev) => {
            modified = true;
          }}
        />
      </div>
    </div>
  </div>
{/if}
