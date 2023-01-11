<script lang="ts">
  import Builder from "./_builder.svelte";
  import sample from "./sample/sample";
  import { getContext } from "svelte";
  import type { PortalService } from "../../../core";

  import * as b from "./builder";

  let loaded = false;
  let editor;

  let state = sample as b.State;

  const app: PortalService = getContext("__app__");
</script>

<Builder
  builder={b.Builder.from_batch(state)}
  open_modal={(comp, opts) => app.utils.small_modal_open(comp, opts)}
  close_modal={() => app.utils.small_modal_close()}
/>
