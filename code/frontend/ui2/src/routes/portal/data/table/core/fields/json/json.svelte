<script lang="ts">
  import type { Column, RowService } from "$lib/services/data";
  import { CEditor } from "$lib/core";

  export let column: Column;
  export let value: any;
  export let row_service: RowService;
  export let onChange;

  let __value = value;

  let editor;
</script>

<CEditor
  bind:editor
  on:change={() => onChange(editor.getValue())}
  code={__value}
/>
