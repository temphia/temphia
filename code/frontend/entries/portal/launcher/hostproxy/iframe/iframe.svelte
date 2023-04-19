<script lang="ts">
  import type { Instance } from "../../../services/engine/launcher";
  import LoadingSpinner from "../../../../xcompo/common/loading_spinner.svelte";
  import IframeInner from "./iframe_execute.svelte";
  import type { PortalService } from "../../../services";
  import { getContext } from "svelte";
  import type { ExecInstanceOptions } from "../../../services/engine/exec_type";

  export let options: Instance;

  const app = getContext("__app__") as PortalService;
  const eapi = app.api_manager.get_engine_api();

  let loading = true;
  let data: ExecInstanceOptions;
  let bootloader = "";
  let iframe: HTMLIFrameElement;
  let startup_payload: any;

  const load = async () => {
    bootloader = await app.launcher.get_bootloader();

    const resp = await eapi.launch_target({
      target_id: Number(options.target_id),
      target_type: options.target_type,
    });
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    data = resp.data;

    startup_payload = app.launcher.last_startup_payload || {};
    app.launcher.last_startup_payload = null;
    loading = false;
  };

  const on_agent_load = () => {
    iframe.contentWindow.postMessage("port_transfer", "*", [
      options.channel.port2,
    ]);
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <IframeInner
    bind:iframe
    exec_data={data}
    name={options.name}
    secret_id={options.id}
    tenant_id={app.options.tenant_id}
    {bootloader}
    {startup_payload}
    on:load={on_agent_load}
  />
{/if}
