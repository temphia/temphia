<script lang="ts">
  import Builder from "./_builder.svelte";
  import sample from "./sample";
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  export let bid = "";

  let loaded = false;
  let editor;

  const app: PortalApp = getContext("__app__");
</script>

<Builder schema={sample} {app} />
