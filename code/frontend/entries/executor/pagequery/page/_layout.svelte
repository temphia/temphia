<script lang="ts">
  export let title = "PageQuery";
</script>

<div class="h-full w-full flex-1 flex flex-col overflow-hidden">
  <div class="flex p-2 justify-between bg-blue-300">
    <div class="flex"><h3 class="text-white mr-4">{title}</h3></div>
    <div class="flex gap-2">
      <button
        class="bg-blue-200 hover:bg-blue-500 rounded font-bold text-white text-sm px-2"
        >↻</button
      >
    </div>
  </div>
  <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
    <div class="container mx-auto py-2">
      <slot />
    </div>
  </main>
</div>


