<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { LoadingSpinner, PortalService } from "../../../core";
  import Iface from "./iface/iface.svelte";

  const app = getContext("__app__") as PortalService;
  const papi = app.api_manager.get_admin_plug_api();
  const bapi = app.api_manager.get_admin_bprint_api();

  let data;
  let loading = true;

  const load = async () => {
    const resp1 = papi.get_agent($params.pid, $params.aid);
    const resp2 = papi.get_plug($params.pid);

    const _resp1 = await resp1;
    const _resp2 = await resp2;

    if (!_resp1.ok || !_resp2.ok) {
      return;
    }

    const ifacefile = _resp1.data["iface_file"] || "";
    const bid = _resp2.data["bprint_id"] || "";

    const resp = await bapi.get_file(bid, ifacefile);
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="p-4">
    <div class="p-4 rounded bg-white">
      <Iface {data} />
    </div>
  </div>
{/if}
