<script lang="ts">
  import Step from "../../../play/stepper/step.svelte";
  import Stepper from "../../../play/stepper/stepper.svelte";

  let final = false;

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
    class="card p-4 text-token border shadow  mx-auto my-4"
    style="max-width: 750px;"
  >
    <Stepper
      buttonCompleteLabel={""}
      on:next={onNextHandler}
      on:back={onBackHandler}
      on:step={onStepHandler}
      on:complete={onCompleteHandler}
    >
      <Step>
        <svelte:fragment slot="header">Going Back.</svelte:fragment>
        <p>
          The current step progress will be tracked at the top. Tap <u>back</u> to
          return to the previous step.
        </p>
      </Step>

      <Step back_locked={final}>
        <svelte:fragment slot="header">Long Form Content.</svelte:fragment>
        <p>The steps will expand to fit content of any length.</p>
        <p>
          Lorem ipsum dolor sit amet consectetur adipisicing elit. Itaque vel
          expedita porro vero, saepe dicta repellendus facilis ab accusamus
          unde, tempora ut nobis eum. Veniam, architecto corrupti. Lorem ipsum
          dolor sit amet consectetur adipisicing elit. Itaque vel expedita porro
          vero, saepe dicta repellendus facilis ab accusamus unde, tempora ut
          nobis eum. Veniam, architecto corrupti. Lorem ipsum dolor sit amet
          consectetur adipisicing elit. Itaque vel expedita porro vero, saepe
          dicta repellendus facilis ab accusamus unde, tempora ut nobis eum.
          Veniam, architecto corrupti. Lorem ipsum dolor sit amet consectetur
          adipisicing elit. Itaque vel expedita porro vero, saepe dicta
          repellendus facilis ab accusamus unde, tempora ut nobis eum. Veniam,
          architecto corrupti. Lorem ipsum dolor sit amet consectetur
          adipisicing elit. Itaque vel expedita porro vero, saepe dicta
          repellendus facilis ab accusamus unde, tempora ut nobis eum. Veniam,
          architecto corrupti.
        </p>
      </Step>
      <Step back_locked={final}>
        <svelte:fragment slot="header">Finished.</svelte:fragment>
        <p>Sheet is read. Go explore.</p>
      </Step>
    </Stepper>
  </div>
</div>
