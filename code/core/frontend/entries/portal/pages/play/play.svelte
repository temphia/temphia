<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../services";
  import PlugCard from "./plug_card.svelte";
  // import View2 from "./view2/view2.svelte";

  const app: PortalService = getContext("__app__");
</script>

<div class="flex items-center justify-center h-screen">
  <PlugCard />
</div>
