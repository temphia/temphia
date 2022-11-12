<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../../../services";
  import type { TableService } from "../../../../services/data/table";

  const app: PortalService = getContext("__app__");

  export let table_service: TableService;
</script>

<div>Grid Layout</div>
