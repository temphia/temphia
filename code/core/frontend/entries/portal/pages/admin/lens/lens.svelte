<script lang="ts">
  import { Router } from "svelte-hash-router";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import LensHelp from "./_help.svelte";
</script>

<div class="p-2">
  <div class="rounded bg-white p-2">
    <div class="flex items-start border-b border-slate-200">
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
            <div class="bg-green-100 p-1 rounded-lg flex space-x-2 flex-row cursor-pointer">
              <Icon name="chevron-left" class="h-4" />
              <p class="text-xs">2021-01-15 [5:30]</p>
              <Icon name="calendar" class="h-4" />
            </div>
            
            <div class="bg-red-100 p-1 rounded-lg flex space-x-2 flex-row cursor-pointer">
              <Icon name="calendar" class="h-4" />
              <p class="text-xs">2021-02-17 [7:30]</p>
              <Icon name="chevron-right" class="h-4" />
            </div>
          </div>
        </div>
      </div>

      <div class="flex-none flex flex-col p-1 gap-1">
        <div class="flex gap-1">
          <select class="font-medium rounded text-sm px-1 py-2">
            <option>Site Index</option>
            <option>Engine Index</option>
            <option>App Index</option>
          </select>

          <button
            on:click={() => {}}
            class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
          >
            <Icon name="information-circle" class="h-6 w-6" solid />
          </button>
        </div>

        <button
          type="submit"
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2"
          >Search</button
        >
      </div>
    </div>

    <div class="flex flex-col" style="min-height: 60vh;">
      <div class="grow">.</div>
      <div class="flex-none h-8 w-full bg-slate-50 rounded flex" />
    </div>
  </div>
</div>

<Router />
