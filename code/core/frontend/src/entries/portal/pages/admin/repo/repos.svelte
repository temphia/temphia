<script lang="ts">
  import { AutoTable, FloatingAdd } from "../../../../../components";
  import Layout from "../layout.svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";
  import { getContext } from "svelte";

  const app: PortalApp = getContext("__app__");
  let repos = [];

  let loading = true;

  const load = async () => {
    const tapi = await app.get_apm().get_tenant_id();
    const resp = await tapi.list_repo();
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }
    repos = resp.data;
    loading = false;
  };

  load();
</script>

<Layout {loading}>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: app.navigator.goto_admin_repo_edit,
      },
      {
        Name: "Delete",
        Class: "bg-red-400",
        Action: async (rid) => {
          const tapi = await app.get_apm().get_tenant_id();
          await tapi.del_repo(rid);
          load();
        },
      },
    ]}
    key_names={[
      ["id", "ID"],
      ["name", "Name"],
      ["provider", "Provider"],
    ]}
    color={["provider"]}
    datas={repos}
  />
</Layout>

<FloatingAdd onClick={app.navigator.goto_admin_repo_new} />
