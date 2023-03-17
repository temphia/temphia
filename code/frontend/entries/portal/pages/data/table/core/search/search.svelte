<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import Layout from "./_layout.svelte";
  import type { Column, TableService } from "../../../../../services/data";

  export let columns: Column[] = [];
  export let table_service: TableService;

  let count = "10";
  let use_filter = false;
  let column = "";
  let pattern = false;
  let search_string = "";
</script>

<Layout {columns} {count} {use_filter} {column} {pattern} {search_string}>
  <div />
</Layout>
