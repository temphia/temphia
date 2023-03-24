<script lang="ts">
  import Element from "./elements/element.svelte";
  import type { PageDashService } from "./service";
  export let service: PageDashService;
  export let data: any;

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

<div class="h-full w-full flex-1 flex flex-col overflow-hidden">
  <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
    <div class="container mx-auto py-2">
      <h3 class="text-gray-700 text-3xl font-medium">{data["name"] || ""}</h3>

      {#each data.sections || [] as section}
        <div
          class="flex flex-wrap justify-center md:justify-between gap-2 py-2"
        >
          {#each section.panels || [] as panel}
            <div class={classIt(panel.class)}>
              <Element data={{}} {panel} />
            </div>
          {/each}
        </div>
      {/each}
    </div>
  </main>
</div>

<div class="hidden">
  <div class="w-28">xm</div>
  <div class="w-32">sm</div>
  <div class="w-48">md</div>
  <div class="w-64">lg</div>
  <div class="w-72">xl</div>
  <div class="w-96">2xl</div>
</div>
