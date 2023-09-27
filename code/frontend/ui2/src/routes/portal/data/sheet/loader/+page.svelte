<script lang="ts">
  import { params } from "$lib/params";

  import { LoadingSpinner, PortalService } from "$lib/core";
  import { getContext } from "svelte";
  import { get } from "svelte/store";

  export let source = $params["source"];
  export let group = $params["dgroup"];

  const app: PortalService = getContext("__app__");

  let loading = true;

  const load = async () => {
    const dsvc = await app.get_data_service();
    const gsvc = await dsvc.group_sheet(source, group);

    const sheets = get(gsvc.sheets);

    if (sheets.length === 0) {
      console.log("@EMPTY sheets")
      return;
    }

    app.nav.data_sheet_render_page(source, group, String(sheets[0]["__id"]));
  };
  load();
</script>

{#if loading}
  <LoadingSpinner />
{/if}
