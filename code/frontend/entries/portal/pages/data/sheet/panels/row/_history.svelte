<script lang="ts">
  import type { SheetService } from "../../../../../services/data";

  export let service: SheetService;
  export let rid;

  

  console.log(service.api.get_row_history(service.sheetid, rid))


</script>

<div>History</div>
