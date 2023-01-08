<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../core";
  import Layout from "../_layout.svelte";
  import Pick from "./_pick.svelte";

  let message = "";

  let slug = "";

  let source = "";
  let folder = "";

  const pick = () => {
    app.utils.small_modal_open(Pick, {});

    source = "default";
    folder = "example";
  };

  const app = getContext("__app__") as PortalService;
</script>

<Layout>
  <p class="text-red-500">{message}</p>

  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="slug">
      Id
    </label>
    <input
      class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
      id="slug"
      type="text"
      placeholder="Slug"
    />

    <p class="italic text-sm mt-2">
      Id should only contain small letters and hypthens.
    </p>
  </div>
  <div class="mb-4">
    <label class="block mb-2 text-sm font-bold text-gray-700" for="name">
      Value
    </label>

    <div class="flex justify-between">
      <ul class="flex flex-wrap space-x-3">
        <li class="flex items-center space-x-2">
          <svg
            xmlns="http://www.w3.org/2000/svg"
            class="w-4 h-4"
            fill="none"
            viewBox="0 0 24 24"
            stroke="currentColor"
            stroke-width="2"
          >
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
            />
          </svg>
          <a href="#" class="text-gray-800">Home</a>
        </li>
        <li>/</li>
        <li class="flex items-center space-x-2 cursor-pointer" on:click={pick}>
          <svg
            class="w-5 h-5"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="currentColor"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <path d="M14 3l6 4v7" />
            <path d="M18 17l-6 4l-6 -4" />
            <path d="M4 14v-7l6 -4" />
            <path d="M4 7l8 5v9" />
            <path d="M20 7l-8 5" />
          </svg>

          <span class="text-gray-400"> Source [{source || "<empty>"}]</span>
        </li>

        <li>/</li>
        <li class="flex items-center space-x-2 cursor-pointer" on:click={pick}>
          <svg
            class="w-5 h-5"
            viewBox="0 0 24 24"
            stroke-width="2"
            stroke="currentColor"
            fill="none"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path stroke="none" d="M0 0h24v24H0z" fill="none" />
            <path
              d="M5 4h4l3 3h7a2 2 0 0 1 2 2v8a2 2 0 0 1 -2 2h-14a2 2 0 0 1 -2 -2v-11a2 2 0 0 1 2 -2"
            />
          </svg>
          <span class="text-gray-400">Folder [{folder || "<empty>"}]</span>
        </li>
      </ul>
    </div>
  </div>

  <div class="flex justify-end">
    <button
      class="p-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline"
    >
      Create
    </button>
  </div>
</Layout>
