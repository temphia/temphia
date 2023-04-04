<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import LensHelp from "./_help.svelte";
  import { getContext } from "svelte";
  import { CEditor, PortalService } from "../../core";

  export let fromDate = null;
  export let toDate = null;
  export let loading = false;
  export let message;
  export let do_query = (qstr) => {};

  const app = getContext("__app__") as PortalService;

  const opts = app.nav.options || {};
  let code = "{}";

  if (opts["filters"]) {
    code = JSON.stringify(opts["filters"], null, 4);
  }

  let editor;
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
          <CEditor bind:editor {code} />
        </div>
      </div>
    </div>

    <div class="flex flex-wrap justify-between text-sm">
      <div class="flex flex-wrap text-sm p-1 gap-1">
        <div
          class="bg-green-100 p-1 rounded-lg flex space-x-1 flex-row items-center cursor-pointer"
        >
          <Icon name="chevron-left" class="h-4" />
          <input
            type="datetime-local"
            class="text-xs bg-green-100"
            bind:value={fromDate}
          />
        </div>

        <div
          class="bg-red-100 p-1 rounded-lg flex space-x-1 flex-row items-center cursor-pointer"
        >
          <input
            type="datetime-local"
            class="text-xs bg-red-100"
            bind:value={toDate}
          />
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

      <div class="flex justify-end p-1 gap-1">
        <button
          on:click={() => {
            do_query(editor.getValue());
          }}
          class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2 flex"
        >
          {#if loading}
            <Icon name="globe" class="h-4 w-4 animate-bounce" solid />
          {/if}

          Search
        </button>

        <button
          on:click={() => {
            app.utils.small_modal_open(LensHelp, {});
          }}
          class="p-1 rounded bg-blue-300 shadow hover:bg-blue-600 flex text-white"
        >
          <Icon name="information-circle" class="h-6 w-6" solid />
        </button>
      </div>
    </div>

    <div>
      <p class="text-red-500">{message}</p>
    </div>
  </div>

  <div class="mt-2 h-full">
    <slot />
  </div>
</div>
