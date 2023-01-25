<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { strHash } from "../../../../lib/utils";
  import type { PortalService } from "../../services";
  import { LoadingSpinner } from "../admin/core";

  export let source = $params.source;

  const app: PortalService = getContext("__app__");
  const api = app.api_manager.get_admin_data_api();

  let loading = false;
  let sources = [];
  let groups = [];

  app.api_manager.self_data.get_data_sources().then((_sources) => {
    sources = _sources;
  });

  const load = async () => {
    const resp = await api.list_group(source);
    if (!resp.ok) {
      return;
    }

    groups = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="w-full h-full bg-gray-100 overflow-auto text-gray-600 body-font">
    <div class="container px-2 py-8 mx-auto max-w-7x1">
      <div class="flex justify-around w-full mb-4 p-4">
        <div class="w-full mb-6 lg:mb-0">
          <h2 class="text-gray-700 text-3xl font-medium">Data Table Groups</h2>
          <div class="h-1 w-64 bg-indigo-500 rounded" />
        </div>

        <div>
          <select
            class="rounded p-2 bg-white"
            value={source}
            on:change={(ev) => {}}
          >
            {#each sources as src}
              <option>{src}</option>
            {/each}
          </select>
        </div>
      </div>
      <div class="flex flex-wrap -m-4">
        {#each groups as group}
          <div class="xl:w-1/3 md:w-1/2 p-4">
            <div
              class="bg-white p-6 rounded-lg border hover:border-blue-400 shadow"
            >
              <img
                class="lg:h-60 xl:h-56 md:h-64 sm:h-72 xs:h-72 h-72  rounded w-full object-cover object-center mb-6"
                src="https://picsum.photos/seed/{strHash(
                  group.name + group.slug
                )}d/800/400"
                alt=""
              />
              <h3
                class="tracking-widest text-indigo-500 text-xs uppercase font-medium title-font"
              >
                {source}
              </h3>
              <h2 class="text-lg text-gray-900 font-medium title-font mb-4">
                {group.name}
              </h2>
              <p class="leading-relaxed selection:bg-red-200 text-base">
                {group.description}
              </p>
              <div class="flex p-5">
                <button
                  on:click={() => app.nav.data_render_table_loader(source, group["slug"])}
                  class="text-center text-sm bg-blue-500 rounded py-2 p-2 text-white mt-2 hover:bg-blue-700"
                >
                  Explore
                </button>
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </div>
{/if}

<div class="relative">
  <button
    class="absolute bottom-2"
    on:click={() => app.nav.data_render_sheet_loader("source1", "group1")}
  >
    Sheets
  </button>
</div>
