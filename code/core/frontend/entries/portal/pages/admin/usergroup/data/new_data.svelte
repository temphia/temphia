<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import { KvEditor } from "../../../../../common";
  import Layout from "../../layout.svelte";

  export let gid = "";

  let getMetaData;

  const app: PortalApp = getContext("__app__");

  let data_source = "";
  let data_group = "";
  let policy = "";

  const save = async () => {
    const uapi = await app.get_apm().get_user_api();
    const resp = await uapi.user_group_add_data(gid, {
      data_source,
      data_group,
      policy,
      user_group: gid,
      extra_meta: getMetaData(),
    });

    if (resp.status !== 200) {
      console.log("Err ", resp);
      return;
    }
  };
</script>

<Layout current_item={"user_groups"}>
  <div class="w-full h-full p-10">
    <div class="bg-white p-2">
      <div class="text-2xl text-indigo-900">User Data Groups</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Data Source</label>
        <input
          type="text"
          bind:value={data_source}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Data Group</label>
        <input
          type="text"
          bind:value={data_group}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3 relative">
        <label class="pb-2 text-gray-700 font-semibold">Policy</label>

        <textarea
          bind:value={policy}
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
