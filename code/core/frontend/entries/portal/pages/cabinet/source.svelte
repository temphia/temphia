<script lang="ts">
  import { getContext } from "svelte";
  import Layout from "./_layout.svelte";
  import PanelNewFolder from "./panels/new_folder.svelte";
  import type { PortalApp } from "../../app";

  export let source;

  const app: PortalApp = getContext("__app__");

  let sources = [];
  app.get_cabinet_sources().then((sresp) => {
    sources = sresp;
  });

  const { open } = getContext("simple-modal");
  let folders = [];

  const load = async () => {
    const api = await app.get_apm().get_cabinet_api(source);
    const resp = await api.list_root();
    folders = await resp.data;
  };

  const complete_new_folder = async (name) => {
    const api = await app.get_apm().get_cabinet_api(source);
    api.new_folder(name);
  };

  const new_folder = () => {
    open(PanelNewFolder, { onNewName: complete_new_folder });
  };

  load();

  // actions
</script>

<Layout {sources}>
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
    <button
      class="px-2 py-1 rounded bg-gray-50 hover:bg-gray-200"
      on:click={new_folder}
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
          d="M9 13h6m-3-3v6m-9 1V7a2 2 0 012-2h6l2 2h6a2 2 0 012 2v8a2 2 0 01-2 2H5a2 2 0 01-2-2z"
        />
      </svg>
    </button>

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
          d="M4 4v5h.582m15.356 2A8.001 8.001 0 004.582 9m0 0H9m11 11v-5h-.581m0 0a8.003 8.003 0 01-15.357-2m15.357 2H15"
        />
      </svg>
    </button>
  </svelte:fragment>

  <svelte:fragment slot="body">
    <div class="flex flex-wrap space-x-4 space-y-2 p-2">
      {#each folders as folder}
        <div
          class="flex flex-col p-1 border rounded-lg bg-white hover:bg-gray-100 cursor-pointer"
          on:click={() => app.navigator.goto_cabinet_folder(source, folder)}
        >
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="h-20 w-20 text-gray-500"
            viewBox="0 0 20 20"
            fill="currentColor"
          >
            <path
              d="M2 6a2 2 0 012-2h5l2 2h5a2 2 0 012 2v6a2 2 0 01-2 2H4a2 2 0 01-2-2V6z"
            />
          </svg>

          <div class="flex justify-center">
            <span class="text-grey-600 text-lg truncate font-medium"
              >{folder}</span
            >
          </div>
        </div>
      {/each}
    </div>
  </svelte:fragment>
</Layout>
