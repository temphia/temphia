<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import SvelteTooltip from "svelte-tooltip";
  import { CEditor } from "../../../../../../shared";
  import { generateId } from "../../../../../../lib/utils";
  import { BetterTextInput } from "../../../../../_shared/common";
  import type { PortalApp } from "../../../../app";

  import WizardLayout from "../../core/wizard_layout.svelte";
  import type { PlugRawSchema } from "./instance";
  import ResourcesPicker from "../../plug/agent/_resources_picker.svelte";
  import NewResource from "../../resource/_new_resource.svelte";
  import { PrimaryButton } from "../../../../../_shared/common";
  import { InstanceHelper } from "./instance";

  export let data: PlugRawSchema;
  export let bid: string;
  export let file: string;
  export let app: PortalApp;

  const instancer = new InstanceHelper(app);

  let slug_check_loading = false;
  let slug_error = "";

  let name = data.name || "";
  let slug = (data.slug || "") + "_" + generateId();
  let resources = {};
  let agent_options = {};

  Object.entries(data.agent_hints || {}).forEach(([aname, ahint]) => {
    agent_options[aname] = {
      name: aname,
    };
  });

  const pick_resource = () => {
    app.simple_modal_open(ResourcesPicker, {
      onSelect: (opts) => console.log(opts),
      app,
    });
  };
  const new_resource = () => {
    app.big_modal_open(NewResource, {
      saveFn: (opts) => console.log(opts),
      app,
    });
  };
</script>

<WizardLayout total_steps={3} curent_step={2}>
  <div class="flex flex-col bg-white border p-8 space-y-4">
    <div class="grid gap-6 mb-2 lg:grid-cols-2">
      <BetterTextInput
        label="Name"
        bind:value={name}
        placeholder="My Plug App"
      />
      <BetterTextInput
        bind:value={slug}
        loading={slug_check_loading}
        error={slug_error}
        label="Slug"
        placeholder="my_plug_app2"
      />
    </div>

    <div class="space-y-2">
      <legend class=" text-base  text-1.5xl font-medium text-gray-900">
        Agents
      </legend>

      <table class="text-left w-full border">
        <thead>
          <tr>
            <th
              class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >
              Name
            </th>

            <th
              class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >
              Executor
            </th>
            <th
              class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >
              Resources
            </th>
          </tr>
        </thead>
        <tbody>
          {#each Object.entries(data.agent_hints) as [agentKey, agent]}
            <tr class="hover:bg-grey-lighter">
              <td class="py-2 px-3 border-b border-grey-light">{agent.name}</td>
              <td class="py-2 px-3 border-b border-grey-light">
                <div>
                  <span
                    class="px-2 text-xs text-fuchsia-600 bg-gray-100 rounded-full"
                  >
                    {agent.executor}
                  </span>
                </div>
              </td>
              <td class="py-2 px-3 border-b border-grey-light">
                <div class="flex flex-col gap-2 ">
                  {#each Object.entries(agent.resources || {}) as [reskey, res]}
                    <div class="p-1 border flex justify-between">
                      <div>
                        <h1>{reskey}</h1>
                        <span class="text-blue-700 underline">{res}</span>
                      </div>

                      <div>
                        <SvelteTooltip tip="Pick resource">
                          <button on:click={pick_resource}>
                            <Icon
                              name="external-link"
                              class="w-6 h-6 text-gray-700"
                            />
                          </button>
                        </SvelteTooltip>

                        <SvelteTooltip tip="Add new resource">
                          <button on:click={new_resource}>
                            <Icon name="plus" class="w-6 h-6 text-gray-700" />
                          </button>
                        </SvelteTooltip>
                      </div>
                    </div>
                  {/each}
                </div></td
              >
            </tr>
          {/each}
        </tbody>
      </table>
    </div>

    <div class="space-y-2">
      <legend class=" text-base  text-1.5xl font-medium text-gray-900">
        Resources
      </legend>

      <div class="flex flex-col gap-2">
        {#each Object.entries(resources || {}) as [res_name, res]}
          <div class="flex justify-between rounded border p-1">
            <div class="flex items-center">
              <Icon name="paper-clip" class="w-6 h-6 text-gray-700" />

              <div class="flex flex-col items-center mx-5 space-y-1">
                <h2 class="text-lg font-medium text-gray-700">{res_name}</h2>
                <div class="px-2 text-xs text-red-500 bg-gray-100 rounded-full">
                  {res["type"]}
                </div>
              </div>
            </div>

            <Icon
              name="x-circle"
              class="w-8 h-8 text-blue-500 cursor-pointer"
            />
          </div>
        {/each}
      </div>
    </div>

    <details>
      <summary class="mb-2 text-sm font-medium text-gray-900"
        >Raw Schema</summary
      >
      <div class="bg-white border p-2">
        <CEditor
          code={JSON.stringify(data, null, 4)}
          container_style="height:20rem;"
          on:change={(ev) => {}}
        />
      </div>
    </details>

    <div class="flex justify-end">
      <PrimaryButton
        onClick={() => (instancer.instance_plug(bid, file, {
            agent_opts: agent_options,
            new_plug_id: slug,
            new_plug_name: name,
            resources: resources}))
        }
        icon="lightning-bolt"
        label="Finish"
      />
    </div>
  </div>
</WizardLayout>
