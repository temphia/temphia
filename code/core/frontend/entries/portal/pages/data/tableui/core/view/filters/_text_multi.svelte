<script lang="ts">
  import { MultiText } from "../../../../../admin/core";
  export let value = "";
  export let column: object;
  let seperator = (column["extra_meta"] || {})["seperator"] || ",";
</script>

<MultiText
  onChange={(v) => {
    value = v;
  }}
  {value}
  {seperator}
/>
