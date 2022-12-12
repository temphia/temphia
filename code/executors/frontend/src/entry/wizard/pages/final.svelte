<script lang="ts">
  import Layout from "../core/layout.svelte";
  import type { WizardManager } from "../service";
  export let manager: WizardManager;
  const store = manager._state;
</script>

<Layout title={manager.wizard_title} showButtons={false}>
  <div>Wizard Finished</div>

  {#if $store.message}
    <div class="p-1 border bg-yellow-100 rounded mt-10">
      {$store.message}
    </div>
  {/if}
</Layout>
