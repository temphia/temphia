<script lang="ts">
  import { LoadingSpinner, PortalService } from "../../../core";
  import ActionPicker from "../../../core/action_picker.svelte";

  export let service: PortalService;
  export let pid: string;
  export let aid: string;

  let loading = true;
  let to_plug_id = "";
  let to_agent_id = "";

  let plugs = [];
  let agents = [];

  let mode: "pick_plug" | "pick_agent" = "pick_plug";

  const load = async () => {
    const api = service.api_manager.get_admin_plug_api();
    const resp = await api.list_plug();
    if (!resp.ok) {
      return;
    }
    plugs = resp.data;
    loading = false
  };

  const pick_plug = async (picked) => {
    loading = true;
    mode = "pick_agent";
    to_plug_id = picked.data.id;

    const api = service.api_manager.get_admin_plug_api();
    const resp = await api.list_agent(to_plug_id);
    if (!resp.ok) {
      return;
    }

    agents = resp.data;
    loading = false;
  };

  const pick_agent = async (picked) => {
    loading = true;
    to_agent_id = picked.data.id;

    service.nav.admin_agent_link_new(pid, aid, {
      to_agent_id,
      to_plug_id,
    });

    service.utils.small_modal_close();
  };

  load();
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else if mode === "pick_plug"}
  <ActionPicker
    actions={plugs.map((pg) => ({
      action: pick_plug,
      icon: "hashtag",
      info: pg.name,
      name: pg.id,
      data: pg,
    }))}
    title="Pick Plug"
  />
{:else if mode === "pick_agent"}
  <ActionPicker
    actions={agents.map((ag) => ({
      action: pick_agent,
      icon: "hashtag",
      info: ag.name,
      name: ag.id,
      data: ag,
    }))}
    title="Pick Agent"
  />
{/if}
