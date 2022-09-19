<script lang="ts">
  import Node from "./_node.svelte";

  const data = {
    cluster_name: "mnop",
    nodes: {
      pqr: {
        id: "pqr",
        tags: ["xyz", "abc"],
        stats: [
          {
            epoch: 1,
            total_mem: 100,
            used_mem: 10,
            total_swap: 20,
            used_swap: 12,
            cpu: 40,
            avg_load: 10,
          },
        ],
      },
    },
  };

  const process = (d: typeof data) => {
    Object.values(d.nodes).forEach((node) => {
      node.stats = node.stats.sort((a, b) => a.epoch - b.epoch);
    });
  };
</script>

<div class="h-full w-full flex-1 flex flex-col overflow-hidden">
  <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
    <div class="container mx-auto px-6 py-8">
      <h3 class="text-gray-700 text-3xl font-medium">Home</h3>

      <div class="flex flex-col mt-8 space-y-4">
        <div class="w-full bg-white rounded p-2">
          <table class="w-full border">
            <thead>
              <tr
                class="bg-gray-200 text-gray-600 uppercase text-sm leading-normal"
              >
                <th class="py-3 px-6 text-left">Properties</th>
                <th class="py-3 px-6 text-left">Value</th>
              </tr>
            </thead>
            <tbody class="text-gray-600 text-sm font-light">
              <tr class="border-b border-gray-200 hover:bg-gray-100">
                <th class="py-3 px-6 text-left">Cluster Name</th>
                <th class="py-3 px-6 text-left">{data["cluster_name"]}</th>
              </tr>

              <tr class="border-b border-gray-200 hover:bg-gray-100">
                <td class="py-3 px-6 text-left">Core Database</td>
                <td class="py-3 px-6 text-left">
                  <div>
                    Vendor: Postgres
                  </div>
                  <div>
                    Name: CoreDB1
                  </div>
                </td>
              </tr>
              <tr class="border-b border-gray-200 hover:bg-gray-100">
                <td class="py-3 px-6 text-left">Nodes</td>
                <td class="py-3 px-6 text-left">
                  {#each Object.entries(data.nodes) as [id, node]}
                    <Node />
                  {/each}
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </main>
</div>
