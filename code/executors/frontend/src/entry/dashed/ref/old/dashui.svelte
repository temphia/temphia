<script>
  import DashRenderer from "./dash_renderer.svelte";
</script>

<div class="h-full w-full flex-1 flex flex-col overflow-hidden">
  <main class="flex-1 overflow-x-hidden overflow-y-auto bg-gray-200">
    <div class="container mx-auto p-6">
      <h3 class="text-gray-700 text-3xl font-medium">Home</h3>

      <DashRenderer />
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
