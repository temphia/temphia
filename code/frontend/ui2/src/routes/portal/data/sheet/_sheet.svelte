<script lang="ts">
  import { getContext, tick } from "svelte";
  import { get } from "svelte/store";
  import type { Writable } from "svelte/store";
  import { LoadingSpinner, PortalService } from "$lib/core";
  import SheetUi from "./_sheet_ui.svelte";
  import AddColumn from "./panels/column/_add_column.svelte";
  import AddSheet from "./panels/sheet/_add_sheet.svelte";
  import EditRow from "./panels/row/edit/edit_row.svelte";
  import AddRow from "./panels/row/edit/add_row.svelte";
  import RemoveSheetDialog from "./panels/sheet/_remove_sheet_dialog.svelte";
  import EditColumn from "./panels/column/_edit_column.svelte";
  import SearchPanel from "./panels/_search_panel.svelte";
  import type { SheetWidget } from "./sheets";
  import type { SheetService, SheetState } from "$lib/services/data";
  import { TargetAppTypeDataSheetWidget } from "../../admin/target/target";
  import RefRecord from "./panels/row/ref_record.svelte";
  import ExtraActions from "./panels/extra_actions.svelte";

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

    let rowid = undefined;

    if (app.nav.options) {
      rowid = app.nav.options["row_cursor_id"];
      app.nav.options = {};
    }

    const ssvc = await gsvc.get_sheet_service(sheetid, rowid);
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
        app.nav.data_sheet_render_page(source, group);
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
    app.utils.small_modal_open(EditColumn, {
      column: ev.detail,
      service: sheet_service,
    });
  };

  const doEditRow = (ev) => {
    app.utils.big_modal_open(EditRow, {
      columns: $state.columns,
      cells: $state.cells,
      service: sheet_service,
      row: ev.detail,
      folder_api,
      gotoSiblingSheet: gotoSheetRow,
      onSave: async (data) => {
        const resp = await sheet_service.update_row_cell(
          ev.detail["__id"],
          data
        );
        if (resp.ok) {
          app.utils.big_modal_close();
        }
        return resp;
      },
    });
  };

  const doAddRow = (dirty_data = {}) => {
    app.utils.big_modal_open(AddRow, {
      columns: $state.columns,
      service: sheet_service,
      dirty_data,
      onSave: async (data) => {
        const resp = await sheet_service.add_row_cell(data);
        if (resp.ok) {
          app.utils.big_modal_close();
        }
        return resp;
      },
    });
  };

  const doSearch = () => {
    app.utils.big_modal_open(SearchPanel, {
      service: sheet_service,
      columns: $state.columns,
    });
  };

  const doActionRunWidget = (ev) => {
    const widget = ev.detail as SheetWidget;
    console.log("@widget", widget);

    app.launcher.instance_by_target({
      invoker_name: "data_sheet",
      target_id: String(widget.id),
      target_name: widget.name,
      target_type: TargetAppTypeDataSheetWidget,
      startup_payload: sheet_service.get_exec_data(selected_rows),
      invoker: sheet_service.get_invoker(widget, app.launcher),
    });

    tick().then(() => {
      app.launcher.plane_float();
    });
  };

  const refPreview = ({ detail }) => {
    console.log("@ref_preview", detail);

    const sheets = get(sheet_service.group.sheets);
    const sheet = sheets.filter((v) => v.__id != detail["sheetid"])[0];

    app.utils.small_modal_open(RefRecord, {
      service: sheet_service,
      onYes: () => {
        gotoSheetRow(sheet.__id, Number(detail["numval"]) - 1);
        app.utils.small_modal_close();
      },
      sheet_name: sheet.name,
      sheet_id: sheet.__id,
      row_id: detail["numval"],
      row_name: detail["value"] || "",
    });
  };

  const doRemoveRowId = ({ detail }) => {
    sheet_service.remove_row_cell(detail);
    selected_rows = [];
  };

  const extraAction = () => {
    app.utils.small_modal_open(ExtraActions, { service: sheet_service });
  };

  const clearSelects = () => {
    selected_rows = [];
  };

  const cloneRow = () => {
    const row = selected_rows[0];
    if (!row) {
      return;
    }
    const data = $state.cells[row] || {};
    delete data["__id"];
    doAddRow(data);
  };

  const gotoSheetRow = (nextsheet, rowid) => {
    app.nav.data_sheet_render_page(source, group, nextsheet, {
      row_cursor_id: rowid,
    });
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
        app.nav.admin_data_activity(source, group, "scells")}
      on:action_goto_rawtable={() =>
        app.nav.data_group_page(source, group)}
      on:add_row={() => doAddRow()}
      on:edit_row={doEditRow}
      on:action_refresh={() => sheet_service.init()}
      on:add_sheet={doAddSheet}
      on:change_sheet={(ev) => {
        app.nav.data_sheet_render_page(source, group, ev.detail);
      }}
      on:mounted={({ detail }) => {
        sheet_service.scroller = detail.scroller;
        sheet_service.close_big_modal = app.utils.big_modal_close;
        sheet_service.close_small_modal = app.utils.small_modal_close;
      }}
      on:action_extra={extraAction}
      on:ref_preview={refPreview}
      on:remove_sheet={doRemoveSheet}
      on:edit_column={doEditColumn}
      on:action_search={doSearch}
      on:action_run_widget={doActionRunWidget}
      on:action_delete_trash={doRemoveRowId}
      on:scroll_top={sheet_service.scroll_top}
      on:scroll_bottom={sheet_service.scroll_bottom}
      on:action_clear_selects={clearSelects}
      on:action_clone={cloneRow}
    />
  {/key}
{/if}
