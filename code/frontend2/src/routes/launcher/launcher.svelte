<script lang="ts">
  import { getContext, onDestroy, onMount } from "svelte";
  import type { PortalService } from "$lib/services/portal/portal";

  const app = getContext("__app__") as PortalService;
  const launcher = app.launcher;

  onMount(() => {
    launcher.plane_show();
  });

  onDestroy(() => {
    launcher.plane_hide();
  });
</script>
