<script lang="ts">
  import { getContext } from "svelte";
  import { generateId } from "../../../../../lib/utils";
  import Step from "../../../../xcompo/stepper/step.svelte";
  import Stepper from "../../../../xcompo/stepper/stepper.svelte";
  import type { PortalService } from "../../../services";
  import { LoadingSpinner } from "../../admin/core";

  import Detail from "./detail.svelte";

  export let data;
  export let source;

  const app: PortalService = getContext("__app__");

  let final = false;

  let importing = true;

  const importPreform = async () => {
    const bapi = app.api_manager.get_admin_bprint_api();
    const resp = await bapi.import({
      slug: data["slug"],
      group: data["type"] || data["group"],
      source: Number(source),
      new_id: (data["slug"] || "").replaceAll(".", "_") + "_" + generateId(),
    });
    if (!resp.ok) {
      console.log("@resp", resp);
      return;
    }
  };

  function onNextHandler(e: any): void {
    console.log("event:next", e.detail["step"]);
    if (e.detail["step"] === 1) {
      importPreform();
    }
  }
  function onBackHandler(e: any): void {
    console.log("event:prev", e.detail);
  }
  function onStepHandler(e: any): void {
    console.log("event:step", e.detail);
  }

  function onCompleteHandler(e: any): void {
    console.log("event:complete", e.detail);
  }
</script>

<div class="w-full bg-gray-50 h-full py-4 px-1">
  <div
    class="card p-4 text-token border shadow  mx-auto my-4 bg-white"
    style="max-width: 750px;"
  >
    <Stepper
      buttonCompleteLabel={""}
      on:next={onNextHandler}
      on:back={onBackHandler}
      on:step={onStepHandler}
      on:complete={onCompleteHandler}
    >
      <Step back_locked={true}>
        <svelte:fragment>
          <Detail {data} />
        </svelte:fragment>
      </Step>

      <Step locked={!importing}>
        <svelte:fragment slot="header"
          >{importing ? "Importing" : "Imported"}</svelte:fragment
        >
        <svelte:fragment>
          {#if importing}
            <LoadingSpinner classes="" />
          {:else}
            <p>
              Blueprint is import, click next to instance or click <button
                >here</button
              > explore
            </p>
          {/if}
        </svelte:fragment>
      </Step>
      <Step back_locked={final}>
        <svelte:fragment slot="header">Final</svelte:fragment>
        <!-- {#if preforming}
          <LoadingSpinner classes="" />
        {:else if message}
          <p class="text-red-500">{message}</p>
        {:else}
          <p>Sheet is ready. Go explore.</p>
        {/if} -->
      </Step>
    </Stepper>
  </div>
</div>
