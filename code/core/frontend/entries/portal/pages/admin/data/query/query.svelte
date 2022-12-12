<script lang="ts">
  import { Router } from "svelte-hash-router";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { CEditor, PortalService } from "../../core";

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
          <CEditor code="" />
        </div>
      </div>
    </div>

    <div class="flex justify-between text-sm">
      <div class="flex text-sm p-1 gap-1">
        <label class="inline-flex items-center">
          <input type="checkbox" class="form-checkbox h-5 w-5" /><span
            class="ml-2 text-gray-700">Raw Query</span
          >
        </label>

        <label class="inline-flex items-center">
            <input type="checkbox" class="form-checkbox h-5 w-5" /><span
              class="ml-2 text-gray-700">Response No Modify</span
            >
          </label>

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
  </div>

  <div class="mt-2 h-full">todo</div>
</div>
