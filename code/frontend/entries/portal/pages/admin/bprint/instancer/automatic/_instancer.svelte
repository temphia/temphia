<script lang="ts">
  import Stepper from "../../../../../xcompo/stepper/stepper.svelte";
  import Step from "../../../../../xcompo/stepper/step.svelte";
  import ItemTable from "./_item_table.svelte";
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../../core";

  export let bid;
  export let bundle_objects;
  export let instancer_type;

  const app = getContext("__app__") as PortalService;
  const bapi = app.api_manager.get_admin_bprint_api();

  let instanceing = false;
  let message = "";
  let instanced_resp = false;
  let instancedData;

  const instance = async () => {
    instanceing = true;
    const resp = await bapi.instance(bid, {
      auto: true,
      instancer_type,
      file: "schema.json",
    });

    if (!resp.ok) {
      message = resp.data;
      instanced_resp = true;
      return;
    }

    instancedData = resp.data;
    instanceing = false;
    instanced_resp = true;
  };

  const onNextHandler = (e: any) => {
    switch (e.detail["step"]) {
      case 0:
        instance();
        console.log("@first");
        break;
      default:
        break;
    }

    console.log("event:next", e.detail);
  };
</script>

<div
  class="card p-2 bg-white border shadow mx-auto"
  style="max-width: 750px;"
>
  <Stepper on:next={onNextHandler}>
    <Step back_locked locked={instanceing}>
      <svelte:fragment slot="header">Blueprint items instance</svelte:fragment>
      <ItemTable items={bundle_objects} />
    </Step>

    <Step back_locked={instanced_resp}>
      <svelte:fragment slot="header"
        >{instanceing ? "Instancing" : "Finished"}</svelte:fragment
      >

      {#if instanceing}
        <LoadingSpinner classes="" />
      {:else if message}
        <p class="bg-red-500">
          {message}
        </p>
      {:else}
        <details>
          <summary>Response</summary>
          <pre><code>{JSON.stringify(instancedData, null, 2)}</code></pre>
        </details>
      {/if}
    </Step>
  </Stepper>
</div>
