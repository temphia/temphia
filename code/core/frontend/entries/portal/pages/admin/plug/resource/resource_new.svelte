<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../app";
  import { PrimaryButton } from "../../../../../_shared/common";
  import Layout from "../../layout.svelte";

  const app: PortalApp = getContext("__app__");

  export let pid: string;
  export let aid: string;

  let slug = "";
  let plug_id = pid;
  let agent_id = aid;
  let resource_id = "";
  let actions = "";
  let policy = "";

  const save = async () => {
    const papi = await app.get_apm().get_plug_api();
    const data = {
      slug,
      plug_id,
      agent_id,
      resource_id,
      actions,
      policy,
    };

    const resp = await papi.agent_resource_new(pid, aid, data);
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
      <div class="text-2xl text-indigo-900">New Agent Resource</div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Slug</label>
        <input
          type="text"
          bind:value={slug}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Plug Id</label>
        <input
          type="text"
          bind:value={plug_id}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>
      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Agent Id</label>
        <input
          type="text"
          bind:value={agent_id}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>

      <div class="flex-col flex py-3">
        <label class="pb-2 text-gray-700 font-semibold">Resource Id</label>
        <input
          type="text"
          bind:value={resource_id}
          disabled
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
        />
      </div>
    </div>

    <div class="flex justify-end">
      {#if slug && resource_id}
        <PrimaryButton onClick={save} icon="save" label="Save" />
      {/if}
    </div>
  </div>
</Layout>
