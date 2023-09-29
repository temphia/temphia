<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { strHash } from "$lib/utils";
  import Dropdown from "$lib/compo/autotable/_dropdown.svelte";
  import type { PortalService } from "$lib/core";
  import { FloatingAdd, LoadingSpinner } from "$lib/core";
  import NewDataPicker from "./_new_data_picker.svelte";

  export let source = "default";

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

  const do_explore = (group: object) => {
    console.log("@group =>", group);
    const slug = group["slug"];

    switch (group["renderer"]) {
      case "sheet":
        app.nav.data_sheets_page(source, slug);
        break;
      default:
        app.nav.data_group_page(source, slug);
        break;
    }
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="w-full h-full overflow-auto">
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
            <div class="card p-6">
              <img
                class="lg:h-60 xl:h-56 md:h-64 sm:h-72 xs:h-72 h-72 rounded w-full object-cover object-center mb-6"
                src="https://picsum.photos/seed/{strHash(
                  group.name + group.slug
                )}d/800/400"
                alt=""
              />

              <div class="flex gap-1">
                <h3
                  class="tracking-widest text-indigo-500 text-xs uppercase font-medium title-font h3"
                >
                  {source}
                </h3>

                {#if group["renderer"]}
                  <h3
                    class="tracking-widest text-blue-500 text-xs uppercase font-medium title-font bg-yellow-400 rounded"
                  >
                    #{group["renderer"] || ""}
                  </h3>
                {/if}
              </div>

              <h2 class="text-lg text-gray-900 font-medium title-font mb-4 h2">
                {group.name}
              </h2>
              <p class="leading-relaxed selection:bg-red-200 text-base">
                {group.description}
              </p>
              <div class="flex p-5 gap-2">
                <button
                  on:click={() => do_explore(group)}
                  class="btn btn-md variant-filled-primary"
                >
                <Icon name="view-list" class="h-5 w-5" />
                  Explore
                </button>

                <button
                  on:click={() =>
                    app.nav.data_group_page(source, group["slug"])}
                  class="btn btn-sm variant-filled-secondary"
                >
                  <Icon name="hashtag" class="h-5 w-5" />
                  <span>Raw</span>
                </button>

                <button
                  on:click={() =>
                    app.nav.admin_data_group(source, group["slug"])}
                  class="btn btn-sm variant-filled-secondary"
                >
                  <Icon name="cog" class="h-5 w-5" />
                  <span>Setting</span>
                </button>
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </div>
{/if}

<FloatingAdd
  onClick={() => {
    app.utils.small_modal_open(NewDataPicker, { app });
  }}
/>
