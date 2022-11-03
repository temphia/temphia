<script lang="ts">
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../core";
  import { params } from "svelte-hash-router";
  import BprintEditor from "./editor/editor.svelte";
  import { Editor } from "./editor/editor";

  export let bid = $params.bid;

  let loading = true;
  const app = getContext("__app__") as PortalService;

  const editor = new Editor(bid, app.api_manager.get_admin_bprint_api());

  editor.load().then(() => {
    loading = false;
  });
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <BprintEditor beditor={editor} service={app} />
{/if}
