<script lang="ts">
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  import Portal from "./_portal.svelte";
  import build from "./services";
  import { onMount } from "svelte";

  let loading = true;

  const app = build();
  app.init().then(() => {
    loading = false;
  });

  console.log("@portal_service", app);
  window["ps"] = app;

  onMount(() => {
    window.onunhandledrejection = (e) => {
      console.log("we got exception, but the app has crashed", e);
      e.preventDefault();
    };
  });
</script>

{#if !loading}
  <Portal {app} />
{/if}

<Tailwind />
