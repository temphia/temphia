<script>
  const templates = [1, 2, 3, 4, 5, 6, 7, 8];
</script>

<div class="-m-4 flex flex-wrap p-2 md:p-0 gap-2 overflow-auto justify-center">
  {#each templates as tpl}
    <div class="w-full p-4 md:w-1/2 lg:w-1/4 shadow rounded">
      <div class="block overflow-hidden rounded">
        <img
          alt=""
          class="block w-full h-auto object-cover object-center cursor-pointer"
          src="https://dummyimage.com/400x200"
        />
      </div>
      <div class="mt-4">
        <h2 class="title-font text-lg font-medium text-gray-900">
          The Catalyzer
        </h2>
      </div>
    </div>
  {/each}
</div>
