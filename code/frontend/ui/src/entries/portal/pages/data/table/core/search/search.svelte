<script lang="ts">
  import Layout from "./_layout.svelte";
  import type { Column, TableService } from "../../../../../services/data";
  import TableRows from "./_table_rows.svelte";

  import { LoadingSpinner } from "../../../../admin/core";

  export let columns: Column[] = [];
  export let table_service: TableService;

  let count = "10";
  let column = "";
  let pattern = false;
  let search_string = "";
  let filter_name = "";

  let data;
  let loading = false;
  let message = "";

  const onSubmit = async () => {
    loading = true;
    const resp = await table_service.fts({
      search_term: search_string,
      count: Number(count),
      search_column: column,
    });
    if (resp.ok) {
      data = resp.data;
      message = "";
    } else {
      message = resp.data;
    }

    loading = false;
  };
</script>

<Layout
  {columns}
  bind:count
  bind:column
  bind:pattern
  bind:search_string
  bind:filter_name
  {onSubmit}
>
  <p class="text-red-400">{message}</p>

  {#if loading}
    <LoadingSpinner />
  {:else if data}
    <TableRows onRowSelect={() => {}} {columns} {data} />
  {/if}
</Layout>
