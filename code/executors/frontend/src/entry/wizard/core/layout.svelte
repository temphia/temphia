<script>
  export let title = "";
  export let prev = () => {};
  export let next = () => {};
  export let showButtons = true;
</script>

<div class="h-full bg-gray-100">
  <div class="flex flex-col w-full h-full">
    <header
      class="flex-shrink flex items-center justify-between h-14 bg-white shadow border-b"
    >
      <div class="px-4">
        <div class="text-xl font-bold text-slate-600">{title}</div>
      </div>

      <div class="px-4">
        <svg
          class="h-6 w-6 cursor-pointer"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"
          />
        </svg>
      </div>
    </header>
    <div class="flex-grow h-32 overflow-auto">
      <div
        class="flex flex-col w-full bg-white rounded-lg shadow p-10 px-8 pt-6 mx-auto my-6 mb-4 transition duration-500 ease-in-out transform sm:w-3/4 lg:w-1/2 3xl:w-1/3 "
      >
        <slot />
      </div>
    </div>

    {#if showButtons}
      <footer
        class="flex-shrink flex justify-between items-center border-t h-12 bg-gray-50"
      >
        <div class="px-2">
          <button
            on:click={prev}
            class="py-2 px-4 border rounded-md cursor-pointer uppercase text-sm font-bold text-blue-500  hover:bg-blue-200"
            >Previous</button
          >
        </div>
        <div class="px-2">
          <button
            on:click={next}
            class="py-2 px-4 border rounded-md cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700"
            >Next</button
          >
        </div>
      </footer>
    {/if}
  </div>
</div>
