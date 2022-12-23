<script lang="ts">
  import type { PortalService } from "../../core";

  export let app: PortalService;
  export let bid;

  app.registry.GetAll("bprint_editor_type");

  const fromRepo = () => {};
</script>

<div class="flex items-center justify-between">
  <h4 class="font-medium text-slate-500">Bprint Editors.</h4>
</div>
<div class="space-y-2 mt-4">
  <div
    on:click={fromRepo}
    class="flex space-x-4 rounded-xl bg-white p-3 shadow-sm hover:border border-blue-500 cursor-pointer"
  >
    <div>
      <h4 class="font-semibold text-gray-600">From Repo</h4>
      <p class="text-sm text-slate-400">
        You can import blueprint from some repo and instance it.
      </p>
    </div>
  </div>
</div>
