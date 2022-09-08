<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";
  import { KvEditor } from "../../../../../common";

  export let did;
  export let wid;

  let data = {};
  let message = "";

  let get_exec_opts;
  let exec_opts_modified = false;

  let get_meta_opts;
  let meta_opts_modified = false;

  let modified = false;

  const app: PortalApp = getContext("__app__");

  const set = (name) => (ev) => {
    data = { ...data, [name]: ev.target.value };
    modified = true;
  };

  const saveHandle = async () => {
    const tapi = await app.get_apm().get_tenant_id();

    const fdata = { ...data };
    if (exec_opts_modified) {
      fdata["exec_meta"] = get_exec_opts();
    }

    if (meta_opts_modified) {
      fdata["extra_meta"] = get_meta_opts();
    }

    const resp = await tapi.update_domain_widget(did, wid, fdata);
    if (resp.data) {
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
        <div class="text-2xl text-indigo-900">Edit Widget</div>
        <p class="text-red-500">{message}</p>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Id</label>
          <input
            type="text"
            disabled
            value={data["id"]}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            value={data["name"]}
            on:change={set("name")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Plug</label>
          <input
            type="text"
            value={data["plug"]}
            on:change={set("plug")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Agent</label>
          <input
            type="text"
            value={data["agent"]}
            on:change={set("agent")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Execution Meta</label>
          <KvEditor
            data={data["exec_meta"]}
            bind:getData={get_exec_opts}
            bind:modified={exec_opts_modified}
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
          <KvEditor
            data={data["extra_meta"]}
            bind:getData={get_meta_opts}
            bind:modified={meta_opts_modified}
          />
        </div>

        <div class="flex py-3">
          {#if modified || meta_opts_modified || exec_opts_modified}
            <button
              class="p-2 bg-blue-400 m-1 w-20 text-white rounded"
              on:click={saveHandle}>Save</button
            >
          {/if}
        </div>
      </div>
    </div>
  </div>
</Layout>
