<script lang="ts">
  import { getContext, tick } from "svelte";
  import { params } from "svelte-hash-router";
  import type { PortalService } from "../../../services";
  import type { TableService } from "../../../services/data";
  import { LoadingSpinner } from "../../admin/core";
  import TableUI from "./_tableui.svelte";

  export let source = $params.source;
  export let group = $params.dgroup;

  const app: PortalService = getContext("__app__");
  let table_service: TableService;
  let loading = true;

  const load = async (table: string) => {
    if (!table) return;

    loading = true;
    const ds = await app.get_data_service();
    const gs = await ds.group_service(source, group);
    table_service = await gs.table_service(table);
    loading = false;
  };

  $: load($params.dtable);

  const execute_widget = (ev) => {
    console.log("@execute_widget", ev.detail);

    const widget = ev.detail || {};

    app.launcher.instance_by_target({
      target_id: String(widget.id),
      invoker_name: "data_table",
      target_name: widget.name,
      target_type: widget.target_type,
      invoker: table_service, // fixme => wrap it
    });

    tick().then(() => {
      app.launcher.plane_float();
    });
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  {#key $params.dtable && $params.layout}
    <TableUI
      layout={$params.layout}
      {table_service}
      view_modal={{
        open: app.utils.small_modal_open,
        close: app.utils.small_modal_close,
      }}
      on:on_change_to_card={(ev) =>
        app.nav.data_render_table(source, group, $params.dtable, "/card")}
      on:on_table_change={(ev) =>
        app.nav.data_render_table(source, group, ev.detail)}
      on:on_change_to_grid={() =>
        app.nav.data_render_table(source, group, $params.dtable)}
      on:admin_data_table={() =>
        app.nav.admin_data_table(source, group, $params.dtable)}
      on:goto_history={() =>
        app.nav.admin_data_activity(source, group, $params.dtable)}
      on:tb_execute_widget={execute_widget}
    />
  {/key}
{/if}
