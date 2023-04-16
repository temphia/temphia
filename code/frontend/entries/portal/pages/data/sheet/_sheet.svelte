<script lang="ts">
  import { getContext, tick } from "svelte";
  import { get, Writable } from "svelte/store";
  import { LoadingSpinner, PortalService } from "../../admin/core";
  import SheetUi from "./_sheet_ui.svelte";
  import AddColumn from "./panels/column/_add_column.svelte";
  import AddSheet from "./panels/sheet/_add_sheet.svelte";
  import EditRow from "./panels/row/edit_row.svelte";
  import AddRow from "./panels/row/add_row.svelte";
  import RemoveSheetDialog from "./panels/sheet/_remove_sheet_dialog.svelte";
  import EditColumn from "./panels/column/_edit_column.svelte";
  import SearchPanel from "./panels/_search_panel.svelte";
  import type { SheetWidget } from "./sheets";
  import type { SheetService, SheetState } from "../../../services/data";
  import { TargetAppTypeDataSheetWidget } from "../../admin/target/target";

  export let source;
  export let group;
  export let sheetid;

  const app: PortalService = getContext("__app__");

  let loading = true;
  let state: Writable<SheetState>;
  let sheet_service: SheetService;
  let force_render_index;
  let folder_api;
  let sheets;

  let selected_rows = [];

  const load = async () => {
    const dsvc = await app.get_data_service();
    const gsvc = await dsvc.group_sheet(source, group);
    if (get(gsvc.sheets).length === 0) {
      return;
    }

    sheets = gsvc.sheets;

    const ssvc = await gsvc.get_sheet_service(sheetid);
    sheet_service = ssvc;
    state = ssvc.state;
    sheet_service.force_render_index;
    folder_api = sheet_service.group.folder_api;
    loading = false;
  };
  load();

  const doAddSheet = () => {
    app.utils.small_modal_open(AddSheet, {
      onAdd: async (name: string, opts: object) => {
        await sheet_service.add_sheet(name, opts);
        app.utils.small_modal_close();
      },
    });
  };

  const doRemoveSheet = () => {
    const sheet = get(sheet_service.group.sheets).filter(
      (s) => s.__id === Number(sheetid)
    )[0];

    app.utils.small_modal_open(RemoveSheetDialog, {
      name: sheet.name,
      sheet_id: sheetid,
      Confirm: async () => {
        await sheet_service.remove_sheet();
        app.utils.small_modal_close();
        app.nav.data_render_sheet_loader(source, group);
      },
      Deny: () => app.utils.small_modal_close(),
    });
  };

  const doAddColumn = () => {
    app.utils.small_modal_open(AddColumn, {
      sheets: $sheets,
      sheetid,
      service: sheet_service,
      onAdd: async (opts) => {
        await sheet_service.add_column(opts);
        app.utils.small_modal_close();
      },
    });
  };

  const doEditColumn = (ev) => {
    app.utils.small_modal_open(EditColumn, { column: ev.detail });
  };

  const doEditRow = (ev) => {
    app.utils.big_modal_open(EditRow, {
      columns: $state.columns,
      cells: $state.cells,
      service: sheet_service,
      row: ev.detail,
      folder_api,
      onSave: async (data) => {
        await sheet_service.update_row_cell(ev.detail["__id"], data);
        app.utils.big_modal_close();
        await sheet_service.init();
      },
    });
  };

  const doAddRow = () => {
    app.utils.big_modal_open(AddRow, {
      columns: $state.columns,
      service: sheet_service,
      onSave: async (data) => {
        await sheet_service.add_row_cell(data);

        app.utils.big_modal_close();
        await sheet_service.init();
      },
    });
  };

  const doSearch = () => {
    app.utils.big_modal_open(SearchPanel, { service: sheet_service });
  };

  const doActionRunWidget = (ev) => {
    const widget = ev.detail as SheetWidget;
    console.log("@widget", widget);

    app.launcher.instance_by_target({
      invoker_name: "data_sheet",
      target_id: String(widget.id),
      invoker_factory: null,
      target_name: widget.name,
      target_type: TargetAppTypeDataSheetWidget,
    });

    tick().then(() => {
      app.launcher.plane_float();
    });
  };

  const doRemoveRowId = ({ detail }) => {
    sheet_service.remove_row_cell(detail);
    selected_rows = [];
  };
</script>

{#if loading || $state.loading}
  <LoadingSpinner />
{:else}
  {#key $force_render_index}
    <SheetUi
      {folder_api}
      bind:selected_rows
      active_sheet={Number(sheetid)}
      cells={$state.cells}
      columns={$state.columns}
      rows={$state.rows}
      sheets={$sheets}
      widgets={$state.widgets}
      profile_genrator={sheet_service.profile_genrator}
      on:add_column={doAddColumn}
      on:action_goto_history={() =>
        app.nav.admin_data_activity(source, group, "sheets")}
      on:action_goto_rawtable={() =>
        app.nav.data_render_table_loader(source, group)}
      on:add_row={doAddRow}
      on:edit_row={doEditRow}
      on:action_refresh={() => sheet_service.init()}
      on:add_sheet={doAddSheet}
      on:change_sheet={(ev) => {
        app.nav.data_render_sheet(source, group, ev.detail);
      }}
      on:remove_sheet={doRemoveSheet}
      on:edit_column={doEditColumn}
      on:action_search={doSearch}
      on:action_run_widget={doActionRunWidget}
      on:action_delete_trash={doRemoveRowId}
    />
  {/key}
{/if}
