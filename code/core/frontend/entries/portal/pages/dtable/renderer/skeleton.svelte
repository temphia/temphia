<div class="h-screen overflow-auto" style="width: calc(-3.5rem + 100vw);">
  <div class="w-full h-full overflow-x-hidden">
    <div class="m-1 pb-1 pl-1 pr-1 shadow bg-white rounded border">
      <div>
        <ul
          class="list-reset flex border-t-1 ml-4 divide divide-light-blue-400"
        >
          <li class="border-0 border-t-1 border-r-1 border-l-1 w-20" />
          <li class="nav-item p-2 border rounded-t inline-block">
            <div class="w-8 h-8 bg-gray-200 animate-pulse" />
          </li>
        </ul>
      </div>
      <div
        class="rounded-t-lg h-10 flex flex-col shadow md:flex-row justify-between items-center"
      >
        <div class="flex p-1 pr-4 gap-x-1">
          {#each Array(5) as _, i}
            <div
              class="py-1 px-3 w-20 h-8 text-gray-700 inline-flex font-light bg-gray-50 rounded hover:bg-gray-100 bg-gray-200 animate-pulse"
            />
          {/each}
        </div>
      </div>
      <div class="flex flex-col space-y-2 p-2">
        {#each Array(20) as _, i}
          <div class="grid grid-cols-5 gap-4">
            <div class="col-span-1 rouned h-10 bg-gray-200 animate-pulse" />
            <div class="col-span-4 rouned h-10 bg-gray-200 animate-pulse" />
          </div>
        {/each}
      </div>
    </div>
    <div class="fixed bottom-8 z-10 right-10 ">
      <div
        class="p-0 w-10 h-10 bg-blue-200 rounded-full active:shadow-lg mouse shadow transition ease-in duration-200 focus:outline-none"
      />
    </div>
  </div>
</div>
