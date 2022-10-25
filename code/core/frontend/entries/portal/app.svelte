<script lang="ts">
  import Tailwind from "../xcompo/common/_tailwind.svelte";
  import Portal from "./_portal.svelte";
  import build from "./services/builder";

  let loading = true;

  const app = build();
  app.init().then(() => {
    loading = false;
  });


  console.log("@EEEEEE", app)


</script>

{#if !loading}
  <Portal {app} />
{/if}

<Tailwind />
