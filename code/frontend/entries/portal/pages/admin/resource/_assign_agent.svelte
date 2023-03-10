<script lang="ts">
  import { LoadingSpinner } from "../../../../adapter_editor/easypage/core";
  import type { PortalService } from "../core";
  import ActionPicker from "../../admin/core/action_picker.svelte";

  export let service: PortalService;
  export let rid: string;

  let loading = true;

  let plugs = [];
  let agents = [];
  let plug_id = "";
  let agent_id = "";

  let mode: "pick_plug" | "pick_agent" = "pick_plug";

  const papi = service.api_manager.get_admin_plug_api();

  const load = async () => {
    loading = true;

    const resp = await papi.list_plug();
    if (!resp.ok) {
      return;
    }

    plugs = resp.data;
    mode = "pick_plug";
    loading = false;
  };

  const pick_plug = async (sdata) => {
    plug_id = sdata["data"]["id"];

    loading = true;
    const resp = await papi.list_agent(plug_id);
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    mode = "pick_agent"
    loading = false;
    agents = resp.data;
  };

  const pick_agent = async (sdata) => {
    agent_id = sdata["data"]["id"];

    service.nav.admin_agent_res_new(plug_id, agent_id, {
      resource_id: rid,
    });

    service.utils.small_modal_close();
  };

  load();
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else if mode === "pick_plug"}
  <ActionPicker
    actions={plugs.map((p) => ({
      action: pick_plug,
      icon: "hashtag",
      info: p["name"],
      name: p["id"],
      data: p,
    }))}
    title="Pick Plug"
  />
{:else if mode === "pick_agent"}
  <ActionPicker
    actions={agents.map((a) => ({
      action: pick_agent,
      icon: "hashtag",
      info: a["name"],
      name: a["id"],
      data: a,
    }))}
    title="Pick Agent"
  />
{:else}
  <div>end</div>
{/if}
