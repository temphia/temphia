<script lang="ts">
  import { getContext } from "svelte";
  import VirtualList from "./_virtaul_list.svelte";
  import type { PortalService } from "../../core";
  import Layout from "./_layout.svelte";
  import PrettyJson from "./_pretty_json.svelte";

  const app = getContext("__app__") as PortalService;
  const lapi = app.api_manager.get_admin_lens_api();

  let datas = [];
  let loading = false;
  let loaded = false;

  let fromDate = "";
  let toDate = "";

  const load = async () => {
    loading = true;
    const queryOpts = {
      //   from: fromDate,
      //   to: toDate,
    };

    if (fromDate) {
      queryOpts["from"] = new Date(fromDate).toISOString();
    }

    if (toDate) {
      queryOpts["to"] = new Date(toDate).toISOString();
    }

    const resp = await lapi.query(queryOpts);
    if (!resp.ok) {
      console.log("@err", resp.data);
      return;
    }

    __open_row_idx = null;

    datas = resp.data;

    loaded = true;
    loading = false;
  };

  $: __open_row_idx = null;
</script>

<Layout do_query={(qstr) => load()} bind:fromDate bind:toDate {loading}>
  {#if loaded}
    <div class="p-2 w-full h-full bg-white rounded">
      <VirtualList items={datas} let:item let:idx>
        <!-- svelte-ignore a11y-click-events-have-key-events -->
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
      <button
        class="flex items-center p-1 text-gray-500 bg-gray-300 rounded-md"
      >
        Previous
      </button>

      <button
        class="p-1 font-bold text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white"
        style="transition: all 0.2s ease;"
      >
        Next
      </button>
    </div>
  {/if}
</Layout>
