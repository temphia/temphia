<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import Layout from "../../layout.svelte";
  import { AutoTable, FloatingAdd } from "../../../../../../shared";

  export let aid: string;
  export let pid: string;
  let exts = [];

  const app: PortalApp = getContext("__app__");

  const load = async () => {
    const papi = await app.get_apm().get_plug_api();
    const resp = await papi.agent_extension_list(pid, aid);
    if (resp.status !== 200) {
      console.log("@err", resp);
      return;
    }
    exts = resp.data;
  };

  load();
</script>

<Layout>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: (id) => {
          app.navigator.goto_admin_agents_ext_edit(pid, aid, id);
        },
        icon: "trash",
      },
      {
        Name: "Delete",
        Action: async (id) => {
          const papi = await app.get_apm().get_plug_api();
          await papi.agent_extension_del(pid, aid, id);
          load();
        },
        drop: true,
        icon: "trash",
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["plug_id", "Plug"],
      ["agent_id", "Agent"],
      ["bprint_id", "Bprint"],
      ["ref_file", "File"],
    ]}
    datas={exts}
  />
</Layout>

<FloatingAdd
  onClick={() => app.navigator.goto_admin_agents_ext_new(pid, aid)}
/>
