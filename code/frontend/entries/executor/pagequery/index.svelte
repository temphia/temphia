<script lang="ts">
  import Tailwind from "./../../xcompo/common/_tailwind.svelte";
  import DialogmodalCompo from "../../xcompo/dialogmodal/dialogmodalcompo.svelte";
  import type { Environment } from "../../../lib/engine/environment";

  import Start from "./page/start/start.svelte";
  import Final from "./page/final/final.svelte";
  import Layout from "./page/_layout.svelte";

  import { LoadingSpinner } from "../../xcompo";
  import { PageQueryService, KEY } from "./service";
  import { setContext } from "svelte";

  export let env: Environment;

  const service = new PageQueryService(env);
  console.log("@pagequery", service);

  let mode: "START" | "END" = "START";
  let loading = false;
  let root_elem;
  let modal;
  let data: any = {};

  let message = "";

  const submit = async (_data) => {
    const resp = await service.submit(_data);
    if (!resp.ok) {
      message = resp.data;
      return;
    }

    data = resp.data;
    mode = "END";
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

  const back = () => {
    mode = "START";
  };

  setContext(KEY, {
    get_service: () => service,
    get_modal: () => modal,
    get_root: () => ({
      elem: root_elem,
      load,
    }),
  });

  load();
</script>

<DialogmodalCompo bind:modal />

<div bind:this={root_elem} class="h-full w-full">
  {#if loading}
    <LoadingSpinner />
  {:else if mode == "START"}
    <Layout actions={{ "↻": load }}>
      <Start
        {message}
        onSubmit={submit}
        {data}
        startup_payload={env.GetExecVars().exec_data}
      />
    </Layout>
  {:else}
    <Layout
      actions={{
        "↻": load,
        "⬅️": () => {
          back();
          load();
        },
      }}
    >
      <Final {data} />
    </Layout>
  {/if}
</div>

<Tailwind />
