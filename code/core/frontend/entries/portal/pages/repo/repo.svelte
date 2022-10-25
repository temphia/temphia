<script lang="ts">
  import { params } from "svelte-hash-router";
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import Skeleton from "./listings/_skeleton.svelte";
  import Listings from "./listings/listings.svelte";

  const app = getContext("__app__") as PortalService;
  const rapi = app.api_manager.get_repo_api();

  let sources;
  let current_source = $params.source;
  let loading = true;
  let items = [];

  (async () => {
    sources = await app.api_manager.get_repo_sources();
    const resp = await rapi.load(current_source);
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
    onChangeSource={(next) => app.nav.repo_source(next)}
    onItemSelect={(item) => {
      app.nav.repo_item(
        current_source,
        item["group"] || item["type"],
        item["slug"]
      );
    }}
    currentSource={current_source}
    {items}
    {sources}
  />
{/if}
