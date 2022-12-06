<script lang="ts">
  import type { PortalService } from "../../core";

  export let app: PortalService;

  const fromRepo = () => {};
  const fromBuilder = () => {};
</script>

<div class="flex items-center justify-between">
  <h4 class="font-medium text-slate-500">Ways to create new data table.</h4>
</div>
<div class="space-y-2 mt-4">
  <div
    on:click={fromRepo}
    class="flex space-x-4 rounded-xl bg-white p-3 shadow-sm hover:border border-blue-500 cursor-pointer"
  >
    <img
      class="aspect-square w-16 rounded-lg bg-center object-cover"
      src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTXeOqFwpuGs8POfrytccjgsihqbnERH0hOXg&amp;usqp=CAU"
      alt=""
    />
    <div>
      <h4 class="font-semibold text-gray-600">From Repo</h4>
      <p class="text-sm text-slate-400">
        You can import blueprint from some repo and instance it.
      </p>
    </div>
  </div>

  <div
    on:click={fromBuilder}
    class="flex space-x-4 rounded-xl bg-white p-3 shadow-sm hover:border border-blue-500 cursor-pointer"
  >
    <img
      class="aspect-square w-16 rounded-lg bg-center object-cover"
      src="https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcTXeOqFwpuGs8POfrytccjgsihqbnERH0hOXg&amp;usqp=CAU"
      alt=""
    />
    <div>
      <h4 class="font-semibold text-gray-600">From Builder</h4>
      <p class="text-sm text-slate-400">
        User Builder to create data table (schema ) from scratch.
      </p>
    </div>
  </div>
</div>
