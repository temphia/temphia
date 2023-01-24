<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import type { PortalService } from "../../services";
  import { LoadingSpinner } from "../admin/core";

  export let source = $params.source;
  export let group = $params.dgroup;
  export let rtype = $params.rtype;

  const app: PortalService = getContext("__app__");

  let loading = true;
  let empty = false;

  const load = async () => {
    const ds = await app.get_data_service();
    const gs = await ds.group_service(source, group);

    const table = gs.default_table();
    if (rtype && rtype !== "default") {
      const render = app.registry.Get("temphia.data_renderer", rtype);
      render({
        app,
        source,
        group,
        rtype,
      });
      return;
    }

    if (!table) {
      loading = false;
      empty = true;
      return;
    }

    app.nav.data_table(source, group, table);
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else if empty}
  <div>Empty Group</div>
{/if}
