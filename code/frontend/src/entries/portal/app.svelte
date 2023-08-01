<script lang="ts">
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  import Portal from "./_portal.svelte";
  import build from "./services";
  import { onMount } from "svelte";
  import { LoadingSpinner } from "../xcompo";
  import Noauth from "./pages/noauth.svelte";

  let loading = true;
  let message = "";

  const app = build();

  if (app) {
    console.log("@portal_service", app);
    window["ps"] = app;

    app.init().then((msg) => {
      message = msg;
      loading = false;
    });
  } else {
    message = "Not Logged";
    loading = false;
  }

  onMount(() => {
    window.onunhandledrejection = (e) => {
      console.log("we got exception, but the app has crashed", e);
      e.preventDefault();
    };
  });
</script>

{#if loading}
  <LoadingSpinner />
{:else if message}
  <Noauth />
{:else}
  <Portal {app} />
{/if}

<Tailwind />
