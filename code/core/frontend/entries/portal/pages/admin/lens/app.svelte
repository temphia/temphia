<script lang="ts">
  import { getContext } from "svelte";
  import VirtualList from "./_virtaul_list.svelte";
  import type { PortalService } from "../core";
  import Layout from "./_layout.svelte";

  const app = getContext("__app__") as PortalService;
  const lapi = app.api_manager.get_admin_lens_api();

  let datas = [];
  let loaded = false;

  const load = async () => {
    const resp = await lapi.query_app({});
    if (!resp.ok) {
      return;
    }
    datas = resp.data;
    loaded = true;
  };
</script>

<Layout do_query={(qstr) => load()} index="app">
  {#if loaded}
    <div class="p-2 w-full h-full bg-white rounded">
      <VirtualList items={datas} let:item>
        <div class="p-1 border border-slate-50">{item}</div>
      </VirtualList>
    </div>

    <div class="font-sans flex justify-between p-1">
      <a
        href="#"
        class="flex items-center p-1 text-gray-500 bg-gray-300 rounded-md"
      >
        Previous
      </a>

      <a
        href="#"
        class="p-1 font-bold text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white"
        style="transition: all 0.2s ease;"
      >
        Next
      </a>
    </div>
  {/if}
</Layout>
