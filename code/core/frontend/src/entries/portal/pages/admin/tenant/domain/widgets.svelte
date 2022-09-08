<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../../components";
  import Layout from "../../layout.svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import { getContext } from "svelte";

  export let did = 0;

  const app: PortalApp = getContext("__app__");
  let widgets = [];

  const load = async () => {
    const tapi = await app.get_apm().get_tenant_id();
    const resp = await tapi.list_domain_widget(did);
    widgets = resp.data;
  };

  load();
</script>

<Layout current_item="ns">
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: (id) => {},
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (pid) => {},
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
    ]}
    datas={widgets}
  />
</Layout>

<FloatingAdd onClick={() => app.navigator.goto_admin_widget_new(did)} />
