<script lang="ts">
  import type { Launcher } from "../services/launcher";

  export let launcher: Launcher;

  const display_state = launcher.display_state;
</script>

<div
  class={$display_state === "HIDDEN"
    ? "hidden"
    : $display_state === "FLOATING"
    ? "floating"
    : "not-floating"}
>
  <iframe
    title="test"
    class="border border-red-600 w-full h-full"
    src="http://www.example.com/"
  />
</div>

<style>
  .hidden {
    display: none;
    width: 0px;
    height: 0px;
  }

  .floating {
    display: block;
    position: fixed;
    width: 90vw;
    height: 90vh;
    z-index: 1000;
  }

  .not-floating {
    display: block;
    width: 90vw;
    height: 90vh;
  }
</style>
