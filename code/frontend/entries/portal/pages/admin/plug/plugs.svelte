<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../core";

  let datas = [];
  let loading = true;
  const app = getContext("__app__") as PortalService;

  const api = app.api_manager.get_admin_plug_api();

  const load = async () => {
    const resp = await api.list_plug();
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_execute = (id: string) =>
    app.nav.admin_plug_dev_execute(id, "default");

  const action_list_agents = (id: string) => app.nav.admin_agents(id);
  const action_edit = (id: string) => app.nav.admin_plug_edit(id);

  const action_list_states = (id: string) => app.nav.admin_plug_states(id);

  const action_show_flowmap = (id: string) =>
    app.nav.admin_plug_dev_flowmap(id);
  const action_delete = async (id: string) => {
    const papi = app.api_manager.get_admin_plug_api();
    await papi.delete_plug(id);
    load();
  };

  const action_reset = async (id) => {
    loading = true;
    const eapi = app.api_manager.get_engine_api();

    const aresp = await api.list_agent(id);
    if (!aresp.ok) {
      loading = false;
      return;
    }

    await Promise.all(
      aresp.data.map(async (data) => {
        await eapi.reset(id, data["id"]);
      })
    ).finally(() => {
      loading = false;
    });
  };

  const action_logs = (id: string) => {
    app.nav.admin_lens_logs({
      filters: {
        plug_id: id,
      },
    });
  };

  const action_new = () => app.nav.admin_plug_new();

  const action_goto_files = (id, data) =>
    app.nav.admin_bprint_files(data["bprint_id"]);
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Execute",
        Class: "bg-blue-400",
        Action: action_execute,
        icon: "lightning-bolt",
      },
      {
        Name: "Agents",
        Class: "bg-green-400",
        Action: action_list_agents,
        icon: "users",
      },
      {
        Name: "Edit",
        Action: action_edit,
        drop: true,
        icon: "pencil-alt",
      },

      {
        Name: "States",
        Action: action_list_states,
        drop: true,
        icon: "database",
      },

      {
        Name: "Flow Map",
        Action: action_show_flowmap,
        drop: true,
        icon: "map",
      },

      {
        Name: "Logs",
        drop: true,
        icon: "document-search",
        Action: action_logs,
      },
      {
        Name: "Files",
        Action: action_goto_files,
        Class: "bg-green-400",
        icon: "document-duplicate",
        drop: true,
      },

      {
        Name: "Reset",
        Action: action_reset,
        drop: true,
        icon: "refresh",
      },

      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: action_delete,
        icon: "trash",
        drop: true,
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["owner", "Owner"],
      ["bprint_id", "Bprint Id"],
    ]}
    color={[]}
    {datas}
    show_drop={true}
  />
{/if}

<FloatingAdd onClick={action_new} />
