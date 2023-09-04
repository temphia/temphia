<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { PortalService, Uploader } from "$lib/core";

  export let app: PortalService;
  export let close_modal;
  export let bid: string;

  export let afterFn: any;

  const api = app.api_manager.get_admin_bprint_api();

  let upload_mode = true;
  let filename = "";
</script>

<div class="inline-flex items-center rounded-md shadow-sm">
  <button
    on:click={() => {
      upload_mode = false;
    }}
    class="text-slate-800 hover:text-blue-600 text-sm bg-white hover:bg-slate-100 border border-slate-200 rounded-l-lg font-medium px-4 py-2 inline-flex space-x-1 items-center {upload_mode
      ? ''
      : 'text-blue-600 bg-slate-100'}"
  >
    <span>
      <Icon name="plus-circle" class="w-4 h-4" />
    </span>
    <span>Create</span>
  </button>

  <button
    on:click={() => {
      upload_mode = true;
    }}
    class="text-slate-800 hover:text-blue-600 text-sm bg-white hover:bg-slate-100 border border-slate-200 rounded-r-lg font-medium px-4 py-2 inline-flex space-x-1 items-center {upload_mode
      ? 'text-blue-600 bg-slate-100'
      : ''}"
  >
    <span>
      <Icon name="upload" class="w-4 h-4" />
    </span>
    <span>Upload</span>
  </button>
</div>

{#key upload_mode}
  {#if upload_mode}
    <Uploader
      uploadFile={async (file, data) => {
        await api.add_file(bid, file, data);
        close_modal();
        if (afterFn) {
          afterFn();
        }
      }}
    />
  {:else}
    <div class="grid grid-cols-1 space-y-2">
      <div class="text-center">
        <h2 class="text-3xl font-bold text-gray-900">Create empty File</h2>
      </div>

      <label
        for="new_file_name"
        class="text-sm font-bold text-gray-500 tracking-wide">File Name</label
      >
      <input
        id="new_file_name"
        class="text-base p-2 border border-gray-300 rounded-lg focus:outline-none focus:border-indigo-500"
        type="text"
        bind:value={filename}
        placeholder="file.txt"
      />

      <button
        on:click={async () => {
          const formdata = new FormData();
          formdata.append("file", "");

          await api.add_file(bid, filename, formdata);
          close_modal();
          if (afterFn) {
            afterFn();
          }
        }}
        class="my-5 w-full flex justify-center bg-blue-500 text-gray-100 p-4  rounded-full tracking-wide font-semibold  focus:outline-none focus:shadow-outline hover:bg-blue-800 shadow-lg cursor-pointer transition ease-in duration-300"
      >
        Create
      </button>
    </div>
  {/if}
{/key}
