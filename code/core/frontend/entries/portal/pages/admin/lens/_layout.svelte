<script lang="ts">
  import { Router } from "svelte-hash-router";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import LensHelp from "./_help.svelte";
  import { getContext } from "svelte";
  import type { PortalService } from "../core";

  export let index;

  const app = getContext("__app__") as PortalService;

  export let do_query = (qstr) => {};
</script>

<div class="p-2 w-full h-full" style="min-height: 80vh;">
  <div class="rounded bg-white p-2">
    <div class="flex items-start">
      <div class="p-1 flex-grow">
        <label
          for="default-search"
          class="mb-2 text-sm font-medium text-gray-900 sr-only ">Search</label
        >
        <div class="relative">
          <textarea
            type="search"
            id="default-search"
            class="block p-2 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 "
            placeholder="Search .."
            required
          />

          <div class="flex text-sm p-1 gap-1">
            <div
              class="bg-green-100 p-1 rounded-lg flex space-x-2 flex-row cursor-pointer"
            >
              <Icon name="chevron-left" class="h-4" />
              <p class="text-xs">2021-01-15 [5:30]</p>
              <Icon name="calendar" class="h-4" />
            </div>

            <div
              class="bg-red-100 p-1 rounded-lg flex space-x-2 flex-row cursor-pointer"
            >
              <Icon name="calendar" class="h-4" />
              <p class="text-xs">2021-02-17 [7:30]</p>
              <Icon name="chevron-right" class="h-4" />
            </div>
          </div>
        </div>
      </div>

      <div class="flex-none flex flex-col p-1 gap-1">
        <div class="flex gap-1">
          <select
            class="font-medium rounded text-sm px-1 py-2"
            value={index}
            on:change={(ev) => app.nav.admin_lens(ev.target["value"])}
          >
            <option value="site">Site Index</option>
            <option value="engine">Engine Index</option>
            <option value="app">App Index</option>
          </select>

          <button
            on:click={() => {}}
            class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
          >
            <Icon name="information-circle" class="h-6 w-6" solid />
          </button>
        </div>

        <button
          on:click={() => do_query("")}
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2"
          >Search</button
        >
      </div>
    </div>
  </div>

  <div class="mt-2 h-full">
    <slot />
  </div>
</div>
