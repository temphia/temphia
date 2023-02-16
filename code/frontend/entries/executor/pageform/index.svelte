<script lang="ts">
  import StartPage from "./pages/start/start.svelte";
  import Layout from "./pages/_layout.svelte";
  import { LoadingSpinner } from "../../xcompo";
  import { LoadResponse, PageFormService } from "./service";

  export let env: any;

  let loading = true;
  let service: PageFormService;
  let data: LoadResponse;

  const load = async () => {
    service = new PageFormService(env);

    const resp = await service.load({
      data_context_type: "",
      options: {},
      rows: [],
    });

    if (!resp) {
      return;
    }

    data = resp;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <Layout>
    <StartPage {data} />
  </Layout>
{/if}
