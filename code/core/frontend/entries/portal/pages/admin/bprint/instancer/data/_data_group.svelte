<script lang="ts">
  import {
    WizardLayout,
    CEditor,
    BetterTextInput,
    PortalService,
    PrimaryButton,
  } from "../../../core";
  import type { NewTableGroup } from "../instance";
  import { InstanceHelper } from "../instance";

  export let data: NewTableGroup;
  export let bid: string;
  export let file: string;
  export let app: PortalService;

  const instancer = new InstanceHelper(app);

  let dyndb_source: string = "";
  let group_name: string = data.name;
  let group_slug: string = data.slug + Math.random().toString(36).slice(7);
  let cabinet_source: string = "";
  let cabinet_folder: string = "";
  let seed_source: string = "";
</script>

<WizardLayout total_steps={3} curent_step={2}>
  <div class="flex flex-col bg-white border p-4 space-y-4">
    <div class="grid gap-6 mb-2 lg:grid-cols-2">
      <BetterTextInput
        placeholder="My Data Group"
        label="Group Name"
        bind:value={group_name}
      />
      <BetterTextInput
        placeholder="datagrp1"
        label="Group Slug"
        bind:value={group_slug}
      />
    </div>

    <div>
      <legend class="text-base text-1.5xl font-medium text-gray-900 mb-2">
        Tables
      </legend>
      <div class="overflow-auto shadow p-1">
        <table class="text-left w-full border">
          <thead>
            <tr>
              <th
                class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
              >
                Name
              </th>
              <th
                class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
              >
                Slug
              </th>
              <th
                class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
              >
                Description
              </th>
              <th
                class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
              >
                Activity Log Type
              </th>
              <th
                class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
              >
                Sync Type
              </th>

              <th
                class="p-1 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
              >
                Seed
              </th>
            </tr>
          </thead>
          <tbody>
            {#each data.tables as table}
              <tr class="hover:bg-grey-lighter">
                <td class="p-1 border-b border-grey-light">{table.name}</td>
                <td class="p-1 border-b border-grey-light">{table.slug}</td>
                <td class="p-1 border-b border-grey-light">
                  {table.description}
                </td>
                <td class="p-1 border-b border-grey-light">
                  <select class="p-1 rounded bg-slate-300">
                    <option>strict</option>
                    <option>lazy</option>
                    <option>none</option>
                  </select>
                </td>
                <td class="p-1 border-b border-grey-light">
                  <select class="p-1 rounded bg-slate-300">
                    <option>none</option>
                    <option>event_only</option>
                    <option>event_and_data</option>
                  </select>
                </td>

                <td class="p-1 border-b border-grey-light">
                  <input type="checkbox" checked />
                </td>
              </tr>
            {/each}
          </tbody>
        </table>
      </div>
    </div>

    <div class="flex justify-around">
      <div>
        <legend class="text-base text-1.5xl font-medium text-gray-900 mb-2">
          Seed Source
        </legend>
        <select class="p-1 rounded bg-slate-300" bind:value={seed_source}>
          <option value="">none</option>
          <option>data</option>
          <option>autogen</option>
        </select>
      </div>

      <div>
        <legend class="text-base text-1.5xl font-medium text-gray-900 mb-2">
          Data Source
        </legend>
        <select class="p-1 rounded bg-slate-300">
          <option value="">none</option>
        </select>
      </div>

      <div>
        <legend class="text-base text-1.5xl font-medium text-gray-900 mb-2">
          Cabinet Source
        </legend>
        <select class="p-1 rounded bg-slate-300">
          <option value="">none</option>
        </select>
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
        onClick={() =>
          instancer.instance_data_group(bid, file, {
            dyndb_source,
            group_name,
            group_slug,
            cabinet_source,
            cabinet_folder,
            seed_source,
          })}
        icon="lightning-bolt"
        label="Finish"
      />
    </div>
  </div>
</WizardLayout>
