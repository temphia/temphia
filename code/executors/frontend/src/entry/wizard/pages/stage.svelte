<script lang="ts">
  import Layout from "../core/layout.svelte";
  import Processing from "../core/processing.svelte";
  import Element from "../elements/element.svelte";
  import type { Manager } from "../service/wizard_types";

  export let manager: Manager;
  const state = manager.get_state();
  const data_sources = $state.data_sources;
</script>

<Layout
  title={manager.wizard_title}
  showButtons={$state.flowState !== "STAGE_PROCESSING"}
  next={manager.stage_next}
>
  {#if $state.flowState === "STAGE_PROCESSING"}
    <Processing />
  {:else}
    {#each $state.fields || [] as field}
      <div class="relative my-4">
        <Element
          {field}
          {data_sources}
          prev_data={$state.prev_data}
          fieldstore={manager.get_field_store(field["name"])}
          error={$state.errors[field["name"]]}
        />
      </div>
    {/each}
  {/if}
</Layout>
