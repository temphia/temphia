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
    const resp = await api.domain_issue_adapter_editor(did);
    if (!resp.ok) {
      message = resp.data;
      loading = false;
      return;
    }

    source = GenerateSRC({
      adapter_editor_token: resp.data["token"],
      adapter_type: resp.data["adapter_type"],
      base_url: app.api_manager.base_url,
      did: did,
      tenant_id: app.api_manager.tenant_id,
    });
    loading = false;
  };

  load();
</script>

<div class="h-full w-full p-2">
  {#if !loading}
    <p class="text-color-red-500">{message}</p>
    <iframe
      srcdoc={source}
      title="Adapter Editor"
      class="w-full h-full transition-all bg-slate-200"
      height="100"
      width="100"
    />
  {:else}
    <LoadingSpinner />
  {/if}
</div>
