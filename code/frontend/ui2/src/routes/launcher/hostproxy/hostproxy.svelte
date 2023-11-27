<script lang="ts">
  import { getContext } from "svelte";
  import { LoadingSpinner, type PortalService } from "$lib/core";
  import Eframe from "$lib/compo/eframe/eframe.svelte";
  import type { Instance } from "$lib/services/portal/launcher/launcher";
  import { BuildExecURL } from "$lib/compo/eframe";

  export let options: Instance;

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_engine_api();

  let loading = true;
  let url = "";
  let token = "";
  let data = {};

  const load = async () => {
    const resp = await api.launch_target({
      target_id: Number(options.target_id),
    });
    if (!resp.ok) {
      return;
    }

    url = BuildExecURL(resp.data);

    token = resp.data["token"];
    data = resp.data;

    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <Eframe
    exec_data={{
      agent_id: data["agent_id"],
      plug_id: data["plug_id"],
      tenant_id: data["tenant_id"],
      exec_token: token,
      startup_payload: null,
    }}
    {url}
    name={options.name || "App"}
  />
{/if}
