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

  const app: PortalService = getContext("__app__");
  let table_service: TableService;
  let loading = true;

  const load = async (table: string) => {
    loading = true;
    const ds = await app.get_data_service();
    const gs = await ds.group_service(source, group);
    table_service = await gs.table_service(table);
    loading = false;
  };

  $: load($params.dtable);
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  {#key $params.dtable && $params.layout}
    {#if $params.layout === "card"}
      <CardLayout />
    {:else}
      <GridLayout
        {table_service}
        on:on_table_change={(ev) =>
          app.nav.data_table(source, group, ev.detail)}
        on:on_change_to_card={(ev) =>
          app.nav.data_table(source, group, $params.dtable, "/card")}
      />
    {/if}
  {/key}
{/if}
