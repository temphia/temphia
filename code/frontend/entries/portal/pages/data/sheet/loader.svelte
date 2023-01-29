<script lang="ts">
  import { params } from "svelte-hash-router";

  import { LoadingSpinner, PortalService } from "../../admin/core";
  import { getContext } from "svelte";

  export let source = $params.source;
  export let group = $params.dgroup;

  const app: PortalService = getContext("__app__");

  let loading = true;

  const load = async () => {
    const dsvc = await app.get_data_service();
    const gsvc = await dsvc.group_sheet(source, group);

    if (gsvc.sheets.length === 0) {
      return;
    }

    app.nav.data_render_sheet(source, group, gsvc.sheets[0]["__id"]);
  };
  load();
</script>

{#if loading}
  <LoadingSpinner />
{/if}
