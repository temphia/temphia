<script lang="ts">
  import { getContext } from "svelte";
  import { params } from "svelte-hash-router";
  import { FloatingAdd, LoadingSpinner, PortalService } from "../../core";

  export let pid = $params.pid;

  let datas = [];
  let loading = true;
  let page = 0;
  let key_cursor = "";
  let cursor_history = [];

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_plug_api();

  const load = async () => {
    loading = true;
    const resp = await api.list_plug_state(pid, {
      key_cursor,
      page,
      no_value: true,
    });
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    datas = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="overflow-auto p-4">
    <table class="min-w-full shadow rounded-lg">
      <thead class="bg-gray-50 border-b">
        <tr>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Key
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Version
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Tag1
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Tag2
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Tag3
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            TTL
          </th>

          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Actions
          </th>
        </tr>
      </thead>
      <tbody>
        {#each datas as data}
          <tr
            class="bg-white border-b transition duration-300 ease-in-out hover:bg-gray-100"
          >
            <td
              class="px-2 py-2 whitespace-nowrap text-sm font-medium text-gray-900"
              >{data["key"] || ""}</td
            >
            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["version"] || ""}
            </td>
            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["tag1"] || ""}
            </td>
            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["tag2"] || ""}
            </td>

            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["tag3"] || ""}
            </td>

            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["ttl"] || ""}
            </td>

            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap flex gap-2"
            >
              <button
                class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 bg-blue-400"
                on:click={() => app.nav.admin_plug_state_edit(pid, data["key"])}
              >
                Edit
              </button>

              <button
                on:click={async () => {
                  const resp = await api.delete_plug_state(pid, data["key"]);
                  load();
                }}
                class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 bg-red-400"
              >
                Delete
              </button>
            </td>
          </tr>
        {/each}
      </tbody>
    </table>
    <div class="flex justify-between p-1">
      <button
        class="flex items-center p-1 text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white font-bold"
        on:click={() => {
          if (cursor_history.length > 0) {
            key_cursor = cursor_history.pop();
          } else {
            key_cursor = "";
          }
          load();
        }}>Previous</button
      >
      <button
        on:click={() => {
          if (datas.length > 0) {
            cursor_history.push(key_cursor);
            key_cursor = datas[datas.length - 1].key || "";
          } else {
            key_cursor = "";
            cursor_history = [];
          }
          load();
        }}
        class="p-1 text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white font-bold"
        style="transition: all 0.2s ease 0s;">Next</button
      >
    </div>
  </div>
{/if}

<FloatingAdd onClick={() => app.nav.admin_plug_state_new(pid)} />
