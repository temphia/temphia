<div class="w-full h-full bg-gray-100 overflow-auto text-gray-600 body-font">
  <div class="container px-5 py-24 mx-auto max-w-7x1">
    <div class="flex justify-between w-full">
      <div class="w-80 h-4 mb-6 lg:mb-0 bg-gray-200 animate-pulse mb-4 p-4" />
      <div class="w-20 h-4 bg-gray-200 animate-pulse mb-4 p-4" />
    </div>
    <div class="flex flex-wrap -m-4">
      {#each Array(10) as _}
        <div class="xl:w-1/3 md:w-1/2 p-4">
          <div class="p-6 rounded-lg bg-gray-200 animate-pulse">
            <div
              class="lg:h-60 xl:h-56 md:h-64 sm:h-72 xs:h-72 h-72 bg-gray-400 animate-pulse  rounded w-full"
            />
            <div class="m-5 h-5 w-20 bg-gray-300 animate-pulse" />
            <div class="m-5 h-5 w-40 bg-gray-300 animate-pulse" />
            <div class="m-5 h-5 w-100 bg-gray-300 animate-pulse" />
            <div class="flex p-5">
              <div class="rounded h-2 w-4 bg-gray-300 animate-pulse" />
            </div>
          </div>
        </div>
      {/each}
    </div>
  </div>
</div>
