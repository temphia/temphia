<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner, PortalService } from "../../core";
  import { GenerateSRC } from "./template";

  export let did = $params.did;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_tenant_api();

  let loading = false;
  let message = "";
  let source = "";

  const load = async () => {
    source = GenerateSRC({
      adapter_editor_token: "",
      adapter_type: "",
      base_url: "",
      did: "",
      tenant_id: "",
    });
  };
</script>

{#if !loading}
  <iframe srcdoc={source} title="Adapter Editor" />
{:else}
  <LoadingSpinner />
{/if}
