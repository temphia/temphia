<script lang="ts">
  import Tailwind from "../common/_tailwind.svelte";
  import Layout from "./core/layout.svelte";
  import Element from "./elements/element.svelte";
  // import Dashui from "./ref/old/dashui.svelte";
  // import Ref2 from "./ref/ref2.svelte";
  import { DashClient } from "./service";

  export let env;

  const dash = new DashClient({
    env: env,
  });

  dash.init();

  let store = dash._state;
</script>

<Tailwind />

<div class="h-full w-full flex flex-col bg-gray-50 space-y-2 p-2 ">
  {#if $store.loaded}
    <div>
      <div class="w-full p-1 flex justify-center bg-white shadow rounded">
        <h3 class="text-xl font-semibold text-gray-500">
          {$store.inner.name}
        </h3>
      </div>
    </div>

    {#each $store.inner.sections as section}
      {#if section.layout === "flex-auto"}
        <div class="flex flex-wrap justify-between space-x-2 space-y-1">
          {#each section.panels as panel}
            <Layout width={panel.width} height={0}>
              <Element {panel} data={$store.inner.data[panel.source] || {}} />
            </Layout>
          {/each}
        </div>
      {:else}
        <div>Not impl</div>
      {/if}
    {/each}
  {:else}
    <div>Loading</div>
  {/if}
</div>

<div class="hidden">
  <div class="w-28">xm</div>
  <div class="w-32">sm</div>
  <div class="w-48">md</div>
  <div class="w-64">lg</div>
  <div class="w-72">xl</div>
  <div class="w-96">2xl</div>
</div>
