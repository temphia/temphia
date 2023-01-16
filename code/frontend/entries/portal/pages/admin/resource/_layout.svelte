<script>
  export let onClick;
</script>

<div class="p-4">
  <div class="bg-white rounded p-4 flex flex-col">
    <slot />

    <div class="flex justify-end">
      <button
        on:click={onClick}
        class="p-2 font-bold text-white bg-blue-500 rounded hover:bg-blue-700 focus:outline-none focus:shadow-outline"
      >
        Create
      </button>
    </div>
  </div>
</div>
