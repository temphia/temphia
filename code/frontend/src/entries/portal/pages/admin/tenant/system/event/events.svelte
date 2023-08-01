<script lang="ts">
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../../core";

  let datas = [];
  let loading = true;
  let last = 0;
  let etype = "";
  let cursor_history = [];

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_tenant_api();

  const load = async () => {
    loading = true;
    const resp = await api.list_system_event({
      etype,
      last,
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
  <div class="overflow-auto p-4 ">
    <table class="min-w-full shadow rounded-lg">
      <thead class="bg-gray-50 border-b">
        <tr>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Id
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Type
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Data
          </th>
          <th
            scope="col"
            class="text-sm font-medium text-gray-900 px-2 py-2 text-left"
          >
            Tenant
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
              >{data["id"] || ""}</td
            >
            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["type"] || ""}
            </td>
            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["data"] || ""}
            </td>
            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap"
            >
              {data["tenant"] || ""}
            </td>

            <td
              class="text-sm text-gray-900 font-light px-2 py-2 whitespace-nowrap flex gap-2"
            >
              <button
                class="p-1 text-white text-sm font-semibold flex self-center shadow rounded hover:scale-110 bg-blue-400"
                on:click={() => {}}
              >
                Edit
              </button>

              <button
                on:click={async () => {}}
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
            last = cursor_history.pop();
          } else {
            last = 0;
          }
          load();
        }}>Previous</button
      >
      <button
        on:click={() => {
          if (datas.length > 0) {
            cursor_history.push(last);
            last = datas[datas.length - 1].id || "";
          } else {
            last = 0;
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
