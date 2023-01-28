<script lang="ts">
  import Sheet from "./sheet.svelte";
  import { params } from "svelte-hash-router";

  import type { LoadingSpinner, PortalService } from "../../admin/core";
  import { getContext } from "svelte";

  export let source = $params.source;
  export let group = $params.dgroup;

  const app: PortalService = getContext("__app__");

  // const load = async () => {    
  // };
  // load();
</script>

<Sheet />
