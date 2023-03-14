<script lang="ts">
  import Stepper from "./stepper/stepper.svelte";
  import Step from "./stepper/step.svelte";

  let locked: boolean = true;
  function onNextHandler(e: any): void {
    console.log("event:next", e.detail);
  }
  function onBackHandler(e: any): void {
    console.log("event:prev", e.detail);
  }
  function onStepHandler(e: any): void {
    console.log("event:step", e.detail);
  }
  function onCompleteHandler(e: any): void {
    console.log("event:complete", e.detail);
    alert("Complete!");
  }
</script>

<div class="w-full bg-gray-50 h-full py-4 px-1">
  <div
    class="card p-4 text-token border shadow  mx-auto my-4"
    style="max-width: 750px;"
  >
    <Stepper
      on:next={onNextHandler}
      on:back={onBackHandler}
      on:step={onStepHandler}
      on:complete={onCompleteHandler}
    >
      <Step>
        <svelte:fragment slot="header">Get Started!</svelte:fragment>
        <p>
          This example Stepper will teach you how to use this component. Tap <u
            >next</u
          > to proceed to the next step.
        </p>
      </Step>
      <Step>
        <svelte:fragment slot="header">Going Back.</svelte:fragment>
        <p>
          The current step progress will be tracked at the top. Tap <u>back</u> to
          return to the previous step.
        </p>
      </Step>
      <Step {locked}>
        <svelte:fragment slot="header">A Locked Step.</svelte:fragment>
        <p>
          This Step component uses the <code>locked</code> property to prevent progress.
          This is ideal for multi-step forms, such as registration. For now we'll
          simulate a successful validation condition using the toggle below.
        </p>
        <div class="card !bg-transparent p-4 text-center">
          <!-- <SlideToggle name="locked-state" bind:checked={locked}
            >This step is <u>{locked ? "Locked" : "Unlocked"}</u></SlideToggle
          > -->
        </div>
      </Step>
      <Step>
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
      <Step>
        <svelte:fragment slot="header">Almost Done.</svelte:fragment>
        <p>
          A Complete button will appear on the last step. When the step is
          unlocked and the button pressed, an <code>on:complete</code> event will
          be fired. You can use this trigger to submit form data to a server.
        </p>
      </Step>
    </Stepper>
  </div>
</div>
