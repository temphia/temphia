<script lang="ts">
  import "../app.postcss";

  import { initRegistry } from "$lib/services/portal/registry/putils";
  import { initializeStores } from "@skeletonlabs/skeleton";

  initRegistry();
  initializeStores();
</script>

<svelte:head>
  <link rel="icon" type="image/png" href="/z/assets/static/logo.png">
</svelte:head>

<slot />
