<script lang="ts">
  import Layout from "../layout.svelte";
  import Dtable from "./_dtable.svelte";

  import { getContext } from "svelte";
  import { DynAdminAPI } from "./dtable2";
  import type { PortalApp } from "../../../../../lib/app/portal";

  export let source;
  export let group;
  export let table;

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let loading = true;
  let data;

  dynapi.get_dtable(source, group, table).then((resp) => {
    data = resp.data;
    loading = false;
  });
</script>

<Layout current_item={"dtable"} {loading}>
  <Dtable {data} {group} {source} />
</Layout>
