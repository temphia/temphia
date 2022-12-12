<script lang="ts">
  import Oicon from "./oicon.svelte";
  import Port from "./_port.svelte";
  export let data;
</script>

<div class="border">
  <header class="bg-gray-50 p-2 pb-3 relative">
    <Port color="blue" name="IN" left={true} />
    <Port color="red" name="OUT" left={false} />

    <h3 class="inline-flex text-lg font-semibold text-gray-700">
      <Oicon name="agent" />
      &nbsp;
      {data.name}
    </h3>

    <span class="text-sm font-medium text-blue-50 rounded bg-blue-400  px-1"
      >{data.slug}</span
    >
  </header>

  <div class="flex flex-col border p-2">
    <div class="flex flex-col p-1 border">
      <header class="bg-gray-50 p-1">
        <h3 class="inline-flex text-gray-700 uppercase">
          <Oicon name="resource" />
          &nbsp; Resources
        </h3>
      </header>

      <div class="flex flex-col">
        {#each data.resources as resource}
          <div class="flex justify-around">
            Id
            <span class="px-1 bg-green-300 text-green-50 rounded"
              >{resource.id}</span
            >
            Slot
            <span class="px-1 bg-green-300 text-green-50 rounded"
              >{resource.slot}</span
            >
          </div>
        {/each}
      </div>
    </div>

    <div class="flex flex-col p-1 border">
      <header class="bg-gray-50 p-1">
        <h3 class="inline-flex text-gray-700 uppercase">
          <Oicon name="ext" />
          &nbsp; Extensions
        </h3>
      </header>

      <div class="flex flex-col">
        {#each data.extensions as ext}
          <div class="flex justify-around">
            Bprint
            <span class="px-1 bg-green-300 text-green-50 rounded"
              >{ext.bprint_id}
            </span>
            File
            <span class="px-1 bg-green-300 text-green-50 rounded"
              >{ext.file}</span
            >
          </div>
        {/each}
      </div>
    </div>
  </div>
</div>
