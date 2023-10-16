<script lang="ts">
  import { getContext } from "svelte";
  import { generateId } from "$lib/utils";

  import Step from "$lib/compo/stepper/step.svelte";
  import Stepper from "$lib/compo/stepper/stepper.svelte";
  import type { PortalService } from "$lib/core";
  import { LoadingSpinner } from "$lib/core";

  import Detail from "./detail.svelte";

  export let data;
  export let source;

  const app: PortalService = getContext("__app__");

  let final = false;

  let importing = true;

  let bid = "";
  let version;

  const importPreform = async () => {
    const bapi = app.api_manager.get_admin_bprint_api();
    const resp = await bapi.import({
      slug: data["slug"],
      group: data["type"] || data["group"],
      source: Number(source),
      version,
      new_id: (data["slug"] || "").replaceAll(".", "_") + "_" + generateId(),
    });
    if (!resp.ok) {
      console.log("@resp", resp);
      return;
    }

    bid = resp.data;
    importing = false;
  };

  const onNextHandler = (e: any) => {
    console.log("event:next", e.detail);

    switch (e.detail["step"]) {
      case 0:
        importPreform();
        break;
      case 1:
        app.nav.admin_bprint_instancer(bid);

      default:
        break;
    }
  };
</script>

<div class="w-full bg-gray-50 h-full py-4 px-1">
  <div
    class="card p-4 text-token border shadow mx-auto my-4 bg-white"
    style="max-width: 750px;"
  >
    <Stepper buttonCompleteLabel={""} on:next={onNextHandler}>
      <Step back_locked={true} locked={!version}>
        <svelte:fragment slot="header">Import</svelte:fragment>

        <svelte:fragment>
          <Detail {data} bind:version />
        </svelte:fragment>
      </Step>

      <Step locked={importing} back_locked={!importing}>
        <svelte:fragment slot="header"
          >{importing ? "Importing" : "Imported"}</svelte:fragment
        >
        <svelte:fragment>
          {#if importing}
            <LoadingSpinner classes="" />
          {:else}
            <p>
              Blueprint is import, click next to instance or click

              <button
                on:click={() => {
                  app.nav.admin_bprints();
                }}
                class="text-blue-600 underline">here</button
              > explore all blueprints.
            </p>
          {/if}
        </svelte:fragment>
      </Step>
      <Step back_locked={final}>
        <svelte:fragment slot="header">Final</svelte:fragment>
        <svelte:fragment />
      </Step>
    </Stepper>
  </div>
</div>
