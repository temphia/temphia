<script lang="ts">
  import { getContext } from "svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";
  import Layout from "../../layout.svelte";
  import { KvEditor } from "../../../../../common";

  const app: PortalApp = getContext("__app__");
  export let did;

  let message = "";

  let get_renderer_opts;
  let renderer_opts_modified = false;

  let get_meta_opts;
  let meta_opts_modified = false;
  let modified = false;

  let renderers = [];
  let domain_data = {};

  const set = (name) => (ev) => {
    domain_data = { ...domain_data, [name]: ev.target.value };
    modified = true;
  };

  (async () => {
    const tapi = await app.get_apm().get_tenant_id();
    const resp1 = tapi.list_renderer();
    const resp2 = tapi.get_tenant_domain(did);
    domain_data = (await resp2).data;
    renderers = (await resp1).data;
  })();

  const saveHandle = async () => {
    const tapi = await app.get_apm().get_tenant_id();
    const data = { ...domain_data };
    if (renderer_opts_modified) {
      data["renderer_opts"] = get_renderer_opts();
    }

    if (meta_opts_modified) {
      data["extra_meta"] = get_meta_opts();
    }

    const resp = await tapi.update_tenant_domain(did, data);
    if (resp.status !== 200) {
      message = resp.data;
      return;
    }

    app.navigator.goto_org_profile();
  };
</script>

<Layout current_item="ns">
  <div class="h-full w-full overflow-auto">
    <div class="md:p-12 bg-indigo-100 flex flex-row flex-wrap">
      <div
        class="md:w-1/2-screen m-0 p-5 bg-white w-full tw-h-full shadow md:rounded-lg"
      >
        <div class="text-2xl text-indigo-900">Edit Domain</div>

        <p class="text-red-500">{message}</p>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">Name</label>
          <input
            type="text"
            value={domain_data["name"] || ""}
            on:change={set("name")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
            placeholder="example.com"
          />
        </div>

        <div class="flex-col flex py-3">
          <label class="pb-2 text-gray-700 font-semibold">About</label>
          <textarea
            value={domain_data["about"] || ""}
            on:change={set("about")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Serve Cabinet Source</label
          >
          <input
            type="text"
            value={domain_data["serve_source"] || ""}
            on:change={set("serve_source")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Serve Cabinet Folder</label
          >
          <input
            type="text"
            value={domain_data["serve_folder"] || ""}
            on:change={set("serve_folder")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Default User Group</label
          >
          <input
            type="text"
            value={domain_data["default_ugroup"] || ""}
            on:change={set("default_ugroup")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Template Brprint</label
          >
          <input
            type="text"
            value={domain_data["template_bprint"] || ""}
            on:change={set("template_bprint")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Child Template Brprint</label
          >
          <input
            type="text"
            value={domain_data["child_template_bprint"] || ""}
            on:change={set("child_template_bprint")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Http Adapter Type</label>

          <select
            value={domain_data["renderer_type"] || ""}
            on:change={set("renderer_type")}
            class="form-select block w-full  p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          >
            {#each renderers as rtype}
              <option value={rtype}>{rtype}</option>
            {/each}
          </select>
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Http Adapter Plug</label>
          <input
            type="text"
            value={domain_data["renderer_plug_id"] || ""}
            on:change={set("renderer_plug_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Http Adapter Agent</label>
          <input
            type="text"
            value={domain_data["renderer_agent_id"] || ""}
            on:change={set("renderer_agent_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Http Adapter Editor Agent</label
          >
          <input
            type="text"
            value={domain_data["editor_agent_id"] || ""}
            on:change={set("editor_agent_id")}
            class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold"
            >Http Adapter Options</label
          >
          <KvEditor
            data={domain_data["renderer_opts"] || {}}
            bind:getData={get_renderer_opts}
            bind:modified={renderer_opts_modified}
          />
        </div>

        <div class="flex-col flex py-3 relative">
          <label class="pb-2 text-gray-700 font-semibold">Extra Meta</label>
          <KvEditor
            data={domain_data["extra_meta"] || {}}
            bind:getData={get_meta_opts}
            bind:modified={meta_opts_modified}
          />
        </div>

        <div class="flex py-3">
          {#if modified || meta_opts_modified || renderer_opts_modified}
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
