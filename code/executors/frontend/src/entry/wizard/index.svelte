<script lang="ts">
  import Tailwind from "../common/_tailwind.svelte";
  import Splash from "./pages/splash.svelte";
  import Stage from "./pages/stage.svelte";
  import Final from "./pages/final.svelte";
  import { WizardManager } from "./service";
  import Modal from "./core/modal.svelte";

  export let env;
  export let options = null;

  const manager = new WizardManager(env, options);
  const store = manager._state;
  manager.init();

  $: _curr_state = $store.flowState;
</script>

<Tailwind />

<Modal>
  {#if _curr_state === "NOT_LOADED"}
    <div>Loading..</div>
  {:else if _curr_state === "SPLASH_LOADED"}
    {#key $store.epoch}
      <Splash {manager} />
    {/key}
  {:else if _curr_state === "FINISHED"}
    <Final {manager} />
  {:else}
    {#key $store.epoch}
      <Stage {manager} />
    {/key}
  {/if}
</Modal>
