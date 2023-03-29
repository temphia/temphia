<script lang="ts">
  import type { PageDashService, LoadResponse } from "../service";
  import Panel from "./panel.svelte";
  import Layout from "./_layout.svelte";
  export let service: PageDashService;
  export let data: LoadResponse;
</script>

<Layout name={data["name"]}>
  {#each data.sections || [] as section}
    <div class="flex flex-wrap justify-center md:justify-between gap-2 py-2">
      {#each section.panels || [] as panel}
        <Panel {panel} {data} {service} />
      {/each}
    </div>
  {/each}
</Layout>
