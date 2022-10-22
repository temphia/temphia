<script lang="ts">
  import { getContext } from "svelte";

  const plugs = [
    {
      plug_name: "Example",
      plug_id: "zzzz",
      agent_id: "yyyyy",
      icon: "",
      ticket: "",
    },

    {
      plug_name: "Notes",
      plug_id: "zzzz",
      agent_id: "yyyyy",
      icon: "",
      ticket: "",
    },

    {
      plug_name: "Example 2",
      plug_id: "zzzz",
      agent_id: "yyyyy",
      icon: "",
      ticket: "",
    },
  ];
</script>

<div
  class="w-full h-full py-10 mx-auto overflow-auto bg-gradient-to-b from-purple-100 to-indigo-100"
>
  <div class="flex justify-center w-full mb-10">
    <div class="relative">
      <input
        type="text"
        class="h-14 w-96 pr-8 pl-5 rounded z-0 focus:shadow focus:outline-none"
        placeholder="Search anything..."
      />
      <div class="absolute top-4 right-3 flex">
        <svg class="h-5 w-5 text-gray-500" viewBox="0 0 24 24" fill="none"
          ><path
            d="M21 21L15 15M17 10C17 13.866 13.866 17 10 17C6.13401 17 3 13.866 3 10C3 6.13401 6.13401 3 10 3C13.866 3 17 6.13401 17 10Z"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          /></svg
        >
      </div>
    </div>
  </div>

  <div class="p-8 flex flex-wrap justify-center gap-4">
    {#each plugs as plug}
      <div
        class="bg-white flex flex-col items-center h-32 w-32 p-2 overflow-hidden shadow-lg rounded-lg cursor-pointer hover:border-2 border-blue-400"
      >
        {#if plug.icon === ""}
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-28 h-28 text-gray-500"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M11 4a2 2 0 114 0v1a1 1 0 001 1h3a1 1 0 011 1v3a1 1 0 01-1 1h-1a2 2 0 100 4h1a1 1 0 011 1v3a1 1 0 01-1 1h-3a1 1 0 01-1-1v-1a2 2 0 10-4 0v1a1 1 0 01-1 1H7a1 1 0 01-1-1v-3a1 1 0 00-1-1H4a2 2 0 110-4h1a1 1 0 001-1V7a1 1 0 011-1h3a1 1 0 001-1V4z"
            />
          </svg>
        {:else}
          <div class="w-28 h-28">
            {@html plug.icon}
            <!-- fixme => escape contents -->
          </div>
        {/if}

        <h3
          class="tracking-widest text-indigo-500 text-sm uppercase font-medium title-font"
        >
          {plug.plug_id}#{plug.agent_id}
        </h3>

        <h2 class="text-lg text-gray-500 font-semibold font-mono">
          {plug.plug_name}
        </h2>
      </div>
    {/each}
  </div>
</div>
