<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";

  const app: PortalService = getContext("__app__");
  const api = app.api_manager.get_admin_data_api();



  app.api_manager.self_data.get_data_sources().then((sources) => {
    


  })


</script>

<div>Data Loader</div>
