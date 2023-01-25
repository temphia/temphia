<script>
  export let title;
  export let onClick;
</script>

<div class="flex flex-col h-full">
  <div class="mb-4 border-b">
    <h1 class="block mb-2 text-lg font-bold text-gray-700" for="slug">
      {title}
    </h1>
  </div>

  <div class="flex flex-col flex-grow">
    <slot />
  </div>

  <div class="flex justify-end border-t">
    <button
      on:click={onClick}
      class="px-2 py-1 font-bold text-white bg-blue-500 hover:bg-blue-700 focus:outline-none focus:shadow-outline rounded mt-1"
      >Save</button
    >
  </div>
</div>
