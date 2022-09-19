<script lang="ts">
  import { getContext } from "svelte";
  import { onMount } from "svelte";
  import { StoreItem } from "../../../../shared";
  import type { PortalApp } from "../../app";
  import ImporterV2 from "../admin/blueprint/importer/importer_v2.svelte";

  export let item;
  export let source;
  export let group;

  const app: PortalApp = getContext("__app__");

  let data = null;
  let repo_api = null;

  onMount(async () => {
    repo_api = await app.get_apm().get_bprint_api();
    let resp = await repo_api.repo_get(source, group, item);
    data = resp.data;
  });
</script>

{#if data}
  <StoreItem
    {data}
    importFunc={async () => {
      app.simple_modal_open(ImporterV2, { data, group, source, app });
    }}
  />
{/if}
