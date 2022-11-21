<script lang="ts">
  import { getContext } from "svelte";
  import VirtualList from "./_virtaul_list.svelte";
  import type { PortalService } from "../core";
  import Layout from "./_layout.svelte";
  import PrettyJson from "./_pretty_json.svelte";

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

  $: __open_row_idx = null;
</script>

<Layout do_query={(qstr) => load()} index="app">
  {#if loaded}
    <div class="p-2 w-full h-full bg-white rounded">
      <VirtualList items={datas} let:item let:idx>
        <div
          on:click={() => {
            if (__open_row_idx === idx) {
              return;
            }
            __open_row_idx = idx;
          }}
          class="flex items-center flex-nowrap w-full border border-slate-200 p-1 gap-2 cursor-pointer"
        >
          <PrettyJson
            data={JSON.parse(item) || {}}
            index={idx}
            is_open={idx === __open_row_idx}
            toggleFunc={() => {
              if (__open_row_idx === idx) {
                __open_row_idx = null;
              } else {
                __open_row_idx = idx;
              }
            }}
          />
        </div>
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
