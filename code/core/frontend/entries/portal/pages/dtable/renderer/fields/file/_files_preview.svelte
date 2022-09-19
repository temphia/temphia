<script lang="ts">
  import type { FolderTktAPI } from "../../../../../../../lib/core/tktapi";

  export let folder_api: FolderTktAPI;
  export let files: string[];
  export let current_file: string;

  $: _current_file = current_file;

  const imagExts = ["jpg", "png", "webp", "svg"];
  const isImage = (file: string) => imagExts.includes(file.split(".").pop());
</script>

<div class="w-full h-full flex flex-col">
  <div class="w-full flex justify-center">
    {#if isImage(_current_file)}
      <img
        src={folder_api.get_file_preview_link(current_file)}
        alt=""
        class="rounded p-1 border"
      />
    {:else}
      <div>AA</div>
    {/if}
  </div>
  <div class="flex">
    <div />
  </div>
</div>
