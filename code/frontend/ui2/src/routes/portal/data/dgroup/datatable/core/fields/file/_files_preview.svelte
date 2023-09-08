<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { FolderTktAPI } from "$lib/services/apiv2";

  export let folder_api: FolderTktAPI;
  export let files: string[];
  export let current_file: string;

  $: _current_file = current_file;
  $: _preview_url = folder_api.getFilePreviewUrl(_current_file)

  const imagExts = ["jpg", "png", "webp", "svg"];
  const isImage = (file: string) => imagExts.includes(file.split(".").pop());
</script>

<div class="w-full h-full flex flex-col">
  <div class="w-full flex justify-center">
    <div
      class="flex absolute bottom-5 left-1/2 z-30 space-x-3 -translate-x-1/2"
    >
      {#each files as file}
        <button
          type="button"
          class="w-3 h-3 rounded-full  border border-gray-500 {file ===
          _current_file
            ? 'bg-white'
            : 'bg-gray-300'}"
          on:click={() => {
            _current_file = file;
          }}
        />
      {/each}
    </div>
    <!-- Slider controls -->
    <button
      type="button"
      class="flex absolute top-0 left-0 z-30 justify-center items-center px-4 h-full cursor-pointer group focus:outline-none"
      data-carousel-prev
    >
      <span
        class="inline-flex justify-center items-center w-8 h-8 rounded-full sm:w-10 sm:h-10 bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
      >
        <Icon
          name="chevron-left"
          class="w-5 h-5 text-white sm:w-6 sm:h-6 dark:text-gray-800"
        />

        <span class="hidden">Previous</span>
      </span>
    </button>
    <button
      type="button"
      class="flex absolute top-0 right-0 z-30 justify-center items-center px-4 h-full cursor-pointer group focus:outline-none"
      data-carousel-next
    >
      <span
        class="inline-flex justify-center items-center w-8 h-8 rounded-full sm:w-10 sm:h-10 bg-white/30 dark:bg-gray-800/30 group-hover:bg-white/50 dark:group-hover:bg-gray-800/60 group-focus:ring-4 group-focus:ring-white dark:group-focus:ring-gray-800/70 group-focus:outline-none"
      >
        <Icon
          name="chevron-right"
          class="w-5 h-5 text-white sm:w-6 sm:h-6 dark:text-gray-800"
        />

        <span class="hidden">Next</span>
      </span>
    </button>

    {#key _current_file}
      {#if isImage(_current_file)}
        <img
          src={_preview_url}
          alt=""
          class="rounded p-1 border"
        />
      {:else}
        <div>AA</div>
      {/if}
    {/key}
  </div>
  <div class="flex">
    <div />
  </div>
</div>
