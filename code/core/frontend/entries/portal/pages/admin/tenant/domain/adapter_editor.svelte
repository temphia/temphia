<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import type { PortalService } from "../../core";

  export let did = $params.did;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_tenant_api();


  const templatedSrc = ``


</script>

<div>Adapter Editor</div>
