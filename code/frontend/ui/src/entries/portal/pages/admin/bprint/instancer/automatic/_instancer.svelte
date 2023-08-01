<script lang="ts">
  import Stepper from "../../../../../../xcompo/stepper/stepper.svelte";
  import Step from "../../../../../../xcompo/stepper/step.svelte";
  import ItemTable from "./_item_table.svelte";
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../../core";
  import FinalTable from "./_final_table.svelte";

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
        break;
      default:
        break;
    }

    console.log("event:next", e.detail);
  };
</script>

<div class="card p-2 bg-white border shadow mx-auto" style="max-width: 750px;">
  <Stepper on:next={onNextHandler} buttonCompleteLabel="">
    <Step back_locked locked={instanceing}>
      <svelte:fragment slot="header">Blueprint items</svelte:fragment>
      <ItemTable items={bundle_objects} />

      <p class="text-sm italic my-1">
        Click next to instance items automatically or click manual to instance individual items.
      </p>
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
        <FinalTable installed_items={instancedData["objects"] || {}} />
      {/if}
    </Step>
  </Stepper>
</div>
