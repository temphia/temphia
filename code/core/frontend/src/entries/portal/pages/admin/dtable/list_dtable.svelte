<script lang="ts">
  import { AutoTable } from "../../../../../components";
  import { FloatingAdd } from "../../../../../components";
  import Layout from "../layout.svelte";

  import { getContext } from "svelte";
  import { DynAdminAPI } from "./dtable2";
  import type { PortalApp } from "../../../../../lib/app/portal";

  export let source;
  export let group;

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let tables = [];
  dynapi.load_group_tables(source, group).then((resp) => {
    tables = resp.data;
  });
</script>

<Layout current_item={"dtable"}>
  <AutoTable
    action_key="slug"
    actions={[
      {
        Class: "bg-green-400",
        Name: "Views",
        Action: (id) => {
          app.navigator.goto_views(source, group, id);
        },
      },
      {
        Name: "Hooks",
        Class: "bg-green-400",
        Action: (id) => {
          app.navigator.goto_hooks(source, group, id);
        },
      },
      {
        Name: "Data",
        Action: (id) => {
          dynapi.goto_dtabe_data(source, group, id);
        },
      },
      {
        Name: "Schema",
        Action: (id) => {
          dynapi.goto_dtable(source, group, id);
        },
      },

      {
        Name: "Edit",
        Action: (id) => {
          app.navigator.goto_dtable_edit(source, group, id);
        },
      },

      {
        Name: "Delete",
        Action: null,
        Class: "bg-red-400",
      },
    ]}
    key_names={[
      ["name", "Name"],
      ["slug", "Slug"],
      ["group_id", "Group"],
      ["description", "Description"],
    ]}
    datas={tables}
  />
</Layout>

<FloatingAdd onClick={() => {}} />
