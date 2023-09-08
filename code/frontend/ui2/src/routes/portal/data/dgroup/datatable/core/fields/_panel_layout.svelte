<script lang="ts">
  import { LoadingSpinner } from "$lib/core";

  export let loading = true;
  export let selected = false;
  export let onSelect;
</script>

<div class="w-full h-full flex flex-col">
  <div class="flex-grow flex flex-col h-32 p-2 space-y-1 overflow-y-auto">
    {#if loading}
      <LoadingSpinner />
    {:else}
      <slot />
    {/if}
  </div>

  <div class="flex-shrink h-10 w-full p-1 border-t flex justify-end">
    {#if selected}
      <div class="flex">
        <button
          class="p-1 shadow bg-green-500 rounded text-white font-semibold hover:bg-green-800"
          on:click={onSelect}
        >
          Select
        </button>
      </div>
    {/if}
  </div>
</div>
