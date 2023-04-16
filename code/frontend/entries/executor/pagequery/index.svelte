<script lang="ts">
  import Tailwind from "./../../xcompo/common/_tailwind.svelte";
  import DialogmodalCompo from "../../xcompo/dialogmodal/dialogmodalcompo.svelte";

  import Start from "./page/start/start.svelte";
  import Final from "./page/final/final.svelte";
  import Layout from "./page/_layout.svelte";

  import { LoadingSpinner } from "../../xcompo";
  import { PageQueryService, KEY } from "./service";
  import { setContext } from "svelte";

  export let env: any;

  const service = new PageQueryService(env);

  let mode: "START" | "END" = "START";
  let loading = false;
  let root_elem;
  let modal;
  let data = {};

  const next = (_data) => {
    data = _data;
    mode = "END";
  };

  const submit = async () => {
    return {
      ok: true,
      data: {},
    };
  };

  const load = async () => {
    loading = true;
    const resp = await service.load();
    if (!resp.ok) {
      return;
    }

    data = resp.data;
    loading = false;
  };

  setContext(KEY, {
    get_service: () => service,
    get_modal: () => modal,
    get_root: () => ({
      elem: root_elem,
      load,
    }),
  });

  load()

</script>

<DialogmodalCompo bind:modal />

<div bind:this={root_elem}>
  {#if loading}
    <LoadingSpinner />
  {:else if mode == "START"}
    <Layout>
      <Start onNext={next} onSubmit={submit} />
    </Layout>
  {:else}
    <Layout>
      <Final />
    </Layout>
  {/if}
</div>

<Tailwind />
