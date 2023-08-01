<script lang="ts">
  import { getContext } from "svelte";
  import { LoadingSpinner } from "../admin/core";
  import type { PortalService } from "../../services";

  const app: PortalService = getContext("__app__");

  app.api_manager.self_data.get_data_sources().then((sources) => {
    app.nav.data_groups(sources[0]);
  });
</script>

<LoadingSpinner />
