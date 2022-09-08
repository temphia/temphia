<script lang="ts">
  import { getContext } from "svelte";

  import type { PortalApp } from "../../../../../../lib/app/portal";

  import Layout from "../../layout.svelte";
  import { DynAdminAPI } from "../dtable2";
  import ViewEdit from "./_view_edit.svelte";

  export let table = "";
  export let group = "";
  export let source = "";
  export let id = 0;

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let columns;
  dynapi.load_tables_column(source, group, table).then((resp) => {
    columns = resp.data;
  });

  let loaded = false;
  let data = {};

  const load = async () => {
    const resp = await dynapi.get_view(source, group, table, id);
    data = resp.data;
    loaded = true;
  };

  load();
</script>

<Layout current_item={"dtable"} loading={!loaded}>
  {#if loaded}
    <ViewEdit {source} {group} {table} {columns} {id} {data} {dynapi} />
  {/if}
</Layout>
