<script lang="ts">
  import type { Instance } from "../../services/engine/launcher";
  import Iframe from "./iframe/iframe.svelte";
  export let options: Instance;
</script>

<Iframe name={options.name} secret_id={options.id} />
