<script lang="ts">
  import { params } from "svelte-hash-router";
  import ListingItem from "./listings/item.svelte";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";

  import Importer from "./importer/importer.svelte";
  import { LoadingSpinner } from "../admin/core";

  const app = getContext("__app__") as PortalService;

  let source = $params.source;
  let group = $params.group;
  let item = $params._;

  let data;

  (async () => {
    const rapi = app.api_manager.get_repo_api();
    const resp = await rapi.getBprint(source, group, item);
    if (!resp.ok) {
      return;
    }
    data = resp.data;
  })();
</script>

{#if data}
  <ListingItem
    {data}
    importFunc={async (item) => {
      app.utils.small_modal_open(Importer, { data, group, source, app });
    }}
  />
{:else}
  <LoadingSpinner />
{/if}
