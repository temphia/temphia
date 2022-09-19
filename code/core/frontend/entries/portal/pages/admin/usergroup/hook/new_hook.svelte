<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import { KvEditor } from "../../../../../common";
  import Layout from "../../layout.svelte";

  const app: PortalApp = getContext("__app__");

  let gid = "";
  let type = "";
  let target = "";
  let data = "";
  let plug_id = "";
  let agent_id = "";
  let client_side = false;

  let getMetaData;

  const save = async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.user_group_add_hook(gid, {
      type,
      target,
      data,
      plug_id,
      agent_id,
      client_side,
      extra_meta: getMetaData(),
    });
    if (resp.status !== 200) {
      console.log("Err ", resp);
      return;
    }
    app.navigator.goto_admin_usergroups_page();
  };
</script>

<Layout current_item={"user_groups"}>
  <div class="w-full h-full p-10">
    <div class="bg-white p-2">
      <div class="text-2xl text-indigo-900">New Hook</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Type</label>
        <input
          type="text"
          bind:value={type}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Target</label>
        <input
          type="text"
          bind:value={target}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Data</label>
        <textarea
          type="text"
          bind:value={data}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Plug</label>
        <input
          type="text"
          bind:value={plug_id}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Agent</label>
        <input
          type="text"
          bind:value={agent_id}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Client Side</label>
        <input
          type="checkbox"
          bind:checked={client_side}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
        <KvEditor data={{}} bind:getData={getMetaData} />
      </div>

      <div class="flex justify-end">
        <button
          on:click={save}
          class="p-2 bg-blue-400 hover:bg-blue-600 m-1 w-20 text-white rounded"
          >Save</button
        >
      </div>
    </div>
  </div>
</Layout>
