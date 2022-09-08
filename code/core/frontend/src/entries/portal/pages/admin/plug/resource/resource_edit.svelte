<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";
  import { PrimaryButton } from "../../../../../common";

  const app: PortalApp = getContext("__app__");

  export let pid: string;
  export let aid: string;
  export let rid: string;

  let data = {};

  let loading = true;
  const load = async () => {
    const papi = await app.get_apm().get_plug_api();
    const resp = await papi.agent_resource_get(pid, aid, rid);
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }
    data = resp.data;
    loading = false;
  };

  let mod_data = {};
  let modified = false;

  const get = (name) => data[name] || "";
  const set = (name) => (ev) => {
    mod_data[name] = ev.target.value;
    modified = true;
  };

  const save = async () => {
    const papi = await app.get_apm().get_plug_api();

    const resp = await papi.agent_resource_update(pid, aid, rid, mod_data);
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }

    app.navigator.goto_admin_agents_page(pid);
  };

  load();
</script>

<Layout {loading}>
  <div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
    <div class="p-5 bg-white w-full ">
      <div class="text-2xl text-indigo-900">Edit Agent Resource</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          value={get("name")}
          on:change={set("name")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Plug</label>
        <input
          type="text"
          value={pid}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Agent</label>
        <input
          type="text"
          value={aid}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Resource Id</label>
        <input
          type="text"
          value={get("resource_id")}
          on:change={set("resource_id")}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex justify-end">
        {#if modified}
          <PrimaryButton onClick={save} icon="save" label="Save" />
        {/if}
      </div>
    </div>
  </div>
</Layout>
