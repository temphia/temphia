<script lang="ts">
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  import Portal from "./portal.svelte";
  import build from "./services/builder";

  let loading = false;

  // const app = build();
  // app.init().then(() => {
  //   loading = false;
  // });
</script>

{#if !loading}
  <Portal />
{/if}

<Tailwind />
