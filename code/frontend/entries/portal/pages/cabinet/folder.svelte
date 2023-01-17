<script lang="ts">
  import Layout from "./_layout.svelte";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import { params } from "svelte-hash-router";

  import PanelUploadFile from "./panels/upload_file.svelte";
  import { LoadingSpinner } from "../admin/core";
  import { humanizeBytes, isImage } from "../../../../lib/utils";
  import type { FolderTktAPI } from "../../../../lib/apiv2";

  const app: PortalService = getContext("__app__");

  const source = $params.source;
  const folder = $params.folder;

  let files = [];
  let files_loaded = false;
  let ticket_loaded = false;
  let preview = false;
  let fapi: FolderTktAPI;

  const cabservice = app.get_cabinet_service();
  const capi = cabservice.get_source_api(source);

  let sources = [];
  cabservice.get_cab_sources().then((sresp) => {
    sources = sresp;
  });

  const load = async () => {
    const resp = await capi.listFolder(folder);
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    fapi = await cabservice.get_folder_api(source, folder);

    files = resp.data;
    files_loaded = true;
    ticket_loaded = true;
  };

  load();

  const finalUpload = async (file, data) => {
    capi.uploadFile(folder, file, data);
    load();
    app.utils.big_modal_close();
  };

  const show_upload_panel = () => {
    app.utils.big_modal_open(PanelUploadFile, { uploadFile: finalUpload });
  };

  $: _files_selected = false;
</script>

<Layout {sources} {source}>
  <svelte:fragment slot="actions_right">
    <div class="p-2 bg-gray-50 font-sans font-thin">
      <label>
        <input type="checkbox" bind:checked={preview} />
        Preview
      </label>
    </div>

    <button
      class="px-2 py-1 rounded bg-gray-50 hover:bg-gray-200"
      on:click={show_upload_panel}
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
        />
      </svg>
    </button>

    {#if _files_selected}
      <button class="px-2 py-1 rounded bg-gray-50 hover:bg-gray-200">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
          />
        </svg>
      </button>
    {/if}
  </svelte:fragment>

  <svelte:fragment slot="body">
    {#if files_loaded && ticket_loaded}
      <table class="w-full">
        <thead>
          <tr
            class="text-md font-semibold tracking-wide text-left text-gray-900 bg-gray-100 uppercase border-b border-gray-600"
          >
            <th class="px-2 py-1" />
            <th class="px-2 py-1">Name</th>
            <th class="px-4 py-1">File Size</th>
            <th class="px-4 py-1">Last Modified</th>
            <th class="px-4 py-1">Actions</th>
          </tr>
        </thead>
        <tbody class="bg-white">
          {#each files as file}
            <tr class="text-gray-700 hover:bg-gray-200">
              <td class="px-2 py-1 border">
                {#if fapi && isImage(file.name) && preview}
                  <img
                    src={fapi.getFilePreviewUrl(file.name)}
                    alt={file.name}
                    class="h-10 border"
                  />
                {:else}
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    class="h-10 w-10 text-gray-500"
                    viewBox="0 0 20 20"
                    fill="currentColor"
                  >
                    <path
                      fill-rule="evenodd"
                      d="M4 4a2 2 0 012-2h4.586A2 2 0 0112 2.586L15.414 6A2 2 0 0116 7.414V16a2 2 0 01-2 2H6a2 2 0 01-2-2V4z"
                      clip-rule="evenodd"
                    />
                  </svg>
                {/if}
              </td>
              <td class="px-4 py-1 text-ms border">{file.name || ""}</td>
              <td class="px-4 py-1 text-xs border">
                <span
                  class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-sm"
                  >{humanizeBytes(file.size || 0)}</span
                >
              </td>
              <td class="px-4 py-1 text-sm border"
                >{(file.last_modified || "").substring(0, 19) || ""}</td
              >
              <td class="px-4 py-1 text-sm border">
                <button
                  on:click={() => app.nav.cab_file(source, folder, file.name)}
                  class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-1 px-2 border border-blue-500 rounded"
                  >Preview</button
                >
                <button
                  on:click={async () => {
                    capi.deleteFile(folder, file.name);
                    load();
                  }}
                  class="bg-red-500 hover:bg-red-700 text-white font-bold py-1 px-2 border border-red-500 rounded"
                  >Delete</button
                >
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    {:else}
      <LoadingSpinner />
    {/if}
  </svelte:fragment>
</Layout>
