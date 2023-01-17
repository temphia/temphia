<script lang="ts">
  import Layout from "./_layout.svelte";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import { params } from "svelte-hash-router";

  import PanelUploadFile from "./panels/upload_file.svelte";
  import { isImage } from "../../../../lib/utils";
  import type { FolderTktAPI } from "../../../../lib/apiv2";

  const app: PortalService = getContext("__app__");

  const source = $params.source;
  const folder = $params.folder;
  const file = $params._;
  let fapi: FolderTktAPI;

  const cabservice = app.get_cabinet_service();
  const capi = cabservice.get_source_api(source);

  let sources = [];
  cabservice.get_cab_sources().then((sresp) => {
    sources = sresp;
  });

  const load = async () => {
    fapi = await cabservice.get_folder_api(source, folder);
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
</script>

<Layout {sources} {source}>
  <svelte:fragment slot="actions_right">
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
  </svelte:fragment>

  <svelte:fragment slot="body">
    {#if fapi && isImage(file)}
      <img
        class="w-full"
        src={fapi.getFileUrl(file)}
        alt=""
        srcset=""
      />
    {:else}
      <div class="flex justify-center py-10">
        <button class="p-2 rounded bg-green-400 hover:bg-green-600 font-sans text-white">Download</button>
      </div>
    {/if}
  </svelte:fragment>
</Layout>
