<script lang="ts">
  import Layout from "./_layout.svelte";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../lib/app/portal";

  export let source;
  export let folder;
  export let file;

  const { open, close } = getContext("simple-modal");
  const app: PortalApp = getContext("__app__");

  let sources = [];
  app.get_cabinet_sources().then((sresp) => {
    sources = sresp;
  });

  let fapi;
  let ticket_loaded = false;
  app.get_folder_api(source, folder).then((resp) => {
    fapi = resp;
    ticket_loaded = true;
  });
</script>

<Layout {source} {sources} >
  <svelte:fragment slot="actions_left">
    <button class="hover:bg-gray-200 rounded">
      <svg
        class="h-6 w-6"
        viewBox="0 0 24 24"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M14.0303 7.46967C14.3232 7.76256 14.3232 8.23744 14.0303 8.53033L10.5607 12L14.0303 15.4697C14.3232 15.7626 14.3232 16.2374 14.0303 16.5303C13.7374 16.8232 13.2626 16.8232 12.9697 16.5303L8.96967 12.5303C8.67678 12.2374 8.67678 11.7626 8.96967 11.4697L12.9697 7.46967C13.2626 7.17678 13.7374 7.17678 14.0303 7.46967Z"
          fill="black"
        />
      </svg>
    </button>
    <button class="hover:bg-gray-200 rounded">
      <svg
        class="h-6 w-6"
        viewBox="0 0 24 24"
        fill="none"
        xmlns="http://www.w3.org/2000/svg"
      >
        <path
          fill-rule="evenodd"
          clip-rule="evenodd"
          d="M9.96967 7.46967C10.2626 7.17678 10.7374 7.17678 11.0303 7.46967L15.0303 11.4697C15.3232 11.7626 15.3232 12.2374 15.0303 12.5303L11.0303 16.5303C10.7374 16.8232 10.2626 16.8232 9.96967 16.5303C9.67678 16.2374 9.67678 15.7626 9.96967 15.4697L13.4393 12L9.96967 8.53033C9.67678 8.23744 9.67678 7.76256 9.96967 7.46967Z"
          fill="black"
        />
      </svg>
    </button>
  </svelte:fragment>

  <svelte:fragment slot="actions_right">
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
          d="M4 16v1a3 3 0 003 3h10a3 3 0 003-3v-1m-4-8l-4-4m0 0L8 8m4-4v12"
        />
      </svg>
    </button>
  </svelte:fragment>

  <div class="flex" slot="body">
    <div class="m-auto p-4">
      {#if ticket_loaded}
        <img
          class="p-1 border bg-white"
          src={fapi.get_file_preview_link(file)}
          alt=""
        />
      {:else}
        <div>Loading</div>
      {/if}
    </div>
  </div>
</Layout>
