<script lang="ts">
  import type { PortalService } from "../../core";
  import { instance_helper } from "./index";

  export let bid: string;
  export let app: PortalService;
  export let bprint: object;

  let bundle_objects;

  (async () => {
    const bapi = app.api_manager.get_admin_bprint_api();
    const resp = await bapi.get_file(bid, "schema.json");
    if (resp.status !== 200) {
      return;
    }
    bundle_objects = resp.data;
  })();
</script>

{#if bundle_objects}
  <h3>Bundle Objects</h3>

  <div class="flex flex-col gap-2">
    {#each Object.entries(bundle_objects) as [bkey, be]}
      <div class="p-2 rounded border flex justify-evenly">
        <p>{bkey}</p>

        <span class="p-1 rounded bg-pink-200">{be["type"]}</span>

        <button
          class="p-1 bg-blue-500 hover:bg-blue-700 text-white rounded"
          on:click={() => {
            instance_helper(app, be["type"], bprint, be["file"]);
            app.utils.small_modal_close();
          }}>Instance</button
        >
      </div>
    {/each}
  </div>
{/if}
