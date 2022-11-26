<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";

  const app: PortalService = getContext("__app__");

  export let source = "";
  export let path = "";
  export let sources = [];
</script>

<div class="w-full h-screen p-2">
  <div class="h-full w-full flex flex-col rounded border">
    <header class="flex-shrink h-12 border-b pb-1">
      <div class="flex justify-between border-b p-1">
        <div class="flex">
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
        </div>

        <div class="flex">
          <slot name="actions_right" />
          <select
            class="rounded p-2 bg-white border"
            value={source}
            on:change={(ev) => app.nav.cab_folders(ev.target["value"])}
          >
            {#each sources as src}
              <option value={src}> {src}</option>
            {/each}
          </select>
        </div>
      </div>
    </header>
    <div class="flex-grow h-32 p-2 overflow-auto">
      <div class="bg-gray-50 h-full w-full rounded">
        <slot name="body" />
      </div>
    </div>
  </div>
</div>
