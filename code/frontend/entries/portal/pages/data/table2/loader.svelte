<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import type { PortalService } from "../../../services";
  import { LoadingSpinner } from "../../admin/core";

  export let source = $params.source;
  export let group = $params.dgroup;

  const app: PortalService = getContext("__app__");

  let loading = true;
  const load = async () => {
    const ds = await app.get_data_service();
    const gs = await ds.group_service(source, group);

    const table = gs.default_table();

    if (!table) {
      loading = false;
      return;
    }

    app.nav.data_render_table(source, group, table);
  };
  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div>Empty Group</div>
{/if}
