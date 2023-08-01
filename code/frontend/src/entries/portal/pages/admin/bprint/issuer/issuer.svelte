<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { PortalService } from "../../core";
  import IssuerResult from "./issuer_result.svelte";

  export let bid: string;
  export let app: PortalService;

  const bapi = app.api_manager.get_admin_bprint_api();

  let plugs = [];
  let all_plugs = false;
  let selected_plugs = [];

  const load = async () => {
    const bapi = app.api_manager.get_admin_bprint_api();
    const resp = await bapi.list_plugs(bid);
    if (resp.status !== 200) {
      return;
    }
    plugs = resp.data;
  };

  load();
</script>

<h4 class="mb-3 block text-base font-medium uppercase text-gray-500">
  Select plugs to include.
</h4>

<table
  class="min-w-full divide-y divide-gray-200 table-fixed dark:divide-gray-700 border rounded"
>
  <thead class="bg-gray-100 dark:bg-gray-700">
    <tr>
      <th scope="col" class="p-4" />
      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
      >
        Id
      </th>
      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
      >
        Name
      </th>
      <th
        scope="col"
        class="py-3 px-6 text-xs font-medium tracking-wider text-left text-gray-700 uppercase dark:text-gray-400"
      >
        Live
      </th>
      <th scope="col" class="p-4" />
    </tr>
  </thead>
  <tbody
    class="bg-white divide-y divide-gray-200 dark:bg-gray-800 dark:divide-gray-700"
  >
    {#each plugs as plug}
      <tr class="hover:bg-gray-100 dark:hover:bg-gray-700">
        <td class="p-4 w-4">
          <div class="flex items-center">
            <input
              id="checkbox-table-{plug.id}"
              type="checkbox"
              on:click={() => {
                if (selected_plugs.includes(plug.id)) {
                  selected_plugs = selected_plugs.filter((v) => v !== plug.id);
                } else {
                  selected_plugs = [plug.id, ...selected_plugs];
                }
              }}
              checked={selected_plugs.includes(plug.id)}
              disabled={all_plugs}
              class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
            />
            <label for="checkbox-table-{plug.id}" class="sr-only"
              >checkbox</label
            >
          </div>
        </td>
        <td
          class="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white"
          >{plug["id"]}</td
        >
        <td
          class="py-4 px-6 text-sm font-medium text-gray-500 whitespace-nowrap dark:text-white"
          >{plug["name"]}</td
        >
        <td
          class="py-4 px-6 text-sm font-medium text-gray-900 whitespace-nowrap dark:text-white"
          >{plug["live"]}</td
        >
        <td class="py-4 px-6 text-sm font-medium text-right whitespace-nowrap">
          <a href="#" class="text-blue-600 dark:text-blue-500 hover:underline"
            >Explore</a
          >
        </td>
      </tr>
    {/each}
  </tbody>
</table>

<div class="flex justify-start gap-2 p-2">
  <input
    id="checkbox-all"
    type="checkbox"
    bind:checked={all_plugs}
    class="w-4 h-4 text-blue-600 bg-gray-100 rounded border-gray-300 focus:ring-blue-500 dark:focus:ring-blue-600 dark:ring-offset-gray-800 focus:ring-2 dark:bg-gray-700 dark:border-gray-600"
  />
  <label for="checkbox-all" class="text-gray-800">All Plugs</label>
</div>

<p class="text-sm text-gray-500">
  That includes all future plugs instanced from this blueprint.
</p>

{#if all_plugs || selected_plugs.length != 0}
  <div class="flex justify-end">
    <button
      class="p-1 flex text-white font-semibold rounded bg-green-400 hover:bg-green-600"
      on:click={async () => {
        const resp = await bapi.issue(bid, {
          all_plugs,
          plug_ids: selected_plugs,
        });
        if (!resp.ok) {
          console.log("Err", resp);
          return;
        }

        app.utils.small_modal_open(IssuerResult, {
          ticket: resp.data,
          close: app.utils.small_modal_close,
        });
      }}
    >
      <Icon name="terminal" class="h-5 w-5" />
      &nbsp; Issue
    </button>
  </div>
{/if}
