<script lang="ts">
  import { LoadingSpinner } from "$lib/core";
  import type { PortalService } from "$lib/core";
  import ActionPicker from "$lib/core/action_picker.svelte";
  import { TargetAppTypeUserGroupApp } from "../../target";

  export let service: PortalService;

  let loading = true;
  let ugroup = "";
  let plug_id = "";
  let agent_id = "";

  let groups = [];
  let plugs = [];
  let agents = [];

  let mode: "pick_group" | "pick_plug" | "pick_agent" = "pick_group";

  const load = async () => {
    const api = service.api_manager.get_admin_ugroup_api();
    const resp = await api.list();
    if (!resp.ok) {
      return;
    }
    groups = resp.data;
    loading = false;
  };

  const pick_group = async (picked) => {
    loading = true;
    mode = "pick_plug";
    ugroup = picked.data.slug;

    const api = service.api_manager.get_admin_plug_api();
    const resp = await api.list_plug();
    if (!resp.ok) {
      return;
    }
    plugs = resp.data;
    loading = false;
  };

  const pick_plug = async (picked) => {
    loading = true;
    mode = "pick_agent";
    plug_id = picked.data.id;

    const api = service.api_manager.get_admin_plug_api();
    const resp = await api.list_agent(plug_id);
    if (!resp.ok) {
      return;
    }

    agents = resp.data;
    loading = false;
  };

  const pick_agent = async (picked) => {
    loading = true;
    agent_id = picked.data.id;

    service.nav.admin_target_app_new({
      target_type: TargetAppTypeUserGroupApp,
      target: ugroup,
      context_type: "app.1",
      plug_id,
      agent_id,
    });

    service.utils.small_modal_close();
  };

  load();
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else if mode === "pick_group"}
  <ActionPicker
    actions={groups.map((ug) => ({
      action: pick_group,
      icon: "hashtag",
      info: ug.name,
      name: ug.slug,
      data: ug,
    }))}
    title="Pick User Group"
  />
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
