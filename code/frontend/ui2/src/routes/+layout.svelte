<script lang="ts">
  import "@skeletonlabs/skeleton/themes/theme-skeleton.css";
  import "@skeletonlabs/skeleton/styles/skeleton.css";
  import { initRegistry } from "$lib/services/portal/registry/putils";

  initRegistry();
</script>

<slot />
