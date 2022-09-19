<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import Layout from "../../layout.svelte";
  import Openrpc from "./openrpc.svelte";

  const app: PortalApp = getContext("__app__");
</script>

<Layout>
  <Openrpc />
</Layout>
