<script lang="ts">
  import { LoadingSpinner } from "../../../../../adapter_editor/easypage/core";
  import type { PortalService } from "../../core";
  import ActionPicker from "../../core/action_picker.svelte";
  import {
    TargetAppTypeDataTableWidget,
  } from "../target";

  export let service: PortalService;

  let loading = true;
  let source = "";
  let group = "";
  let table = "";
  let plug_id = "";
  let agent_id = "";

  let sources = [];
  let groups = [];
  let tables = [];
  let plugs = [];
  let agents = [];

  let mode:
    | "pick_source"
    | "pick_group"
    | "pick_table"
    | "pick_plug"
    | "pick_agent" = "pick_source";

  const load = async () => {
    sources = await service.api_manager.self_data.get_data_sources();
    loading = false;
  };

  const pick_source = async (picked) => {
    loading = true;
    mode = "pick_group";
    source = picked.name;

    const api = await service.api_manager.get_admin_data_api();
    const resp = await api.list_group(source);
    if (!resp.ok) {
      return;
    }
    groups = resp.data;
    loading = false;
  };

  const pick_group = async (picked) => {
    loading = true;
    mode = "pick_table";
    group = picked.data.slug;

    const api = service.api_manager.get_admin_data_api();
    const resp = await api.list_tables(source, group);
    if (!resp.ok) {
      return;
    }
    tables = resp.data;
    loading = false;
  };

  const pick_table = async (picked) => {
    loading = true;
    mode = "pick_plug";
    table = picked.data.slug;

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
      target_type: TargetAppTypeDataTableWidget,
      target: `${source}/${group}/${table}`,
      context_type: "global.1",
      plug_id,
      agent_id,
    });

    service.utils.small_modal_close();
  };

  load();
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else if mode === "pick_source"}
  <ActionPicker
    actions={sources.map((ug) => ({
      action: pick_source,
      icon: "hashtag",
      info: "",
      name: ug,
    }))}
    title="Pick Data source"
  />
{:else if mode === "pick_group"}
  <ActionPicker
    actions={groups.map((ug) => ({
      action: pick_group,
      icon: "hashtag",
      info: ug.description,
      name: ug.slug,
      data: ug,
    }))}
    title="Pick Data Group"
  />
{:else if mode === "pick_table"}
  <ActionPicker
    actions={tables.map((ug) => ({
      action: pick_table,
      icon: "hashtag",
      info: ug.name,
      name: ug.slug,
      data: ug,
    }))}
    title="Pick Data Table"
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
