<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "$lib/core";
  import { LoadingSpinner } from "$lib/core";

  export let template = "";

  const app: PortalService = getContext("__app__");
  let loading = true;

  const sapi = app.api_manager.get_self_api();

  let templates = [];
  const load = async () => {
    const resp = await sapi.list_sheet_templates();
    if (!resp.ok) {
      return;
    }

    templates = resp.data;
    loading = false;
    console.log("@templates", templates);
  };

  load()


</script>

<div class="flex flex-wrap p-2 gap-2 overflow-auto justify-center">
  {#if loading}
    <LoadingSpinner classes="" />
  {:else}
    {#each Object.entries(templates) as [tplkey, tpl]}
      <!-- svelte-ignore a11y-click-events-have-key-events -->
      <div class="p-4 w-52 m-1 shadow rounded hover:bg-blue-100   {tplkey === template ? "border border-blue-400" : ""}" on:click={() => {
        template = tplkey
      }}>
        <div class="block overflow-hidden rounded">
          <img
            alt=""
            class="block w-full h-auto object-cover object-center cursor-pointer"
            src="https://dummyimage.com/400x200"
          />
        </div>
        <div class="mt-4">
          <h2 class="title-font text-lg font-medium text-gray-900">
            {tpl.name}
          </h2>
        </div>
      </div>
    {/each}
  {/if}
</div>
