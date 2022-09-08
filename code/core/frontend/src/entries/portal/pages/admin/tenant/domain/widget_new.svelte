<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";
  import { KvEditor } from "../../../../../common";

  export let did;

  let name = "";
  let plug = "";
  let agent = "";
  let exec_meta = {};
  let extra_meta = {};

  let message = "";

  const app: PortalApp = getContext("__app__");

  const saveHandle = async () => {
    const tapi = await app.get_apm().get_tenant_id();
    const resp = await tapi.add_domain_widget(did, {
      name,
      plug,
      agent,
      exec_meta,
      extra_meta,
    });

    if (resp.status !== 200) {
      message = resp.data;
      return;
    }

    app.navigator.goto_admin_org();
  };
</script>

<Layout current_item="ns">
  <div class="h-full w-full overflow-auto">
    <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
      <div
        class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
      >
        <div class="text-2xl text-indigo-900">Add Widget</div>

        <p class="text-red-500">{message}</p>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            bind:value={name}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Plug</label>
          <input
            type="text"
            bind:value={plug}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Agent</label>
          <input
            type="text"
            bind:value={agent}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Execution Meta</label
          >
          <KvEditor bind:data={exec_meta} />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
          <KvEditor bind:data={extra_meta} />
        </div>

        <div class="flex py-3">
          <button
            on:click={saveHandle}
            class="p-2 bg-blue-400 m-1 w-20 text-white rounded">Save</button
          >
        </div>
      </div>
    </div>
  </div>
</Layout>
