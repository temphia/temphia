<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { Column, RowService } from "../../../../../../services/data";

  import { KeyPrimary } from "../field";
  import RefPanel from "./ref_panel.svelte";

  export let value;
  export let column: Column;
  export let onChange: (value: any) => void;
  export let row_service: RowService;

  const loader = (cursor: number) => {
    return row_service.ref_load({
      column: column.slug,
      type: column.ref_type,
      target: column.ref_target,
      object: column.ref_object,
      cursor_row_id: cursor,
    });
  };

  const openPanel = () => {
    row_service.open_model(RefPanel, {
      loader,
      onRowSelect: (row: object) => {
        console.log("ROW", row);
        onChange(row[KeyPrimary]);

        // fixme => ref_copy here
        row_service.close_model();
      },
    });
  };

  $: _cached_row_id = "";
</script>

<div class="flex w-full">
  <div
    class="p-2 w-full rounded-lg bg-gray-100 outline-none focus:bg-gray-200 mr-1"
  >
    <div class="inline-flex bg-yellow-50 text-gray-600 rounded px-1">
      <span>{_cached_row_id}</span>
      <span class="font-semibold text-gray-800 ml-1 mt-1">{value}</span>
    </div>
  </div>

  <button on:click={openPanel}>
    <Icon name="link" class="h-6 w-6" />
  </button>
</div>
