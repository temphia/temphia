<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { generateId } from "../../../../../../../lib/utils";
  import IframeExecute from "../../../../../launcher/hostproxy/iframe/iframe_execute.svelte";
  import type { ExecInstanceOptions } from "../../../../../services/engine/exec_type";
  import { LoadingSpinner, PortalService } from "../../../core";

  export let pid = $params.pid;
  export let aid = $params.aid;

  const app = getContext("__app__") as PortalService;
  const eapi = app.api_manager.get_engine_api();

  // asas

  let data: ExecInstanceOptions;
  let sid = "";
  let loading = true;
  let bootloader = "";

  let iframe;

  let chan = new MessageChannel();
  chan.port1.onmessage = () => {};

  const load = async (pid: string, aid: string) => {
    if (!pid || !aid) {
      return;
    }

    bootloader = await app.launcher.get_bootloader();

    const resp = await eapi.launch_admin({
      plug_id: pid,
      agent_id: aid,
    });
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }

    data = resp.data;
    sid = generateId();
    loading = false;
  };

  load(pid, aid);
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <IframeExecute
    bind:iframe
    exec_data={data}
    name="{pid}/{aid}"
    secret_id={sid}
    tenant_id={app.options.tenant_id}
    {bootloader}
    startup_payload={{}}
    on:load={() => {
      iframe.contentWindow.postMessage("port_transfer", "*", [chan.port2]);
    }}
  />
{/if}
