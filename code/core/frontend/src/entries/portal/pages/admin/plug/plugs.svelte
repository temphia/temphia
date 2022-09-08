<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import Layout from "../layout.svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";
  import { getContext } from "svelte";
  import Flowmap from "./flowmap/flowmap.svelte";

  const app: PortalApp = getContext("__app__");
  let plugs = [];

  const load = async () => {
    const papi = await app.get_apm().get_plug_api();
    const resp = await papi.list_plug();
    plugs = resp.data;
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
        Action: (pid) => {
          app.navigator.iframe_plug_launch(pid, "default");
        },
        icon: "lightning-bolt",
      },
      {
        Name: "Agents",
        Class: "bg-green-400",
        Action: (id) => {
          app.navigator.goto_admin_agents_page(id);
        },
        icon: "users",
      },
      {
        Name: "Edit",
        Action: (id) => {
          app.navigator.goto_admin_plug_page(id);
        },
        drop: true,
        icon: "pencil-alt",
      },

      {
        Name: "Resources",
        Action: (id) => {
          app.navigator.goto_admin_plug_resources(id);
        },
        drop: true,
        icon: "paper-clip",
      },

      {
        Name: "Flow Map",
        Action: (id) => {
          app.big_modal_open(Flowmap, { pid: id });
        },
        drop: true,
        icon: "map",
      },

      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (pid) => {
          const papi = await app.get_apm().get_plug_api();
          await papi.del_plug(pid);
          load();
        },
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
    color={["executor"]}
    datas={plugs}
    show_drop={true}
  />
</Layout>
