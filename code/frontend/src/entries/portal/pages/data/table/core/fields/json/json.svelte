<script lang="ts">
  import type { Column, RowService } from "../../../../../../services/data";
  import { CEditor } from "../../../../../admin/core";

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
