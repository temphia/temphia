<script lang="ts">
  import type { FolderTktAPI } from "../../../../../../../lib/core/tktapi";
  import type { RowEditor } from "../../../../../../../lib/service/dyn";

  export let row_editor: RowEditor;
  export let folder_api: FolderTktAPI;

  export let onSelect: (file: string) => void;

  let files = [];
  const load = async () => {
    files = await folder_api.list();
  };

  load();

  let selectd = "";
</script>

<div class="w-full">
  <div
    class="flex flex-wrap space-x-2 space-y-2 overflow-auto"
    style="max-height: 80vh;"
  >
    {#each files as file}
      <div
        on:click={() => {
          selectd = file.name;
        }}
        class="flex flex-col p-1 border rounded-lg bg-white hover:bg-gray-600 cursor-pointer {selectd ===
        file.name
          ? 'bg-gray-400'
          : ''}"
      >
        <img
          src={folder_api.get_file_preview_link(file.name)}
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

  <div class="flex justify-between pt-4 border-t">
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
