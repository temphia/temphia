<script lang="ts">
  import {
    FloatingAdd,
    AutoTable,
    FloatingAction,
  } from "../../../../../../components";
  import Layout from "../../layout.svelte";

  import { getContext } from "svelte";
  import { DynAdminAPI } from "../dtable2";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  export let source;
  export let group;
  export let table;

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let views = [];

  dynapi.list_view(source, group, table).then((resp) => {
    views = resp.data;
  });
</script>

<Layout current_item={"dtable"}>
  <AutoTable
    action_key="id"
    actions={[
      {
        Name: "Edit",
        Action: (id) => {
          app.navigator.goto_edit_view(source, group, table, id);
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
      ["id", "Id"],
      ["table_id", "Table"],
      ["group_id", "Group"],
    ]}
    datas={views}
  />
</Layout>

<FloatingAdd
  onClick={() => {
    app.navigator.goto_add_view(source, group, table);
  }}
/>
