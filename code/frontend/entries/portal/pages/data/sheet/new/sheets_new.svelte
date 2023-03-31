<script lang="ts">
  import Step from "../../../play/stepper/step.svelte";
  import Stepper from "../../../play/stepper/stepper.svelte";
  import Template from "./_template.svelte";

  let final = false;

  let name = "";
  let info = "";
  let template = "";



  function onNextHandler(e: any): void {
    console.log("event:next", e.detail["step"]);
    if (e.detail["step"] === 1) {
      final = true;
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
      <Step locked={!name || !info}>
        <svelte:fragment slot="header">New Sheet</svelte:fragment>

        <svelte:fragment>
          <div class="flex-col flex py-3">
            <label for="" class="pb-2 text-gray-700 font-semibold">Name</label>
            <input
              type="text"
              bind:value={name}
              class="p-2 rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            />
          </div>

          <div class="flex-col flex py-3">
            <label for="" class="pb-2 text-gray-700 font-semibold">Info</label>
            <textarea
              bind:value={info}
              class="p-2 rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            />
          </div>
        </svelte:fragment>
      </Step>

      <Step back_locked={final} locked={!template}>
        <svelte:fragment slot="header">Select Template</svelte:fragment>

        <svelte:fragment>
          <Template bind:template={template} />
        </svelte:fragment>
      </Step>
      <Step back_locked={final}>
        <svelte:fragment slot="header">Finished.</svelte:fragment>
        <p>Sheet is ready. Go explore.</p>
      </Step>
    </Stepper>
  </div>
</div>
