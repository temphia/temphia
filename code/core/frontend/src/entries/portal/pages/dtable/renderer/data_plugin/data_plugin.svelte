<script lang="ts">
  import Layout from "./layout.svelte";
  export let show;
  export let plugin_target;
</script>

<Layout bind:show>
  <div class="h-full w-full" bind:this={plugin_target} />
</Layout>
