<script lang="ts">
  import type { PortalApp } from "../../../../../lib/app/portal";
  import { instance } from "./bprint_util";

  export let bid: string;
  export let app: PortalApp;
  export let bprint: object;

  let bundle_objects;

  (async () => {
    const bapi = await app.get_apm().get_bprint_api();
    const resp = await bapi.bprint_get_file(bid, "schema.json");
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
            instance(app, be["type"], bprint, be["file"]);
            app.simple_modal_close()
          }}>Instance</button
        >
      </div>
    {/each}
  </div>
{/if}
