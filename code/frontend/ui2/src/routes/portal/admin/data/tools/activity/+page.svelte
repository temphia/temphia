<script lang="ts">
  import { getContext } from "svelte";
  import { AutoTable, LoadingSpinner, PortalService } from "$lib/core";
  import { params } from "$lib/params";

  export let source = $params["source"];
  export let group = $params["group"];
  export let table = $params["table"];

  const app = getContext("__app__") as PortalService;
  const api = app.api_manager.get_admin_data_api();

  let message = "";
  let datas = [];
  let loading = true;
  let offset_history = [];
  let offset = 0;

  const load = async () => {
    loading = true;
    const resp = await api.list_table_activity(source, group, table, offset);
    if (!resp.ok) {
      message = resp.data;
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
  <AutoTable
    action_key="id"
    show_drop={true}
    actions={[
      {
        Class: "bg-green-400",
        Name: "preview",
        Action: () => {},
        drop: true,
      },

      {
        Class: "bg-blue-400",
        Name: "Follow User",
        Action: () => {},
        drop: true,
      },

      {
        Class: "bg-yellow-400",
        Name: "Follow Row",
        Action: () => {},
        drop: true,
      },
    ]}
    key_names={[
      ["id", "Id"],
      ["type", "Type"],
      ["row_id", "Row Id"],
      ["row_version", "Row Version"],
      ["init_sign", "Init Sign"],
      ["user_id", "User Id"],
      ["user_sign", "User Sign"],
      ["payload", "Payload"],
      ["message", "Message"],
      ["created_at", "Created At"],
    ]}
    color={["type", "user_id"]}
    {datas}
  />

  <div class="flex justify-between p-1">
    <button
      class="flex items-center p-1 text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white font-bold"
      on:click={() => {
        if (offset_history.length > 0) {
          offset = offset_history.pop();
        } else {
          offset = 0;
        }
        load();
      }}>Previous</button
    >
    <button
      on:click={() => {
        if (datas.length > 0) {
          offset_history.push(offset);
          offset = (datas[datas.length - 1] || {}).id || 0;
        } else {
          offset_history = [];
          offset = 0;
        }

        load();
      }}
      class="p-1 text-gray-500 bg-gray-300 rounded-md hover:bg-teal-400 hover:text-white font-bold"
      style="transition: all 0.2s ease 0s;">Next</button
    >
  </div>
{/if}
