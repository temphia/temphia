<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../lib/app/portal";

  const app: PortalApp = getContext("__app__");

  app.get_dyn_sources().then((resp) => {
    if (!resp) {
      return;
    }
    app.navigator.goto_dtable_source(resp[0]);
  });
</script>
