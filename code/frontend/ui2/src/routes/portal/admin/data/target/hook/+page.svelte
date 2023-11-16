<script lang="ts">
  // import Hooks from "./hooks.svelte";
  import { params } from "$lib/params";
  import { getContext } from "svelte";
  import type { PortalService } from "$lib/core";

  export let source = $params["source"];
  export let group = $params["group"];
  export let table = $params["table"];

  const app = getContext("__app__") as PortalService;
</script>

<div>Data Hooks</div>

<!-- <Hooks
  action_new={() => app.nav.admin_data_hook(source, group, table)}
  ttype="data"
  target={`${source}/${group}/${table}`}
/> -->
