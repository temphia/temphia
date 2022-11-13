<script lang="ts">
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  import Portal from "./_portal.svelte";
  import build from "./services";

  let loading = true;

  const app = build();
  app.init().then(() => {
    loading = false;
  });

  console.log("@portal_service", app);
  window["ps"] = app
</script>

{#if !loading}
  <Portal {app} />
{/if}

<Tailwind />
