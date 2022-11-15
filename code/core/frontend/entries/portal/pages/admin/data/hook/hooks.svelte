<script lang="ts">
  import Hooks from "../../target/hook/hooks.svelte";
  import { params } from "svelte-hash-router";
  import { getContext } from "svelte";
  import type { PortalService } from "../../core";

  export let source = $params.source;
  export let group = $params.group;
  export let table = $params.group;

  const app = getContext("__app__") as PortalService;
</script>

<Hooks
  action_new={() => app.nav.admin_data_hook(source, group, table)}
  ttype="data"
  target={`${source}/${group}/${table}`}
/>
