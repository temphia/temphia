<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../../../services";
  import type { TableService } from "../../../../services/data/table";
  import { action_builder } from "../../actions";
  import Datatable from "./datatable/datatable.svelte";

  const app: PortalService = getContext("__app__");

  export let table_service: TableService;

  const row_service = table_service.get_row_service();

  const data_store = table_service.state.data_store;
  const nav_store = table_service.state.nav_store;

  let abuilder = action_builder(table_service, app, []);
</script>

<Datatable
  actions={abuilder.actions}
  active_table={table_service.table_slug}
  all_tables={table_service.all_tables}
  columns={$data_store.column_order}
  columns_index={$data_store.indexed_column}
  hooks={[]}
  main_column=""
  loading={$nav_store.loading}
  rows={$data_store.rows}
  rows_index={$data_store.indexed_rows}
  selectedRows={[]}
/>
