<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import type { PortalService } from "../services";

  export let show_running = false;

  const app = getContext("__app__") as PortalService;

  const user_apps = app.api_manager.self_data.get_user_apps();
</script>

<div
  class="w-full h-full py-10 mx-auto overflow-auto bg-gradient-to-b from-purple-100 to-indigo-100"
>
  <div class="flex justify-center w-full mb-10">
    <div class="relative">
      <input
        type="text"
        class="h-14 w-96 pr-8 pl-5 rounded z-0 focus:shadow focus:outline-none"
        placeholder="Search anything..."
      />
      <div class="absolute top-4 right-3 flex">
        <svg class="h-5 w-5 text-gray-500" viewBox="0 0 24 24" fill="none"
          ><path
            d="M21 21L15 15M17 10C17 13.866 13.866 17 10 17C6.13401 17 3 13.866 3 10C3 6.13401 6.13401 3 10 3C13.866 3 17 6.13401 17 10Z"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          /></svg
        >
      </div>
    </div>
  </div>

  <div class="p-8 flex flex-wrap justify-center gap-4">
    {#each user_apps as uapp}
      <div
        on:click={() => app.nav.launch_target(uapp["target_id"], uapp["name"])}
        class="bg-white flex flex-col items-center h-32 w-32 p-2 overflow-hidden shadow-lg rounded-lg cursor-pointer hover:border-2 border-blue-400"
      >
        {#if !uapp["icon"]}
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-28 h-28 text-gray-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 110-4h1a1 1 0 001-1V7a1 1 0 011-1h3a1 1 0 001-1V4z"
            />
          </svg>
        {:else}
          <div class="w-28 h-28">
            {@html uapp["icon"]}
            <!-- fixme => escape contents -->
          </div>
        {/if}
        
        <h2 class="text-lg text-gray-500 font-semibold font-mono">
          {uapp["name"]}
        </h2>
      </div>
    {/each}
  </div>
</div>

{#if show_running}
  <div class="fixed bottom-10 right-10">
    <button
      on:click={() => app.nav.launcher()}
      class="p-0 w-12 h-12 bg-blue-500 rounded-full hover:bg-blue-900 active:shadow-lg mouse shadow transition ease-in duration-200 focus:outline-none"
    >
      <Icon name="lightning-bolt" class="text-white p-1" solid />
    </button>
  </div>
{/if}
