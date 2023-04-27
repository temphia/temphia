<script lang="ts">
  import type { SheetService, SheetState } from "../../../../services/data";
  import { LoadingSpinner } from "../../../admin/core";
  import type { SheetColumn } from "../sheets";
  import SheetInner from "../_sheet_inner.svelte";

  export let current = undefined;
  export let onSelect = (val: { __id: number; ref_value: string }) => {};
  export let service: SheetService;
  export let column: SheetColumn;

  let loading = true;
  let state;

  const load = async () => {
    if (!column.refsheet) {
      console.log("invalid refsheet", column);
      return;
    }

    loading = true;

    state = await service.ref_sheet_query({
      column_id: column.__id,
      row_cursor_id: 0,
      target_source: column.extraopts["ref_source"],
      target_group: column.extraopts["ref_group"],
      target_sheet_id: column.refsheet,
    });

    loading = false;
  };

  const pick_row = (ev) => {
    const rowid = ev.detail["__id"];
    const colcells = state.cells[rowid] || {};
    const refcell = colcells[column.refcolumn] || {};
    const ref_value = String(refcell["value"] || "");
    onSelect({ __id: rowid, ref_value });
  };

  load();
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else}
  <SheetInner
    editable={false}
    cells={state["cells"] || {}}
    columns={state["columns"] || []}
    rows={state["rows"] || []}
    folder_api={null}
    selected_rows={[]}
    on:pick_row={pick_row}
  />
{/if}
