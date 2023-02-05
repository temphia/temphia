<script lang="ts">
  import type { Writable } from "svelte/store";
  import type { SheetService, SheetState } from "../../../../services/data";
  import { LoadingSpinner } from "../../../admin/core";
  import type { SheetColumn } from "../sheets";

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

    const sservice = await service.get_sibling_sheet(column.refsheet);
    if (!sservice) {
      return;
    }
    state = service.state;
    loading = false;

    console.log("@state", $state);
  };

  load();
</script>

{#if loading}
  <LoadingSpinner classes="" />
{:else}
  <div>Ref</div>
{/if}
