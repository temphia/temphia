<script lang="ts">
  import { getContext } from "svelte";
  import { PortalService, LoadingSpinner } from "../core";

  const app = getContext("__app__") as PortalService;

  app.api_manager.cache.get_data_sources().then((_sources) => {
    app.nav.admin_data_groups(_sources[0]);
  });
</script>

<LoadingSpinner />
