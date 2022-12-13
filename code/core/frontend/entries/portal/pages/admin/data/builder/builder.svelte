<script lang="ts">
  import Builder from "./_builder.svelte";
  import sample from "./sample";
  import { getContext } from "svelte";
  import type { PortalService } from "../../core";

  export let bid = "";

  let loaded = false;
  let editor;

  // DataSchemaEditor

  const app: PortalService = getContext("__app__");
</script>

<Builder
  schema={sample}
  open_modal={app.utils.small_modal_open}
  close_modal={app.utils.small_modal_close}
/>
