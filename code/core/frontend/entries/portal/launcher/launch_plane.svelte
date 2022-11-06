<script lang="ts">
  import type { Writable } from "svelte/store";

  export let store: Writable<boolean>;

  $: console.log("#STORE", $store);

  // let display_state: Writable<"HIDDEN" | "FLOATING" | "NOT_FLOATING"> =
  //   writable("HIDDEN");
</script>

<div style="display: {$store ? 'block' : 'none'};">Launch Plane</div>

<!-- <style>
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
  }

  .not-floating {
    display: block;
    width: 90vw;
    height: 90vh;
  }
</style> -->
