<script lang="ts">
  import { LoadingSpinner, PortalService } from "../../core";

  export let id;
  export let app: PortalService;

  let loading = false;
  let clean_old = false;

  let filedata;

  let mode: "SELECT" | "RESULT" = "SELECT";

  const preform = async () => {
    loading = true;

    if (!filedata) {
      return;
    }

    const papi = app.api_manager.get_admin_plug_api();
    const resp = await papi.import_plug_state(id, clean_old, filedata);
    if (!resp.ok) {
      return;
    }

    mode = "RESULT";
    loading = false;
  };
</script>

<div class="p-4 text-center overflow-y-auto flex flex-col gap-2">
  {#if loading}
    <LoadingSpinner classes="" />
  {:else if mode === "SELECT"}
    <h3 class="mb-2 text-center text-2xl font-bold text-gray-800">Import State</h3>

    <label>
      JSON file
      <input
        type="file"
        on:change={(ev) => {
          const file = ev.target["files"][0];

          if (file) {
            const reader = new FileReader();

            reader.onload = (event) => {
              filedata = event.target.result;
              console.log(filedata);
            };
            reader.readAsText(file);
          }
        }}
      />
    </label>

    <label>
      Clean Previous States
      <input type="checkbox" bind:checked={clean_old} />
    </label>

    <div class="mt-6 flex justify-end gap-2">
      {#if filedata}
        <button
          on:click={preform}
          class="py-2 px-4 inline-flex justify-center items-center gap-2 rounded-md border border-transparent font-semibold bg-blue-500 text-white hover:bg-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 transition-all text-sm"
        >
          Ok
        </button>
      {/if}

      <button
        on:click={() => app.utils.small_modal_close()}
        class="py-2 px-4 inline-flex justify-center items-center gap-2 rounded-md border border-transparent font-semibold bg-red-500 text-white hover:bg-red-600 focus:outline-none focus:ring-2 focus:ring-red-500 focus:ring-offset-2 transition-all text-sm"
      >
        Cancel
      </button>
    </div>
  {:else}
    <h3 class="mb-2 text-2xl font-bold text-gray-800 text-center">Done</h3>
  {/if}
</div>
