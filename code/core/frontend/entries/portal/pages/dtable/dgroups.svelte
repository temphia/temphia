<script lang="ts">
  import { strHash } from "../../../../lib/utils";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../app";

  const app: PortalApp = getContext("__app__");

  export let source;

  let groups = [];
  let sources = [];

  const load = async () => {
    const sresp = await app.get_dyn_sources();
    sources = sresp;

    const api = await app.get_apm().get_dyn_api();
    const resp = await api.list_group(source);
    groups = resp.data;
  };

  load();
</script>

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
          on:change={(ev) => {
            app.navigator.goto_dtable_source(ev.target["value"]);
          }}
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
          <div class="bg-white p-6 rounded-lg">
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
                on:click={() =>
                  app.navigator.goto_dtable_group(source, group["slug"])}
                class="text-center text-sm bg-blue-500 rounded py-2 p-2 text-white mt-2"
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
