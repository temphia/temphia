<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../../services";
  const app = getContext("__app__") as PortalService;
  
  const rapi = app.api_manager.get_admin_repo_api();

</script>

<div>Repos</div>
