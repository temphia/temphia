<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";
  import Openrpc from "./openrpc.svelte";

  const app: PortalApp = getContext("__app__");
</script>

<Layout>
  <Openrpc />
</Layout>
