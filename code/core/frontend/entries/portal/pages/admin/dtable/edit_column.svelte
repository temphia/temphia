<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../app";
  import Layout from "../layout.svelte";
  import { DynAdminAPI } from "./dtable2";
  import Column from "./_column.svelte";

  export let source;
  export let group;
  export let table;
  export let column;

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let loading = true;
  let data;
  dynapi.get_column(source, group, table, column).then((resp) => {
    data = resp.data;
    loading = false;
  });
</script>

<Layout current_item={"dtable"} {loading}>
  <Column {data} {source} {group} {table} />
</Layout>
