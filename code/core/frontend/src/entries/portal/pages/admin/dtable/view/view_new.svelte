<script lang="ts">
  import Layout from "../../layout.svelte";
  import ViewNew from "./_view_new.svelte";

  import { getContext } from "svelte";
  import { DynAdminAPI } from "../dtable2";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  export let table = "";
  export let group = "";
  export let source = "";

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let columns;
  dynapi.load_tables_column(source, group, table).then((resp) => {
    columns = resp.data;
  });
</script>

<Layout current_item={"dtable"}>
  {#if columns}
    <ViewNew {columns} {source} {group} {table} {dynapi} />
  {/if}
</Layout>
