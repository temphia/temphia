<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import Layout from "../../layout.svelte";
  import { PrimaryButton } from "../../../../../_shared/common";

  const app: PortalApp = getContext("__app__");

  export let pid: string;
  export let aid: string;

  let name = "";
  let from_plug_id = pid;
  let from_agent_id = aid;
  let to_plug_id = "";
  let to_agent_id = "";
  let to_handler = "";

  const save = async () => {
    const papi = await app.get_apm().get_plug_api();
    const data = {
      name,
      from_plug_id,
      from_agent_id,
      to_plug_id,
      to_agent_id,
      to_handler,
    };

    const resp = await papi.agent_link_new(pid, aid, data);
    if (resp.status !== 200) {
      console.log("Err", resp);
      return;
    }
    app.navigator.goto_admin_agents_page(pid);
  };
</script>

<Layout>
  <div class="h-full w-full bg-indigo-100 p-10 overflow-auto">
    <div class="p-5 bg-white w-full ">
      <div class="text-2xl text-indigo-900">New Link</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Name</label>
        <input
          type="text"
          bind:value={name}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">From Plug</label>
        <input
          type="text"
          bind:value={from_plug_id}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">From Agent</label>
        <input
          type="text"
          bind:value={from_agent_id}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">To Plug</label>
        <input
          type="text"
          bind:value={to_plug_id}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">To Agent</label>
        <input
          type="text"
          bind:value={to_agent_id}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex justify-end">
        {#if name && to_plug_id && to_agent_id}
          <PrimaryButton onClick={save} icon="save" label="Save" />
        {/if}
      </div>
    </div>
  </div>
</Layout>
