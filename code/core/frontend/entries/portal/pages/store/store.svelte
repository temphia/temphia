<script lang="ts">
  import { getContext } from "svelte";
  import { StoreListings, StoreItem } from "../../../../shared";
  import Storelist from "../../../../shared/skeleton/storelist.svelte";
  import type { PortalApp } from "../../app";

  const app: PortalApp = getContext("__app__");

  let sources = [];
  let current;
  let data = [];
  app.get_store_sources().then((_sources) => {
    sources = Object.keys(_sources);
    load(sources[0]);
  });

  const load = async (nextsrc: string) => {
    const api = await app.get_apm().get_bprint_api();
    const resp = await api.repo_list(nextsrc);
    data = resp.data;
    current = nextsrc;
  };

  export let itemSelect;
</script>

{#key current}
  {#if current}
    <StoreListings
      currentSource={current}
      items={data}
      onChangeSource={load}
      onItemSelect={(item) => {
        if (itemSelect) {
          itemSelect(item);
        } else {
          app.navigator.goto_repo_item(
            current,
            item["group"] || item["type"],
            item["slug"]
          );
        }
      }}
      {sources}
    />
  {:else}
    <Storelist />
  {/if}
{/key}
