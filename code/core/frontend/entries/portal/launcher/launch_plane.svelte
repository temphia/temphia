<script lang="ts">
  import type { Writable } from "svelte/store";

  export let store: Writable<boolean>;

  $: console.log("#STORE", $store);
</script>

<div style="display: {$store ? 'block' : 'none'};">Launch Plane</div>
