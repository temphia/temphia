<script lang="ts">
  import type { PortalService } from "../../core";

  export let app: PortalService;
  export let bid: string;
  export let file: string;

  const etypes = Array.from(app.xtmgr.bprint_editors.keys());
</script>

<div class="flex items-center justify-between">
  <h4 class="text-slate-800 font-semibold text-xl p-1 mb-2">Bprint Editors</h4>
</div>
<div class="flex flex-col gap-2">
  {#each etypes as be}
    <div class="flex justify-between rounded-xl bg-white p-3 shadow-sm border ">
      <h4 class="font-semibold text-gray-600">{be}</h4>
      <button
        class="p-1 bg-blue-400 hover:bg-blue-600 rounded text-white"
        on:click={() => {
          app.utils.small_modal_close();
          app.nav.admin_bprint_editor(bid, be, file);
        }}>Open</button
      >
    </div>
  {/each}
</div>
