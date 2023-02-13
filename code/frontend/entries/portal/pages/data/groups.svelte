<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { strHash } from "../../../../lib/utils";
  import Dropdown from "../../../xcompo/autotable/_dropdown.svelte";
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

  const do_explore = (group: object) => {
    console.log("@group =>", group);
    const slug = group["slug"];

    switch (group["renderer"]) {
      case "sheet":
        app.nav.data_render_sheet_loader(source, slug);
        break;
      default:
        app.nav.data_render_table_loader(source, slug);
        break;
    }
  };
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

              <div class="flex gap-1">
                <h3
                  class="tracking-widest text-indigo-500 text-xs uppercase font-medium title-font"
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

              <h2 class="text-lg text-gray-900 font-medium title-font mb-4">
                {group.name}
              </h2>
              <p class="leading-relaxed selection:bg-red-200 text-base">
                {group.description}
              </p>
              <div class="flex p-5 gap-2">
                <button
                  on:click={() => do_explore(group)}
                  class="text-center text-sm bg-blue-500 rounded py-2 p-2 text-white hover:bg-blue-700"
                >
                  Explore
                </button>

                <Dropdown>
                  <button
                    on:click={() =>
                      app.nav.data_render_table_loader(source, group["slug"])}
                    class="flex justify-between rounded-sm px-4 py-2 text-sm capitalize text-gray-700 hover:bg-blue-500 hover:text-white"
                  >
                    <Icon name="hashtag" class="h-5 w-5" />
                    <span>Raw</span>
                  </button>

                  <button
                    on:click={() =>
                      app.nav.admin_data_group(source, group["slug"])}
                    class="flex justify-between rounded-sm px-4 py-2 text-sm capitalize text-gray-700 hover:bg-blue-500 hover:text-white"
                  >
                    <Icon name="cog" class="h-5 w-5" />
                    <span>Setting</span>
                  </button>
                </Dropdown>
              </div>
            </div>
          </div>
        {/each}
      </div>
    </div>
  </div>
{/if}
