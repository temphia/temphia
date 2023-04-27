<script lang="ts">
  import Element from "../../elements/element.svelte";
  import type { SubmitResponse } from "../../service";
  export let data: SubmitResponse;

  // const data1: SubmitResponse = {
  //   data: {},
  //   elements: [
  //     {
  //       data_opts: {},
  //       info: "",
  //       name: "",
  //       source: "",
  //       type: "",
  //       view_opts: {},
  //     },
  //   ],
  //   stage: "",
  // };

  // data = data1;
</script>

<div class="container mx-auto py-2 rounded bg-white p-2 h-full">
  <h1
    class="text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4"
  >
    {data.stage || ""}
  </h1>

  <div class="flex flex-col">
    {#each data["elements"] || [] as eValue}
      <Element data={data.data} element={eValue} />
    {/each}
  </div>
</div>



