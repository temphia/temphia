<script lang="ts">
  import Layout from "../core/layout.svelte";
  import Element from "../elements/element.svelte";

  import type { Manager } from "../service/wizard_types";
  export let manager: Manager;
  const store = manager.get_state();
  const data_sources = $store.data_sources;

  const start = () => {
    manager.splash_next();
  };
</script>

<Layout title={manager.wizard_title} showButtons={false}>
  {#if $store.message}
    <div class="p-1 border bg-yellow-100 rounded">
      {$store.message}
    </div>
  {/if}

  {#each $store.fields || [] as field}
    <div class="relative my-4">
      <Element
        {field}
        {data_sources}
        fieldstore={manager.get_field_store(field["name"])}
      />
    </div>
  {/each}

  <div class="p-1 flex justify-center">
    <button
      on:click={start}
      class="py-2 px-4 border rounded-md cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700"
      >Start</button
    >
  </div>
</Layout>
