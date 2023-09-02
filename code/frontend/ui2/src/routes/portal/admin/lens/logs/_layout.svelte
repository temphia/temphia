<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import LensHelp from "./_help.svelte";
  import { getContext } from "svelte";
  import { CEditor, PortalService } from "$lib/core";

  export let fromDate = null;
  export let toDate = null;
  export let loading = false;
  export let count = "20";
  export let message;
  export let do_query = (qstr) => {};

  const app = getContext("__app__") as PortalService;

  const opts = app.nav.options || {};
  let code = "{}";

  if (opts["filters"]) {
    code = JSON.stringify(opts["filters"], null, 4);
  }

  const subtractTimeFromDate = (objDate, sec) => {
    const numberOfMlSeconds = objDate.getTime();
    const offset = objDate.getTimezoneOffset() * 60 * 1000;
    const addMlSeconds = sec * 1000;
    return new Date((numberOfMlSeconds - offset) - (addMlSeconds ));
  };

  const date_range = {
    "5min": ["Last 5 minute", 5 * 60],
    "15min": ["Last 15 minute", 15 * 60],
    "1hour": ["Last hour", 60 * 60],
    "1day": ["Last day", 60 * 60 * 24],
    "1week": ["Last week", 60 * 60 * 24 * 7],
    "15days": ["Last 15 days", 60 * 60 * 24 * 15],
    "1mon": ["Last 1 month", 60 * 60 * 24 * 30],
  };

  let range;
  const reset_range = () => {
    console.log("@reset_value", fromDate);
    range = "custom";
  };

  $: if (date_range[range]) {
    console.log("@range", date_range[range]);
    fromDate = subtractTimeFromDate(new Date(), date_range[range][1])
      .toISOString()
      .slice(0, 16);
    console.log("@range_from", fromDate);
  }

  let editor;
</script>

<div class="p-2 w-full h-full" style="min-height: 80vh;">
  <div class="rounded bg-white p-2">
    <div class="flex items-start">
      <div class="p-1 flex-grow">
        <label
          for="default-search"
          class="mb-2 text-sm font-medium text-gray-900 sr-only">Search</label
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

          {#key range}
            {#key fromDate}
              <input
                type="datetime-local"
                on:change={reset_range}
                class="text-xs bg-green-100"
                bind:value={fromDate}
              />
            {/key}
          {/key}
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
          <select class="font-medium rounded text-sm p-1" bind:value={count}>
            <option value="20">20</option>
            <option value="50">50</option>
            <option value="100">100</option>
            <option value="500">500</option>
            <option value="1000">1000</option>
          </select>

          <span class="ml-2 text-gray-700">Logs count</span>
        </label>
      </div>

      <div class="flex justify-end p-1 gap-1">
        <label class="inline-flex items-center">
          <select bind:value={range} class="font-medium rounded text-sm p-1">
            <option value="custom" />

            {#each Object.entries(date_range) as [dkey, dval]}
              <option value={dkey}>{dval[0]}</option>
            {/each}
          </select>
        </label>

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

  <div class="mt-2 w-full">
    <slot />
  </div>
</div>
