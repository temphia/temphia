<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import type { PortalService } from "../../services";
  import { LoadingSpinner } from "../admin/core";

  // export let source = $params.source;
  // export let group = $params.dgroup;

  const app: PortalService = getContext("__app__");
</script>

<div id="data_renderer_root">
  <LoadingSpinner />
</div>
