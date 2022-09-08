<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";
  const app: PortalApp = getContext("__app__");
  export let bid: string;
  export let file: string;
</script>

<Layout>
  <div>Data Table</div>
</Layout>


