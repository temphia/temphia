<script lang="ts">
  import Element from "../../elements/element.svelte";
  import type { SubmitResponse } from "../../service";
  export let data: SubmitResponse;

</script>

<div class="container mx-auto py-2 p-2 h-full">
  <h1
    class="text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4"
  >
    {data.stage || ""}
  </h1>

  <div class="flex flex-col gap-2">
    {#each data["elements"] || [] as eValue}
      <Element data={data.data} element={eValue} />
    {/each}
  </div>
</div>



