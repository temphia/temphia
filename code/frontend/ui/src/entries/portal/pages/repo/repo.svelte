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

  const load = async (__source: string) => {
    loading = true;
    sources = await app.api_manager.self_data.get_repo_sources();
    const resp = await rapi.load(__source);
    if (!resp.ok) {
      return;
    }
    items = resp.data;
    loading = false;
  };

  $: load($params.source);
</script>

{#if loading}
  <Skeleton />
{:else if $params.source}
  <Listings
    onChangeSource={(next) => app.nav.repo_source(next)}
    onItemSelect={(item) => {
      app.nav.repo_item(
        $params.source,
        item["group"] || item["type"],
        item["slug"]
      );
    }}
    currentSource={$params.source}
    {items}
    {sources}
  />
{/if}
