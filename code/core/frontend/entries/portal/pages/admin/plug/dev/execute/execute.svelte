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

  const load = async (pid: string, aid: string) => {
    if (!pid || !aid) {
      return;
    }

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
  <IframeExecute exec_data={data} name="TODO" secret_id={sid} />
{/if}
