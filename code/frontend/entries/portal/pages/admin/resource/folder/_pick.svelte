<script lang="ts">
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../core";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let onSelect;
  export let source = "";
  export let folder = "";

  const app = getContext("__app__") as PortalService;

  const sapi = app.api_manager.get_self_api();

  let loading = true;
  let sources = [];

  let folders = [];

  const load = async () => {
    const resp = await sapi.list_cabinet_sources();
    if (!resp.ok) {
      return;
    }
    sources = resp.data;
    if (!source) {
      source = sources[0] || "";
    }
  };

  const loadCabinet = async (src) => {
    console.log("@load_cabinet", src);
    if (!src) {
      return;
    }

    loading = true;
    folder = "";

    const capi = app.api_manager.get_cabinet(src);
    const resp = await capi.listRoot();
    if (!resp.ok) {
      return;
    }

    folders = resp.data;
    loading = false;
  };

  load();

  $: loadCabinet(source);
</script>

{#if loading}
  <LoadingSpinner classes="w-full" />
{:else}
  <div class="p-2 flex flex-row justify-end">
    <select class="p-1 rounded border" bind:value={source}>
      {#each sources as s}
        <option>{s}</option>
      {/each}
    </select>
  </div>

  <table class="w-full">
    <thead
      ><tr
        class="text-md font-semibold tracking-wide text-left text-gray-900 bg-gray-100 uppercase border-b border-gray-600"
        ><th class="px-2 py-1" /> <th class="px-2 py-1">Name</th>
        <th class="px-4 py-1">File Size</th>
        <th class="px-4 py-1">Last Modified</th>
        <th class="px-4 py-1" />
      </tr></thead
    >
    <tbody class="bg-white">
      {#each folders as _folder}
        <tr class="text-gray-700 hover:bg-gray-200"
          ><td class="px-2 py-1 border">
            <Icon name="folder" class="w-10" />
          </td>
          <td class="px-4 py-1 text-ms border">{_folder}</td>
          <td class="px-4 py-1 text-xs border"
            ><span
              class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-sm"
              >--</span
            ></td
          >
          <td class="px-4 py-1 text-sm border">
            <span
              class="px-2 py-1 font-semibold leading-tight text-green-700 bg-green-100 rounded-sm"
              >--</span
            >
          </td>

          <td class="px-4 py-1 text-sm border">
            <input
              type="checkbox"
              checked={folder === _folder}
              on:click={() => {
                if (folder === _folder) {
                  folder = "";
                } else {
                  folder = _folder;
                }
              }}
            />
          </td>
        </tr>
      {/each}
    </tbody>
  </table>

  <div class="flex justify-end mt-2">
    <button
      on:click={() => {
        onSelect(source, folder);
      }}
      class="p-1 font-bold text-white bg-blue-500 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline"
    >
      Select
    </button>
  </div>
{/if}
