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

  <div class="p-8 flex flex-wrap justify-center gap-4">
    {#each Array(20) as _, i}
      <div
        class="bg-white flex flex-col items-center h-30 w-40 p-4 overflow-hidden"
      >
        <div class="w-28 h-28 bg-gray-200 animate-pulse" />
        <div class="w-20 h-6 bg-gray-200 m-1 rounded animate-pulse" />
      </div>
    {/each}
  </div>
</div>
