<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import type { PortalService } from "../../services";
  import { LoadingSpinner } from "../admin/core";

  export let source = $params.source;
  export let group = $params.group;

  const app: PortalService = getContext("__app__");

  const load = async () => {
    const ds = await app.get_data_service();
    const gs = await ds.group_service(source, group);
    app.nav.data_table(source, group, gs.default_table());
  };

  load();
</script>

<LoadingSpinner />
