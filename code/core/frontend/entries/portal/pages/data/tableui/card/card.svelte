<script lang="ts">
  import type { TableService } from "../../../../services/data/table";
  import Renderer from "../core/renderer/renderer.svelte";
  import RendererLayout from "../core/renderer/layout.svelte";

  export let table_service: TableService;
  export let show_editor;

  export let hooks: object[];
  export let actions: object[];
  export let selected_rows = [];

  const row_service = table_service.get_row_service();

  const data_store = table_service.state.data_store;
  const dirty_store = table_service.state.dirty_store;
  const nav_store = table_service.state.nav_store;
</script>

<RendererLayout
  {actions}
  active_table={table_service.table_slug}
  all_tables={table_service.all_tables}
  loading={false}
  {selected_rows}
  {hooks}
  on:on_hook_click
  on:on_new_row
  on:on_table_change
  on:on_change_to_grid
  layout={"card"}
>
  <div class="flex w-full" style="height:calc(100vh - 6em);">
    <div>Card</div>
  </div>
</RendererLayout>
