<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { generateId } from "../../../../../../../lib/utils";
  import IframeExecute from "../../../../../launcher/hostproxy/iframe/iframe_execute.svelte";
  import type { ExecInstanceOptions } from "../../../../../services/engine/exec_type";
  import { LoadingSpinner, PortalService } from "../../../core";

  const app = getContext("__app__") as PortalService;
  const eapi = app.api_manager.get_engine_api();

  let data: ExecInstanceOptions;
  let sid = "";
  let loading = true;
  let bootloader = "";

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

  load($params.pid, $params.aid);
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <IframeExecute
    exec_data={data}
    name="{$params.pid}/{$params.aid}"
    secret_id={sid}
    tenant_id={app.options.tenant_id}
    {bootloader}
  />
{/if}
