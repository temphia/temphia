<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import Builder from "../admin/bprint/editor/data_schema/index.svelte";

  const app: PortalService = getContext("__app__");
</script>

<Builder />
