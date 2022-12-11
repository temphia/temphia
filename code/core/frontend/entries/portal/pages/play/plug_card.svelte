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

<div class="bg-white shadow rounded-lg w-64">
  <header class="p-6 border-b border-gray-200">
    <h3 class="text-lg font-semibold text-gray-700">{data.root_plug.name}</h3>
    <h4 class="text-sm font-medium text-gray-700">
      Slug: {data.root_plug.slug}
    </h4>
  </header>

  <body class="flex flex-col p-2 ">
    {#each data.root_plug.agents as agent}
      <div class="border">
        <header class="bg-gray-50">
          <h3 class="text-lg font-semibold text-gray-700 inline-flex">
            <Oicon name="agent" />
            {agent.name}
          </h3>
          <h4 class="text-sm font-medium text-gray-700">
            Slug: {agent.slug}
          </h4>
        </header>

        <body class="flex flex-col border p-2">
          <div class="flex flex-col">
            <h4>Links</h4>
            <ul class="list-disc pl-4">
              {#each agent.links as link}
                <li>{link.to}</li>
              {/each}
            </ul>
          </div>

          <div class="flex flex-col">
            <h4>Resources</h4>
            <ul class="list-disc pl-4">
              {#each agent.resources as resource}
                <li>{resource.id} |> {resource.slot}</li>
              {/each}
            </ul>
          </div>

          <div class="flex flex-col">
            <h4>Extensions</h4>
            <ul class="list-disc pl-4">
              {#each agent.extensions as extension}
                <li>{extension.id} |> {extension.slot}</li>
              {/each}
            </ul>
          </div>
        </body>
      </div>
    {/each}
  </body>

  <h4 class="text-sm font-medium text-gray-700">Resources:</h4>
  <ul class="list-disc pl-4">
    {#each data.root_plug.resources as resource}
      <li>{resource}</li>
    {/each}
  </ul>
</div>



<!-- <div class="h-scree flex w-screen items-center self-center p-10">
    <div class="w-64 rounded-lg bg-white shadow">
      <header class="border-b border-gray-200 p-2">
        <h3 class="text-lg font-semibold text-gray-700">Plug One</h3>
        <div> <span class="text-sm font-medium text-gray-700 rounded bg-slate-300 px-1">p1</span></div>
        
      </header>
      <div class="flex flex-col p-2">
        <div class="border">
          <header class="bg-gray-50 p-2">
            <h3 class="inline-flex text-lg font-semibold text-gray-700">
              <svg xmlns="http://www.w3.org/2000/svg" aria-hidden="true" viewBox="0 0 24 24" stroke="currentColor" fill="none" name="users" class="h-5 w-5"><path stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M12 4.354a4 4 0 110 5.292M15 21H3v-1a6 6 0 0112 0v1zm0 0h6v-1a6 6 0 00-9-5.197M13 7a4 4 0 11-8 0 4 4 0 018 0z"></path></svg> &nbsp;default
            </h3>
            <h4 class="text-sm font-medium text-gray-700">
              <span class="px-1 bg-gray-200 rounded">default</span>
            </h4>
          </header>
  
          <div class="flex flex-col p-2 gap-2 rounded">
            
            <div class="flex flex-col border p-1">
              <h4>Resources</h4>
  
              <div class="pl-4">
                <span>xyz |&gt; pqr</span>
              </div>
  
            </div>
            <div class="flex flex-col border p-1">
              <h4>Extensions</h4>
              <ul class="list-disc pl-4">
                <li>undefined |&gt; undefined</li>
              </ul>
            </div>
          </div>
        </div>
      </div>
  
          <div class="flex flex-col p-2">
                <h4 class="text-sm font-medium text-gray-700">Resources:</h4>
          </div>
  
  
  
  
    </div> 
  </div> -->