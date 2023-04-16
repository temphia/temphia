<script lang="ts">
  import Tailwind from "./../../xcompo/common/_tailwind.svelte";

  import Start from "./page/start/start.svelte";
  import Final from "./page/final/final.svelte";
  import Layout from "./page/_layout.svelte";

  import { LoadingSpinner } from "../../xcompo";

  export let env: any;

  let mode: "START" | "END" = "START";
  let loading = false;
</script>

{#if loading}
  <LoadingSpinner />
{:else if mode == "START"}
  <Layout>
    <Start />
  </Layout>
{:else}
  <Layout>
    <Final />
  </Layout>
{/if}

<Tailwind />
