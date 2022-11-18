<script lang="ts">
  import type { Instance } from "../../services/engine/launcher";
  import Iframe from "./iframe/iframe.svelte";
  export let options: Instance;
</script>

<Iframe {options} />
