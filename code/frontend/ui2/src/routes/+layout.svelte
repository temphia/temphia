<script lang="ts">
  import "../app.postcss";

  import { initRegistry } from "$lib/services/portal/registry/putils";
  import { initializeStores } from "@skeletonlabs/skeleton";

  initRegistry();
  initializeStores();
</script>

<slot />
