<script lang="ts">
  import { LoadingSpinner } from "../../../xcompo";
  import Element from "../elements/element.svelte";
  import type { PageDashService, Panel, LoadResponse } from "../service";
  export let service: PageDashService;

  export let data: LoadResponse;
  export let panel: Panel;

  const view_opts = panel.view_opts || {};
  let loading = true;
  let panel_data = {};

  if (panel.source.startsWith("data/")) {
    console.log("@data", data, panel.source)
    

    panel_data = data[panel.source.replace("data/", "")] || {};

    console.log("@panel_data", panel_data)

    loading = false;
  }

  const classIt = (mode) => {
    switch (mode) {
      case "xm":
        return "w-28";
      case "sm":
        return "w-32";
      case "md":
        return "w-48";
      case "lg":
        return "w-64";
      case "xl":
        return "w-72";
      case "2xl":
        return "w-96";
      default:
        return "w-64";
    }
  };
</script>

<div class="rounded p-1 bg-white {classIt(view_opts["width"])}">
  {#if loading}
    <LoadingSpinner classes="" />
  {:else}
    <Element {data} {panel} />
  {/if}
</div>
