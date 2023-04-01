<script lang="ts">
  import { params } from "svelte-hash-router";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";

  import Importer from "./importerv2/index.svelte";
  import { LoadingSpinner } from "../admin/core";

  const app = getContext("__app__") as PortalService;

  let group = $params.group;
  let item = $params._;

  let data;

  (async () => {
    const rapi = app.api_manager.get_repo_api();
    const resp = await rapi.getBprint($params.source, group, item);
    if (!resp.ok) {
      return;
    }
    data = resp.data;
  })();
</script>

{#if data}
  <Importer {data} source={$params.source} />
{:else}
  <LoadingSpinner />
{/if}
