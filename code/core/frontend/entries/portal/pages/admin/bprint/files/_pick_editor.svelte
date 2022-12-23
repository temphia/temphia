<script lang="ts">
  import type { PortalService } from "../../core";

  export let app: PortalService;
  export let bid;
  export let file;

  // app.registry.GetAll("bprint_editor_type");

  const BprintEditors = [{ name: "Data Schema" }, { name: "Data Migrator" }];

  const fromRepo = () => {};
</script>

<div class="flex items-center justify-between">
  <h4 class="text-slate-800 font-semibold text-xl p-1 mb-2">Bprint Editors</h4>
</div>
<div class="flex flex-col gap-2">
  {#each BprintEditors as be}
    <div
      on:click={fromRepo}
      class="flex justify-between rounded-xl bg-white p-3 shadow-sm border "
    >
      <h4 class="font-semibold text-gray-600">{be.name}</h4>
      <button class="p-1 bg-blue-400 hover:bg-blue-600 rounded text-white">Open</button>
    </div>
  {/each}
</div>
