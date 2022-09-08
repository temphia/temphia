<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";
  import { KvEditor, ActionButton } from "../../../../../common";
  import PlugPick from "../../core/plug_pick.svelte";

  const app: PortalApp = getContext("__app__");

  export let name = "";
  export let about = "";
  export let default_ugroup = "";
  export let template_bprint = "";
  export let child_template_bprint = "";

  export let renderer_type = "";
  export let renderer_plug_id = "";
  export let renderer_agent_id = "";
  export let renderer_opts = {};
  export let editor_agent_id = "";
  export let serve_source = "";
  export let serve_folder = "";
  export let extra_meta = {};

  let renderers = [];
  let message = "";

  (async () => {
    const tapi = await app.get_apm().get_tenant_id();
    const resp = await tapi.list_renderer();
    renderers = resp.data;
  })();

  const saveHandle = async () => {
    const data = {
      name,
      about,
      default_ugroup,
      template_bprint,
      child_template_bprint,

      renderer_type,
      renderer_plug_id,
      renderer_agent_id,
      renderer_opts,

      editor_agent_id,
      serve_source,
      serve_folder,
      extra_meta,
    };

    const tapi = await app.get_apm().get_tenant_id();
    const resp = await tapi.add_tenant_domain(data);
    if (resp.status !== 200) {
      message = resp.data;
      return;
    }
    app.navigator.goto_admin_org();
  };

  const rendererPick = () => {
    app.simple_modal_open(PlugPick, {
      app,
      selected_plug: renderer_plug_id,
      selected_agent: renderer_agent_id,
      onSelected: (data) => {
        renderer_plug_id = data.plug;
        renderer_agent_id = data.agent;
      },
    });
  };
  const editorPick = () => {
    app.simple_modal_open(PlugPick, {
      app,
      selected_plug: renderer_plug_id,
      selected_agent: editor_agent_id,
      onSelected: (data) => {
        renderer_plug_id = data.plug;
        editor_agent_id = data.agent;
      },
    });
  };
</script>

<Layout current_item="ns">
  <div class="h-full w-full overflow-auto">
    <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
      <div
        class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
      >
        <div class="text-2xl text-indigo-900">Add Domain</div>
        <p class="text-red-500">{message}</p>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            bind:value={name}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="example.com"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">About</label>
          <textarea
            bind:value={about}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Serve Cabinet Source</label
          >
          <input
            type="text"
            bind:value={serve_source}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Serve Cabinet Folder</label
          >
          <input
            type="text"
            bind:value={serve_folder}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Default User Group</label
          >
          <input
            type="text"
            bind:value={default_ugroup}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Template Brprint</label
          >
          <input
            type="text"
            bind:value={template_bprint}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Child Template Brprint</label
          >
          <input
            type="text"
            bind:value={child_template_bprint}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Http Adapter Type</label>

          <select
            bind:value={renderer_type}
            class="form-select block w-full  p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          >
            {#each renderers as rtype}
              <option value={rtype}>{rtype}</option>
            {/each}
          </select>
        </div>

        <div class="flex-col flex py-3 relative">
          <div class="absolute right-1">
            <ActionButton icon_name="link" name="pick" onClick={rendererPick} />
          </div>

          <label class="pb-2 text-gray-700 font-semibold">Http Adapter Plug</label>
          <input
            type="text"
            bind:value={renderer_plug_id}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Http Adapter Agent</label>
          <input
            type="text"
            bind:value={renderer_agent_id}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <div class="absolute right-1">
            <ActionButton icon_name="link" name="pick" onClick={editorPick} />
          </div>
          <label class="pb-2 text-gray-700 font-semibold"
            >Http Adapter Editor Agent</label
          >
          <input
            type="text"
            bind:value={editor_agent_id}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Http Adapter Options</label
          >
          <KvEditor bind:data={renderer_opts} />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
          <KvEditor bind:data={extra_meta} />
        </div>

        {#if name}
          <div class="flex py-3">
            <button
              class="p-2 bg-blue-400 m-1 w-20 text-white rounded"
              on:click={saveHandle}>Save</button
            >
          </div>
        {/if}
      </div>
    </div>
  </div>
</Layout>
