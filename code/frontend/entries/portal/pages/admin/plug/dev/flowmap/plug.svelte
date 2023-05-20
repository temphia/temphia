<script lang="ts">
  import type { FormatedPlug } from "./formatter";

  export let data: FormatedPlug;

  const plug = data.plug;
</script>

<div
  class="bg-white p-4 w-full overflow-auto"
  style="height: calc(100% - 1rem)"
>


  <div class="flex justify-center">
    <div class="inline-flex items-center p-1">
      <h2 class="line-clamp-1 text-lg font-light uppercase text-gray-800">
        {plug.name}
      </h2>
      <span class="ml-1 rounded bg-gray-100 p-0.5 text-xs text-gray-800"
        >{plug.id}</span
      >
    </div>
  </div>

  <div class="flex flex-col gap-2">
    {#each data.agents as agent}
      <div class="relative flex h-32 justify-center rounded-lg">
        <div
          class="w-full transform transition-transform duration-500 ease-in-out hover:scale-110"
        >
          <div
            class="absolute inset-0 flex justify-center bg-yellow-400 opacity-50"
          >
            <div class="inline-flex items-start">
              <h3 class="uppercase">{agent.name || ""}</h3>
              <span
                class="ml-1 mt-1 rounded bg-gray-100 p-0.5 text-xs text-gray-800"
                >{agent.id}</span
              >
            </div>
          </div>
        </div>
        <!-- EXTENSIONS -->
        <!-- <div class="absolute -left-2 top-5 flex flex-col items-start space-y-2">
          {#each extensions[agent.id] || [] as extkey}
            <div class="flex space-x-5">
              <button
                class="flex items-center rounded-lg bg-white px-2 py-1 font-medium text-gray-600 shadow hover:bg-gray-300"
                >{extkey}</button
              >
            </div>
          {/each}
        </div> -->

        <!-- RESOURCES -->

         <div class="absolute -right-2 top-5 flex flex-col items-end space-y-2">
          {#each Object.keys(data.resources[agent.id] || [])  || [] as [rkey, rdata]}
            <div class="flex space-x-5">
              <button
                class="flex items-center rounded-lg bg-white px-2 py-1 font-medium text-gray-600 shadow hover:bg-gray-300"
                >{rkey}</button
              >
            </div>
          {/each}
        </div> 

        <div
          id="agent-in-port-{agent.id}"
          class="h-4 w-10 absolute -right-2 bottom-1 text-red-300 border rounded flex bg-white text-sm justify-between items-center"
        >
          IN
        </div>

        <div
          id="agent-out-port-{agent.id}"
          class="h-4 w-10 absolute -left-2 bottom-1 text-red-300 border rounded flex bg-white text-sm justify-between items-center"
        >
          OUT
        </div>

        <div class="absolute -bottom-1 right-1/2">
          <div
            id="bottom-plug-${plug.id}"
            class="h-2 w-2 rounded-full bg-red-400 hover:scale-150 hover:bg-red-600"
          />
        </div>
      </div>
    {/each}
  </div>
</div>
