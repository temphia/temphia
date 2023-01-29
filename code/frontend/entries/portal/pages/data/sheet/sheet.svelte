<script lang="ts">
  import { params } from "svelte-hash-router";

  import { LoadingSpinner, PortalService } from "../../admin/core";
  import { getContext } from "svelte";
  import SheetUi from "./_sheet_ui.svelte";
  import type { SheetService, SheetState } from "../../../services/data/sheet";
  import type { Writable } from "svelte/store";

  export let source = $params.source;
  export let group = $params.dgroup;
  export let sheetid = $params.sheet;

  const app: PortalService = getContext("__app__");

  let loading = true;
  let state: Writable<SheetState>;
  let sheet_service: SheetService;

  let sheets;
  let rows = []

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
    on:add_column={(ev) => {}}
  />
{/if}
