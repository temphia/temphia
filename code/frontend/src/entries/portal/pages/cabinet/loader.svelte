<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import { LoadingSpinner } from "../admin/core";

  const app: PortalService = getContext("__app__");

  const load = async () => {
    const cservice = app.get_cabinet_service();
    const sources = await cservice.get_cab_sources();

    if (!sources || sources.length === 0) {
      console.log("no cabinet sources found");
      return;
    }
    app.nav.cab_folders(sources[0]);
  };

  load();
</script>

<LoadingSpinner />
