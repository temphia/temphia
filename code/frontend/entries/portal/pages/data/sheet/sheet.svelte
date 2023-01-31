<script lang="ts">
  import { params } from "svelte-hash-router";
  import { getContext } from "svelte";
  import type { Writable } from "svelte/store";

  import { LoadingSpinner, PortalService } from "../../admin/core";
  import type { SheetService, SheetState } from "../../../services/data";

  import SheetUi from "./_sheet_ui.svelte";
  import AddColumn from "./panels/_add_column.svelte";
  import AddSheet from "./panels/_add_sheet.svelte";
  import EditRow from "./panels/_edit_row.svelte";
  import AddRow from "./panels/_add_row.svelte";

  export let source = $params.source;
  export let group = $params.dgroup;
  export let sheetid = $params.sheet;

  const app: PortalService = getContext("__app__");

  let loading = true;
  let state: Writable<SheetState>;
  let sheet_service: SheetService;

  let sheets;

  const load = async () => {
    const dsvc = await app.get_data_service();
    const gsvc = await dsvc.group_sheet(source, group);
    if (gsvc.sheets.length === 0) {
      return;
    }

    const ssvc = await gsvc.get_sheet_service(sheetid);
    sheet_service = ssvc;
    state = ssvc.state;
    sheets = sheet_service.group.sheets;

    loading = false;
  };
  load();

  const doAddSheet = () => {
    app.utils.small_modal_open(AddSheet, {});
  };

  const doAddColumn = () => {
    app.utils.small_modal_open(AddColumn, {
      onAdd: (name, ctype, opts) => {
        app.utils.small_modal_close();
      },
    });
  };

  const doEditRow = (ev) => {
    app.utils.big_modal_open(EditRow, {
      columns: $state.columns,
      cells: $state.cells,
      row: ev.detail,
    });
  };

  const doAddRow = () => {
    app.utils.big_modal_open(AddRow, { columns: $state.columns });
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <SheetUi
    active_sheet={Number(sheetid)}
    cells={$state.cells}
    columns={$state.columns}
    rows={$state.rows}
    {sheets}
    on:add_column={doAddColumn}
    on:action_goto_history={() =>
      app.nav.admin_data_activity(source, group, "sheets")}
    on:action_goto_rawtable={() =>
      app.nav.data_render_table_loader(source, group)}
    on:add_row={doAddRow}
    on:add_sheet={doAddSheet}
    on:edit_row={doEditRow}
    on:action_refresh={() => {}}
  />
{/if}
