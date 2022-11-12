<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import type { PortalService } from "../../services";
  import type { TableService } from "../../services/data/table";
  import { LoadingSpinner } from "../admin/core";
  import CardLayout from "./layout/card/card.svelte";
  import GridLayout from "./layout/grid/grid.svelte";

  export let source = $params.source;
  export let group = $params.dgroup;
  export let table = $params.dtable;
  export let layout = $params.layout;

  const app: PortalService = getContext("__app__");
  let table_service: TableService;
  let loading = true;

  const load = async () => {
    const ds = await app.get_data_service();
    const gs = await ds.group_service(source, group);
    table_service = await gs.table_service(table);
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else if layout === "card"}
  <CardLayout />
{:else}
  <GridLayout {table_service} />
{/if}
