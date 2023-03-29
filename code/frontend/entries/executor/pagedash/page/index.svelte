<script lang="ts">
  import type { PageDashService, LoadResponse } from "../service";
  import Panel from "./panel.svelte";
  import Layout from "./_layout.svelte";
  export let service: PageDashService;
  export let data: LoadResponse;

  $: console.log("@data", data);
</script>

<Layout name={data["name"]}>
  {#each data.sections || [] as section}
    <span class="m-2 p-1 inline-flex rounded bg-gray-100"
      ><h2 class="text-gray-600 ">{section.name || ""}</h2>
      <svg
        viewBox="0 0 24 24"
        stroke="none"
        fill="currentColor"
        name="chevron-down"
        class="w-6"
        ><path
          fill-rule="evenodd"
          d="M5.293 7.293a1 1 0 011.414 0L10 10.586l3.293-3.293a1 1 0 111.414 1.414l-4 4a1 1 0 01-1.414 0l-4-4a1 1 0 010-1.414z"
          clip-rule="evenodd"
        /></svg
      ></span
    >

    <div
      class="flex flex-wrap justify-center md:justify-between gap-2 py-2 rounded border"
    >
      {#each section.panels || [] as panel}
        <Panel {panel} {data} {service} />
      {/each}
    </div>
  {/each}
</Layout>
