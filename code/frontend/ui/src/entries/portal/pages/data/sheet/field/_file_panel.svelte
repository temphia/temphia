<script lang="ts">
  import type { FolderTktAPI } from "../../../../../../lib/apiv2";
  import { LoadingSpinner } from "../../../admin/core";

  export let folder_api: FolderTktAPI;
  export let onFileAdd = (file) => {};

  let loading = true;
  let files = [];

  const load = async () => {
    const resp = await folder_api.list();
    if (!resp.ok) {
      return;
    }
    files = resp.data;
    loading = false;
  };
  load();
</script>

<div class="h-full w-full bg-gray-50">
  {#if loading}
    <LoadingSpinner classes="" />
  {:else}
    <div class="w-full flex gap-1 flex-wrap">
      {#each files as file}
        <div
          class="w-24 overflow-hidden border hover:border-blue-500 cursor-pointer border-gray-100 rounded"
        >
          <img
            class="h-auto w-24 p-1 rounded"
            src={folder_api.getFilePreviewUrl(file.name)}
            alt=""
          />
          <span
            class="text-xs text-gray-600 truncate overflow-hidden whitespace-nowrap"
          >
            {file.name}</span
          >
          <button
            on:click={() => onFileAdd(file.name)}
            class="hover:bg-blue-500 w-full rounded text-gray-600 hover:text-white  font-semibold text-xs bg-blue-100"
            >Add</button
          >
        </div>
      {/each}
    </div>
  {/if}
</div>
