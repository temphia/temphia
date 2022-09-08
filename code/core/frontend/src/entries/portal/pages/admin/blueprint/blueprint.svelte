<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../lib/app/portal";

  import { Editor } from "./editor/editor";
  import BprintEditor from "./editor/editor.svelte";
  export let bid;

  let loaded = false;
  let editor;

  const app: PortalApp = getContext("__app__");

  const load = async () => {
    const api = await app.get_apm().get_bprint_api();
    editor = new Editor(bid, api);
    await editor.load();
    loaded = true;
  };

  load();
</script>

{#if loaded}
  <BprintEditor beditor={editor} />
{:else}
  <div>Loading..</div>
{/if}
