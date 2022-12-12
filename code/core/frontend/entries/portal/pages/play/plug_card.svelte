<script>
  import Oicon from "./oicon.svelte";

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

<div class="bg-white shadow rounded-lg w-80 relative">
  <header class="p-6 border-b border-gray-200 relative">
    <div class="h-4 w-6 absolute -right-2 top-1/2 text-red-300 border rounded flex bg-white text-sm justify-between items-center">
      IN
      <div class="w-2 h-2 rounded-full bg-red-400 hover:bg-red-400" />
    </div>
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
      <div class="border">
        <header class="bg-gray-50 p-2 relative">
          <div
            class="h-2 w-2 rounded-full absolute -right-2 top-1/2 bg-blue-400 hover:bg-blue-400"
          />

          <h3 class="inline-flex text-lg font-semibold text-gray-700">
            <Oicon name="agent" />
            &nbsp;
            {agent.name}
          </h3>

          <span
            class="text-sm font-medium text-blue-50 rounded bg-blue-400  px-1"
            >{agent.slug}</span
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
              {#each agent.resources as resource}
                <div class="flex justify-around">
                  Id
                  <span class="px-1 bg-fuchsia-300 text-fuchsia-50 rounded"
                    >{resource.id}</span
                  >
                  Slot
                  <span class="px-1 bg-fuchsia-300 text-fuchsia-50 rounded"
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
              {#each agent.extensions as ext}
                <div class="flex justify-around">
                  Bprint
                  <span class="px-1 bg-fuchsia-300 text-fuchsia-50 rounded"
                    >{ext.bprint_id}
                  </span>
                  File
                  <span class="px-1 bg-fuchsia-300 text-fuchsia-50 rounded"
                    >{ext.file}</span
                  >
                </div>
              {/each}
            </div>
          </div>
        </div>
      </div>
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
              <span class="px-1 bg-fuchsia-300 text-fuchsia-50 rounded"
                >{resource}</span
              >
              Type
              <span class="px-1 bg-fuchsia-300 text-fuchsia-50 rounded"
                >{""}</span
              >
            </div>
          {/each}
        </div>
      </div>
    </footer>
  {/if}
</div>
