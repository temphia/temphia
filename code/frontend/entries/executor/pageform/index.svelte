<script lang="ts">
  import StartPage from "./pages/start/start.svelte";
  import Layout from "./pages/_layout.svelte";
  import { LoadingSpinner } from "../../xcompo";
  import {  PageFormService } from "./service";

  export let env: any;

  let loading = true;
  let service: PageFormService;
  let data: any;

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

  const submit = async (ev) => {
    console.log(ev.detail);
    loading = true;
    const resp = await service.submit(data.stage, ev.detail);
    if (!resp) {
      return;
    }
    data = resp;
    loading = false;
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <Layout>
    <StartPage {data} on:submit={submit} />
  </Layout>
{/if}
