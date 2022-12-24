<script lang="ts">
  import { Router } from "svelte-hash-router";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import LensHelp from "./_help.svelte";
  import { getContext } from "svelte";
  import type { PortalService } from "../../core";

  export let index;

  const app = getContext("__app__") as PortalService;

  export let do_query = (qstr) => {};
</script>

<div class="flex justify-end pt-2 pr-2">
  <select
    class="px-2 py-1 rounded-full bg-white hover:text-white hover:bg-slate-500 border border-slate-600"
    value={index}
    on:change={(ev) => app.nav.admin_lens(ev.target["value"])}
  >
    <option value="site">Site Index</option>
    <option value="engine">Engine Index</option>
    <option value="app">App Index</option>
  </select>
</div>

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
        </div>
      </div>

      <div class="flex p-1 gap-1">
        <button
          on:click={() => do_query("")}
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2"
          >Search</button
        >

        <button
          on:click={() => {}}
          class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
        >
          <Icon name="information-circle" class="h-6 w-6" solid />
        </button>
      </div>
    </div>

    <div class="flex flex-wrap justify-between text-sm">
      <div class="flex text-sm p-1 gap-1">
        <div
          class="bg-green-100 p-1 rounded-lg flex space-x-1 flex-row cursor-pointer"
        >
          <Icon name="chevron-left" class="h-4" />
          <input type="datetime-local" class="text-xs" />
          <!-- <p class="text-xs">2021-01-15 [5:30]</p> -->
          <Icon name="calendar" class="h-4" />
        </div>

        <div
          class="bg-red-100 p-1 rounded-lg flex space-x-1 flex-row cursor-pointer"
        >
          <Icon name="calendar" class="h-4" />
          <input type="datetime-local" class="text-xs" />
          <!-- <p class="text-xs">2021-02-17 [7:30]</p> -->
          <Icon name="chevron-right" class="h-4" />
        </div>

        <label class="inline-flex items-center">
          <input type="checkbox" class="form-checkbox h-5 w-5" checked /><span
            class="ml-2 text-gray-700">Unscope Organization</span
          >
        </label>

        <label class="inline-flex items-center">
          <select class="font-medium rounded text-sm p-1">
            <option value="50">50</option>
            <option value="100">100</option>
            <option value="500">500</option>
            <option value="1000">1000</option>
          </select>

          <span class="ml-2 text-gray-700">Logs count</span>
        </label>
      </div>
    </div>
  </div>

  <div class="mt-2 h-full">
    <slot />
  </div>
</div>
