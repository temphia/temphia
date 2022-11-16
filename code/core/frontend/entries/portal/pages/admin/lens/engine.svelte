<script lang="ts">
  import { getContext } from "svelte";
  import VirtualList from "./_virtaul_list.svelte";
  import type { PortalService } from "../core";
  import Layout from "./_layout.svelte";

  const app = getContext("__app__") as PortalService;
  const lapi = app.api_manager.get_admin_lens_api();

  let datas = [];

  const load = async () => {
    const resp = await lapi.query_engine({});
    if (!resp.ok) {
      return;
    }
    datas = resp.data;
  };
</script>

<Layout do_query={(qstr) => load()} index="engine">
  {#if datas && datas.length > 0}
    <div class="p-2 w-full h-full bg-white rounded">
      <VirtualList items={datas} let:item>
        <div class="p-1 border border-slate-50">{item}</div>
      </VirtualList>
    </div>
  {/if}
</Layout>
