<script lang="ts">
  import type { FolderTktAPI } from "$lib/services/apiv2";
  import type { RowService } from "$lib/services/data";

  export let row_editor: RowService;
  export let folder_api: FolderTktAPI;

  export let onSelect: (file: string) => void;

  let files = [];
  const load = async () => {
    const resp = await folder_api.list()
    if (!resp.ok) {
      return
    }
    files = resp.data;
  };

  load();

  let selectd = "";
</script>

<div class="w-full h-full flex flex-col">
  <div class="grow flex flex-wrap space-x-2 space-y-2 overflow-auto items-start">
    {#each files as file}
      <div
        on:click={() => {
          selectd = file.name;
        }}
        class="flex flex-col p-1 border rounded-lg bg-white hover:bg-gray-200 cursor-pointer {selectd ===
        file.name
          ? 'bg-gray-400'
          : ''}"
      >
        <img
          src={folder_api.getFilePreviewUrl(file.name)}
          class="w-32 border"
          alt=""
        />

        <div class="flex justify-center">
          <span class="text-grey-600 text-lg w-28 truncate font-medium"
            >{file.name}</span
          >
        </div>
      </div>
    {/each}
  </div>

  <div class="flex grow-0 justify-between pt-4 border-t">
    <div class="text-gray-600">
      {selectd === "" ? "Select a file" : `File Selected ${selectd}`}
    </div>

    <div>
      <button
        on:click={load}
        class="px-2 py-1 shadow bg-blue-500 hover:bg-blue-700 rounded text-white font-semibold"
      >
        Refresh
      </button>

      {#if selectd !== ""}
        <button
          on:click={() => onSelect(selectd)}
          class="px-2 py-1 shadow bg-green-500 hover:bg-green-700 rounded text-white font-semibold"
        >
          Select
        </button>
      {/if}
    </div>
  </div>
</div>
