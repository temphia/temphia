<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import Layout from "./_layout.svelte";
  import PanelNewFolder from "./panels/new_folder.svelte";

  import { params } from "svelte-hash-router";

  const app: PortalService = getContext("__app__");
  const cservice = app.get_cabinet_service();
  const capi = cservice.get_source_api($params.source);

  let sources = [];
  let folders = [];
  let loading = false;

  cservice.get_cab_sources().then((sresp) => {
    sources = sresp;
  });

  const load = async () => {
    const resp = await capi.listRoot();
    if (!resp.ok) {
      return;
    }

    folders = resp.data;
    loading = true;
  };

  const complete_new_folder = async (name) =>
    capi.newFolder($params.folder, name);

  const new_folder = () => {
    app.utils.small_modal_open(PanelNewFolder, {
      onNewName: complete_new_folder,
    });
  };

  load();

  // actions
</script>

<Layout {sources}>
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
          on:click={() => app.nav.cab_folder($params.source, folder)}
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
