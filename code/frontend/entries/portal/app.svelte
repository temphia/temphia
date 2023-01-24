<script lang="ts">
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  import Portal from "./_portal.svelte";
  import build from "./services";
  import { onMount } from "svelte";
  import Sheets from "./pages/data/sheets/sheets.svelte";

  let loading = true;

  const app = build();
  app.init().then(() => {
    app.registry.RegisterFactory("temphia.data_renderer", "sheet", (opts) => {
      new Sheets({
        target: opts["target"],
        props: opts["props"] || {},
      });
    });

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
