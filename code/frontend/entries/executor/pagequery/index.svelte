<script lang="ts">
  import Tailwind from "./../../xcompo/common/_tailwind.svelte";

  import Start from "./page/start/start.svelte";
  import Final from "./page/final/final.svelte";
  import Layout from "./page/_layout.svelte";

  import { LoadingSpinner } from "../../xcompo";

  export let env: any;

  let mode: "START" | "END" = "START";
  let loading = false;

  let data = {};
</script>

{#if loading}
  <LoadingSpinner />
{:else if mode == "START"}
  <Layout>
    <Start
      onNext={(_data) => {
        data = _data;
        mode = "END";
      }}
      onSubmit={async () => {
        return {
          ok: true,
          data: {},
        };
      }}
    />
  </Layout>
{:else}
  <Layout>
    <Final />
  </Layout>
{/if}

<Tailwind />
