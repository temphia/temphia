<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import Layout from "../layout.svelte";
  import LensHelp from "./_lens_help.svelte";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";

  const app: PortalApp = getContext("__app__");
</script>

<Layout>
  <div class="p-2">
    <div class="rounded bg-white p-2">
      <form class="p-2 border-b border-slate-200">
        <label
          for="default-search"
          class="mb-2 text-sm font-medium text-gray-900 sr-only ">Search</label
        >
        <div class="relative">
          <div
            class="flex absolute inset-y-0 left-0 items-center pl-3 pointer-events-none"
          >
            <svg
              class="w-5 h-5 text-gray-500 dark:text-gray-400"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
              ><path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"
              /></svg
            >
          </div>
          <input
            type="search"
            id="default-search"
            class="block p-4 pl-10 w-full text-sm text-gray-900 bg-gray-50 rounded-lg border border-gray-300 focus:ring-blue-500 focus:border-blue-500 "
            placeholder="Search .."
            required
          />

          <div class="absolute right-2.5 bottom-2.5 flex gap-2">
            <select class="font-medium rounded text-sm px-1 py-2">
              <option>Site Index</option>
              <option>Engine Index</option>
              <option>App Index</option>
            </select>

            <button
              type="submit"
              class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2"
              >Search</button
            >
            <button
              type="button"
              on:click={() => {
                app.big_modal_open(LensHelp, {});
              }}
              class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
            >
              <Icon name="information-circle" class="h-6 w-6" solid />
            </button>
          </div>
        </div>
      </form>

      <div class="flex flex-col" style="min-height: 60vh;">
        <div class="grow">.</div>
        <div class="flex-none h-8 w-full bg-slate-50 rounded flex">
          <div class="flex text-sm p-1">
            <Icon name="calendar" class="h-5 w-5" />
            <span class="text-zinc-800">From</span> 2022/04/04 [6:33]
            <span class="text-zinc-800">To</span> 2022/04/05 [6:33]
          </div>
        </div>
      </div>
    </div>
  </div>
</Layout>
