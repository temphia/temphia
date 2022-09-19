<script lang="ts">
  import { getContext } from "svelte";
  import { CEditor } from "../../../../../../components";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  import InstallerLayout from "../_layout.svelte";
  import { installAsPlug } from "./installer";

  export let bid;
  export let data;
  export let schema;

  const app: PortalApp = getContext("__app__");

  let installer;
  app
    .get_apm()
    .get_bprint_api()
    .then((bapi) => {
      installer = installAsPlug(bapi);
    });

  const _all_agents = Object.keys(schema.agent_hints || {}) || [];
  const _all_resources = Object.keys(schema.resource_hints || {}) || [];

  $: _agent_hints = [..._all_agents];
  $: _resource_hints = [..._all_resources];

  let _schema_modified = false;
  let _new_plug_id = "";

  const pick_agent = (_agent) => () => {
    const index = _agent_hints.indexOf(_agent);
    if (index > -1) {
      _agent_hints.splice(index, 1);
    } else {
      _agent_hints.push(_agent);
    }
    _agent_hints = [..._agent_hints];
  };

  const pick_resource = (_resource) => () => {
    const index = _resource_hints.indexOf(_resource);
    if (index > -1) {
      _resource_hints.splice(index, 1);
    } else {
      _resource_hints.push(_resource);
    }
    _resource_hints = [..._resource_hints];
  };

  let editor;
</script>

<InstallerLayout
  last_stage_name="Install"
  description={data["description"]}
  files={data["files"]}
  type={data["type"]}
  name={data["name"]}
  slug={data["slug"]}
  subtype={data["sub_type"]}
  source=""
  final_func={app.big_modal_close}
  last_page_func={() => {
    let _schema = "";
    if (_schema_modified) {
      _schema = editor.getValue();
    }

    installer({
      agents: _agent_hints,
      resources: _resource_hints,
      bprint_id: bid,
      new_plug_id: _new_plug_id,
      schema: _schema,
    });
  }}
>
  <div slot="options" class="h-full w-full">
    <div class="flex flex-col relative">
      <button
        class="p-1 text-gray-500 right-1 top-1/2 border absolute bg-white rounded"
        >Check</button
      >
      <label class="leading-loose">Id</label>
      <input
        type="text"
        bind:value={_new_plug_id}
        class="px-4 py-2 border w-full sm:text-sm rounded focus:outline-none focus:border-indigo-500"
        placeholder="Optional"
      />
    </div>

    {#if _all_agents.length > 0}
      <legend class="text-base mt-5 text-1.5xl font-medium text-gray-900"
        >Auto Create Agents</legend
      >
      <div class="mt-2 space-y-4">
        {#each _all_agents as ah}
          <div class="flex place-items-center">
            <div class="flex items-center h-5">
              <input
                type="checkbox"
                on:change={pick_agent(ah)}
                checked={_agent_hints.includes(ah)}
                class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
              />
            </div>
            <div class="ml-3 text-sm">
              <label for="comments" class="font-regular text-gray-700"
                >{ah}</label
              >
            </div>
          </div>
        {/each}
      </div>
    {/if}

    {#if _all_resources.length > 0}
      <legend class="text-base mt-5 text-1.5xl font-medium text-gray-900"
        >Auto Create Resources</legend
      >
      <div class="mt-2 space-y-4">
        {#each _all_resources as ah}
          <div class="flex place-items-center">
            <div class="flex items-center h-5">
              <input
                type="checkbox"
                checked={_resource_hints.includes(ah)}
                on:change={pick_resource(ah)}
                class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"
              />
            </div>
            <div class="ml-3 text-sm">
              <label for="comments" class="font-regular text-gray-700"
                >{ah}</label
              >
            </div>
          </div>
        {/each}
      </div>
    {/if}

    <div class="flex-col flex py-3 relative">
      <button
        class="p-1 text-gray-500 right-1 top-1/2 border absolute bg-white rounded"
        >Check</button
      >
      <legend class="text-base font-medium text-gray-900 mt-5">Schema</legend>
      <CEditor
        bind:editor
        code={JSON.stringify(schema, null, 4)}
        container_style="height:20rem;"
        on:change={(ev) => {}}
      />
    </div>
  </div>

  <div slot="final" class="" />
</InstallerLayout>
