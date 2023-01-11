<div
  class="w-full h-full py-10 mx-auto overflow-auto bg-gradient-to-b from-purple-100 to-indigo-100"
>
  <div class="flex justify-center w-full mb-10">
    <div class="relative">
      <div
        type="text"
        class="h-14 w-96 pr-8 pl-5 rounded z-0 focus:shadow focus:outline-none bg-white animate-pulse"
      />
      <div class="absolute top-4 right-3 flex">
        <div class="h-6 w-6 bg-gray-200 animate-pulse" />
      </div>
    </div>
  </div>

  <div
    class="p-2 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 2xl:grid-cols-5 gap-2"
  >
    {#each Array(10) as _, i}
      <div class="bg-white p-4 overflow-hidden rounded-lg">
        <div class="w-58 h-44 bg-gray-200 animate-pulse" />
        <div class="mt-8 h-32 w-full space-y-3">
          <div class="w-20 h-6 bg-gray-200 rounded-full animate-pulse" />
          <div class="w-full h-4 bg-gray-200 rounded-full animate-pulse" />
          <div class="w-full h-4 bg-gray-200 rounded-full animate-pulse" />
          <div class="w-1/2 h-4 bg-gray-200 rounded-full animate-pulse" />
        </div>
      </div>
    {/each}
  </div>
</div>
