<script lang="ts">
  import type { Writable } from "svelte/store";
  import type { SheetService, SheetState } from "../../../../services/data";
  import { LoadingSpinner } from "../../../admin/core";
  import type { SheetColumn } from "../sheets";
  import SheetInner from "../_sheet_inner.svelte";

  export let current = undefined;
  export let onSelect = (val) => {};
  export let service: SheetService;
  export let column: SheetColumn;

  let loading = true;
  let state: Writable<SheetState>;

  const load = async () => {
    if (!column.refsheet) {
      console.log("invalid refsheet", column);
      return;
    }

    console.log("@column", column);

    const sservice = await service.get_sibling_sheet(column.refsheet);
    if (!sservice) {
      return;
    }
    state = sservice.state;
    loading = false;

    console.log("@state", $state);
  };

  load();
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else}
  <SheetInner
    editable={false}
    cells={$state.cells}
    columns={$state.columns}
    rows={$state.rows}
    folder_api={null}
    selected_rows={[]}
  />
{/if}
