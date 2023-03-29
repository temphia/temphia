<script lang="ts">
  export let name: string = "";
</script>

<div class="h-full w-full flex-1 flex flex-col overflow-hidden">
  <div class="flex p-2 justify-between bg-blue-300">
    <div class="flex"><h3 class="text-white mr-4">Dash [{name}]</h3></div>
    <div class="flex gap-2" />
  </div>
  <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
    <div class="container mx-auto py-2">
      <slot />
    </div>
  </main>
</div>

<div class="hidden">
  <div class="w-28">xm</div>
  <div class="w-32">sm</div>
  <div class="w-48">md</div>
  <div class="w-64">lg</div>
  <div class="w-72">xl</div>
  <div class="w-96">2xl</div>
</div>
