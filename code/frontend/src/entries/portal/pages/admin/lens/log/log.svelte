<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalService } from "../../core";
  import Layout from "./_layout.svelte";
  import { time_ago } from "../../../../../../lib/vendor/timeago";

  const app = getContext("__app__") as PortalService;
  const lapi = app.api_manager.get_admin_lens_api();

  let datas = [];
  let loading = false;
  let loaded = false;

  let fromDate = "";
  let toDate = "";
  let count = "20"

  let message = "";

  const load = async (filter_query: string) => {
    loading = true;
    let filters = {};

    console.log("@using_filter", filter_query);

    try {
      filters = JSON.parse(filter_query || "{}");
    } catch (error) {
      message = error;
      loading = false;
      return;
    }

    const queryOpts = {
      filters,
      count: Number(count)
    };

    if (fromDate) {
      queryOpts["from"] = new Date(fromDate).toISOString();
    }

    if (toDate) {
      queryOpts["to"] = new Date(toDate).toISOString();
    }

    const resp = await lapi.query(queryOpts);
    if (!resp.ok) {
      message = resp.data;
      loading = false;
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

<Layout {message} do_query={load} bind:fromDate bind:toDate bind:count {loading}>
  {#if loaded}
    <div class="p-2 bg-white rounded w-full overflow-auto" style="max-width: 100vw;" >
      <table class="w-full">
        <thead class="bg-white border-b">
          <tr>
            <th
              scope="col"
              class="text-sm font-medium text-gray-900 p-1 text-left"
            >
              #
            </th>
            <th
            scope="col"
            class="text-sm font-medium text-gray-900 p-1 text-left"
          >
            Time
          </th>
            <th
              scope="col"
              class="text-sm font-medium text-gray-900 p-1 text-left"
            >
              Id
            </th>
            <th
              scope="col"
              class="text-sm font-medium text-gray-900 p-1 text-left"
            >
              Label
            </th>

            <th
            scope="col"
            class="text-sm font-medium text-gray-900 p-1 text-left"
          >
            Node Id
          </th>
            <th
              scope="col"
              class="text-sm font-medium text-gray-900 p-1 text-left"
            >
              Tenant Id
            </th>

            <th
            scope="col"
            class="text-sm font-medium text-gray-900 p-1 text-left"
          >
            Index
          </th>

          <th
          scope="col"
          class="text-sm font-medium text-gray-900 p-1 text-left"
        >
          Service
        </th>

            <th
              scope="col"
              class="text-sm font-medium text-gray-900 p-1 text-left"
            >
              Message
            </th>
            

            <th
              scope="col"
              class="text-sm font-medium text-gray-900 p-1 text-left"
            >
              Data
            </th>
          </tr>
        </thead>
        <tbody>
          {#each datas as rdata, idx}
            {@const data = JSON.parse(rdata) || {}}
            {@const show_data = __open_row_idx === idx}
            {@const white =   idx %2 == 0}

            <tr class="border-b {white ? "bg-gray-100": ""}">
              <td
                class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
              >
                {idx}
              </td>

              <td
              class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
              >
              <span>{time_ago(new Date(data["time"]))}</span>

              [{data["time"] || ""}]

              </td>

              <td
                class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
              >
                {data["log_event_id"] || ""}
              </td>

              <td
                class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
              >
                {data["level"] || ""}
              </td>



              <td
              class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
            >
              {data["node_id"] || ""}
            </td>


              <td
                class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
              >
                {data["tenant_id"] || ""}
              </td>

              <td
              class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
            >
              {data["index"] || ""}
            </td>

            <td
            class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
          >
            {data["service_id"] || ""}
          </td>

              <td
                class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
              >
                {data["message"] || ""}
              </td>
              <td
                class="p-1 whitespace-nowrap text-sm font-medium text-gray-900"
              >
                <div class="p-2">
                  {#if show_data}
                    <pre class="bg-slate-100 rounded border">
                      {JSON.stringify(data, null, "\t")}
                    </pre>
                  {/if}

                  <button
                    on:click={() => {
                      if (show_data) {
                        __open_row_idx = null;
                        return;
                      }
                      __open_row_idx = idx;
                    }}
                    class="underline text-blue-500"
                    >{show_data ? "hide" : "show"}</button
                  >
                </div>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
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
