<script lang="ts">
  import Oicon from "./oicon.svelte";
  import Agent from "./_agent.svelte";
  import Port from "./_port.svelte";

  const data = {
    root_plug: {
      name: "Plug One",
      slug: "p1",
      agents: [
        {
          name: "default",
          slug: "default",
          links: [{ to: "aa" }],
          resources: [{ id: "xyz", slot: "pqr" }],
          extensions: [{ bprint_id: "mno", file: "aa.js" }],
        },
      ],
      resources: [],
    },

    other_plugs: {
      p2: {
        name: "Plug Two",
        slug: "p2",
      },
    },
  };
</script>

<div class="bg-white shadow rounded-lg w-80">
  <header class="p-6 border-b border-gray-200">
    <h3 class="text-lg font-semibold text-gray-700 inline-flex">
      <Oicon name="plug" />
      &nbsp;
      {data.root_plug.name}
    </h3>

    <div>
      <span class="text-sm font-medium text-red-50 rounded bg-red-400 px-1"
        >{data.root_plug.slug}</span
      >
    </div>
  </header>

  <div class="flex flex-col p-2 ">
    {#each data.root_plug.agents as agent}
      <Agent data={agent} />
    {/each}
  </div>

  {#if data.root_plug.resources}
    <footer class="p-2 border-t border-gray-200">
      <div class="border">
        <header class="bg-gray-50 p-2 border-b">
          <h3 class="inline-flex text-lg font-semibold text-gray-700">
            <Oicon name="resource" />
            &nbsp; Resources
          </h3>
        </header>

        <div class="flex flex-col">
          {#each data.root_plug.resources as resource}
            <div class="flex justify-around">
              Id
              <span class="px-1 bg-green-300 text-green-50 rounded"
                >{resource}</span
              >
              Type
              <span class="px-1 bg-green-300 text-green-50 rounded"
                >{""}</span
              >
            </div>
          {/each}
        </div>
      </div>
    </footer>
  {/if}


</div>
