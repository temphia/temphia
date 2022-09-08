<script lang="ts">
  import { getContext } from "svelte";
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import type { PortalApp } from "../../../../../lib/app/portal";
  import Layout from "../layout.svelte";

  export let pid

  let resources = [
    // {
    //   namespace: "demo3",
    //   id: "c6vd0ttmecapm6na51g0",
    //   name: "External Ping",
    //   type: "slot",
    // },
  ];

  const app: PortalApp = getContext("__app__");

  const load = async () => {
    const rapi = await app.get_apm().get_resource_api();
    const resp = await rapi.resource_list(pid)
    resources = resp.data;
  };
  load();
</script>

<Layout current_item="resources">
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: (id) => app.navigator.goto_admin_resource_edit(id),
        icon: "pencil-alt",
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (id) => {
          const rapi = await app.get_apm().get_resource_api();
          await rapi.resource_remove(id);
          load();
        },
        icon: "trash",
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["id", "Id"],
      ["type", "Type"],
      ["schema", "Schema"],
    ]}
    color={["type"]}
    datas={resources}
  />
</Layout>

<FloatingAdd onClick={() => app.navigator.goto_admin_resource_new(pid)} />
