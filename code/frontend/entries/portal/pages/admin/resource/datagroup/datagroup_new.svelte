<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../core";

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_resource_api();
</script>

<div>Datagroup Edit</div>
