
<script lang="ts">
    import { LoadingSpinner, PortalService } from "$lib/core";
  
    export let id;
    export let app: PortalService;
  
    let loading = false;
  
    let mode: "SELECT" | "RESULT" = "SELECT";
  
    const load = async () => {
      loading = true;
      const papi = app.api_manager.get_admin_plug_api();
      const resp = await papi.export_plug_state(id);
      if (!resp.ok) {
        return;
      }
  
      let anchor = document.createElement("a");
      document.body.appendChild(anchor);
  
      let data = resp.data;
      if (typeof data === "object") {
        data = JSON.stringify(resp.data, null, 4);
      }
  
      let objectUrl = window.URL.createObjectURL(new Blob([data]));
  
      anchor.href = objectUrl;
      anchor.download = "plug_state.json";
  
      anchor.click();
      window.URL.revokeObjectURL(objectUrl);
  
      mode = "RESULT";
      loading = false;
    };
  </script>
  
  {#if loading}
    <LoadingSpinner classes="" />
  {:else if mode === "SELECT"}
    <div class="p-2 text-center overflow-y-auto">
      <h3 class="mb-2 text-2xl font-bold text-gray-800">Export State</h3>
      <p class="text-gray-500">Do you want to export plug states ?</p>
  
      <div class="mt-6 flex justify-end gap-2">
        <button
          on:click={() => load()}
          class="py-2 px-4 inline-flex justify-center items-center gap-2 rounded-md border border-transparent font-semibold bg-blue-500 text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-all text-sm"
        >
          Ok
        </button>
  
        <button
          on:click={() => app.utils.small_modal_close()}
          class="py-2 px-4 inline-flex justify-center items-center gap-2 rounded-md border border-transparent font-semibold bg-red-500 text-white hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition-all text-sm"
        >
          Cancel
        </button>
      </div>
  
      
    </div>
  {:else}
    <div class="p-2 text-center overflow-y-auto">
      <h3 class="mb-2 text-2xl font-bold text-gray-800">Done</h3>
    </div>
  {/if}
  
