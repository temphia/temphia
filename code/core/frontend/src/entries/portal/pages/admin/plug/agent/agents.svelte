<script lang="ts">
  import { getContext } from "svelte";
  import { AutoTable, FloatingAction } from "../../../../../../components";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  import Layout from "../../layout.svelte";

  export let pid;
  let agents = [];

  const app: PortalApp = getContext("__app__");

  const load = async () => {
    const papi = await app.get_apm().get_plug_api();
    const resp = await papi.list_agent(pid);
    agents = resp.data;
  };

  load();
</script>

<Layout>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Execute",
        Class: "bg-blue-400",

        Action: (agent_id) => {
          app.navigator.iframe_plug_launch(pid, agent_id);
        },
        icon: "lightning-bolt",
      },
      {
        Name: "Edit",
        Action: async (aid) => app.navigator.goto_admin_agent_page(pid, aid),
        drop: true,
        icon: "pencil-alt",
      },
      {
        Name: "Extern Execute",
        Action: (agent_id) => {
          app.navigator.extern_plug_launch(pid, agent_id);
        },
        drop: true,
        icon: "lightning-bolt",
      },

      {
        Name: "Duplicate",
        Action: null,
        drop: true,
        icon: "duplicate",
      },

      {
        Name: "Docs",
        drop: true,
        icon: "book-open",
        Action: (aid) => {
          app.navigator.goto_admin_agent_ifacedoc_page(pid, aid);
        },
      },

      {
        Name: "Links",
        drop: true,
        icon: "link",
        Action: (aid) => {
          app.navigator.goto_admin_agents_links(pid, aid);
        },
      },

      {
        Name: "Extensions",
        drop: true,
        icon: "puzzle",
        Action: (aid) => {
          app.navigator.goto_admin_agents_exts(pid, aid);
        },
      },

      {
        Name: "Resources",
        drop: true,
        icon: "paper-clip",
        Action: (aid) => {
          app.navigator.goto_admin_agents_resources(pid, aid);
        },
      },

      {
        Name: "Dev Shell",
        drop: true,
        icon: "terminal",
        Action: (aid) => {
          app.navigator.goto_admin_agent_shell_page(pid, aid);
        },
      },

      {
        Name: "Delete",
        Action: null,
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
    datas={agents}
    show_drop={true}
    color={["type", "executor"]}
  />
</Layout>

<FloatingAction />
