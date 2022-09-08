<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";

  import { KvEditor } from "../../../../../common";
  import PrimaryButton from "../../../../../common/action_button/primary_button.svelte";

  const app: PortalApp = getContext("__app__");

  export let pid: string;
  export let aid: string;

  let name = "";
  let plug_id = pid;
  let agent_id = aid;
  let brpint_id = "";
  let ref_file = "";
  let extra_meta = {};

  let meta_modified = false;

  const save = async () => {
    const papi = await app.get_apm().get_plug_api();
    const data = {
      name,
      plug_id,
      agent_id,
      brpint_id,
      ref_file,
      extra_meta,
    };

    const resp = await papi.agent_extension_new(pid, aid, data);
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
      <div class="text-2xl text-indigo-900">New Extension</div>

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
          bind:value={plug_id}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Agent</label>
        <input
          type="text"
          bind:value={agent_id}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Bprint</label>
        <input
          type="text"
          bind:value={brpint_id}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">File</label>
        <input
          type="text"
          bind:value={ref_file}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
        <KvEditor bind:data={extra_meta} bind:modified={meta_modified} />
      </div>

      <div class="flex justify-end">
        {#if plug_id && agent_id && brpint_id && ref_file}
          <PrimaryButton onClick={save} icon="save" label="Save" />
        {/if}
      </div>
    </div>
  </div>
</Layout>
