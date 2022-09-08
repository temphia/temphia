<script lang="ts">
  import Layout from "../../layout.svelte";
  import { getContext } from "svelte";
  import { DynAdminAPI } from "../dtable2";
  import HookEdit from "./_hook_edit.svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  export let id = "";
  export let table = "";
  export let group = "";
  export let source = "";

  const app: PortalApp = getContext("__app__");
  const dynapi = new DynAdminAPI(app);

  let data;
  dynapi.get_hook(source, group, table, Number(id)).then((resp) => {
    data = resp.data;
  });
</script>

<Layout current_item={"dtable"} loading={data === null}>
  {#if data}
    <HookEdit
      {data}
      group_id={group}
      table_id={table}
      {source}
      id={Number(id)}
    />
  {/if}
</Layout>
