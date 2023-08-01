<script lang="ts">
  import SheetInner from "../../../../portal/pages/data/sheet/_sheet_inner.svelte";
  import { formatCells } from "../../../../portal/services/data/sheet/format";
  import type { Element } from "../../service";

  export let element: Element;
  export let data = {};

  const cells = formatCells(data["cells"]);
</script>

<div class="flex flex-col p-2 rounded border">
  <SheetInner
    {cells}
    columns={data["columns"] || []}
    rows={data["rows"] || []}
    editable={false}
    pick_label=""
  />
</div>
