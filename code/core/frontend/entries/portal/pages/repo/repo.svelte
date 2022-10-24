<script lang="ts">
  import { params } from "svelte-hash-router";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import Skeleton from "./listings/_skeleton.svelte";
  import Listings from "./listings/listings.svelte";

  const app = getContext("__app__") as PortalService;
  const rapi = app.api_manager.get_repo_api();

  let sources;
  let loading = true;
  let items = [];

  (async () => {
    sources = await app.api_manager.get_repo_sources();
    const resp = await rapi.load($params.source);
    if (!resp.ok) {
      return;
    }
    items = resp.data;
    loading = false;
  })();
</script>

{#if loading}
  <Skeleton />
{:else}
  <Listings
    onChangeSource={null}
    onItemSelect={null}
    currentSource={null}
    {items}
    {sources}
  />
{/if}
