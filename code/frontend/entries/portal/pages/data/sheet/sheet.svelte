<script>
  import { CreateStore } from "./store";
  import SheetUI from "./_sheet_ui.svelte";

  const store = CreateStore();

  $: _state = store.state;
</script>

<SheetUI
  active_sheet={1}
  cells={$_state.cells}
  columns={$_state.columns}
  rows={$_state.rows}
  sheets={$_state.sheets}
  on:add_column={(ev) => {
    const args = ev.detail;
    store.add_column(args.name, args.ctype, args.opts);
  }}
/>
