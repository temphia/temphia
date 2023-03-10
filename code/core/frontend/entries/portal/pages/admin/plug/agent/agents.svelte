<script lang="ts">
  import { getContext } from "svelte";
  import {
    AutoTable,
    LoadingSpinner,
    FloatingAdd,
    PortalService,
  } from "../../core";
  import { params } from "svelte-hash-router";

  export let pid = $params.pid;

  let datas = [];
  let loading = true;
  const app = getContext("__app__") as PortalService;

  const load = async () => {
    const api = app.api_manager.get_admin_plug_api();
    const resp = await api.list_agent(pid);
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  // actions

  const action_execute = (id: string) =>
    app.nav.admin_plug_dev_execute(pid, id);
  const action_agent_links = (id: string) => app.nav.admin_agent_links(pid, id);
  const action_agent_exts = (id: string) => app.nav.admin_agent_ext(pid, id);
  const action_agent_resources = (id: string) => {};
  const action_dev_shell = (id: string) =>
    app.nav.admin_plug_dev_shell(pid, id);
  const action_dev_docs = (id: string) => app.nav.admin_plug_dev_docs(pid, id);
  const action_delete = async (id: string) => {
    const api = app.api_manager.get_admin_plug_api();
    const resp = await api.delete_agent(pid, id);
    if (!resp.ok) {
      return;
    }
    load();
  };

  const action_new = () => app.nav.admin_agent_new(pid);
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
        Name: "Edit",
        Action: null, //async (aid) => app.navigator.goto_admin_agent_page(pid, aid),
        drop: true,
        icon: "pencil-alt",
      },
      {
        Name: "Extern Execute",
        Action: (agent_id) => {},
        drop: true,
        icon: "lightning-bolt",
      },

      {
        Name: "Duplicate",
        Action: (id) => {},
        drop: true,
        icon: "duplicate",
      },

      {
        Name: "Docs",
        drop: true,
        icon: "book-open",
        Action: action_dev_docs,
      },

      {
        Name: "Links",
        drop: true,
        icon: "link",
        Action: action_agent_links,
      },

      {
        Name: "Extensions",
        drop: true,
        icon: "puzzle",
        Action: action_agent_exts,
      },

      {
        Name: "Resources",
        drop: true,
        icon: "paper-clip",
        Action: action_agent_resources,
      },

      {
        Name: "Dev Shell",
        drop: true,
        icon: "terminal",
        Action: action_dev_shell,
      },

      {
        Name: "Delete",
        Action: action_delete,
        drop: true,
        icon: "trash",
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["type", "Type"],
      ["executor", "Executor"],
      ["plug_id", "Plug Id"],
    ]}
    {datas}
    show_drop={true}
    color={["type", "executor"]}
  />
{/if}

<FloatingAdd onClick={action_new} />
