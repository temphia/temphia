<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import Toggle from "../../../../components/button/toggle.svelte";
  import BetterTextInput from "../../../common/better_text_input.svelte";

  const tables = [
    {
      name: "big table1",
      slug: "b1",
      description: "this is very big table",
    },
  ];
</script>

<div class="flex flex-col bg-white border p-16 space-y-4">
  <div class="grid gap-6 mb-2 lg:grid-cols-2">
    <BetterTextInput
      placeholder="My Data Group"
      label="Group Name"
      value="My Data Group"
    />
    <BetterTextInput placeholder="datagrp1" label="Group Slug" value="dg1" />
  </div>

  <div>
    <legend class="text-base text-1.5xl font-medium text-gray-900">
      Tables
    </legend>
    <div class="space-y-2 p-4 overflow-auto">
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
              Slug
            </th>
            <th
              class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >
              Description
            </th>
            <th
              class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >
              Activity Log Type
            </th>
            <th
              class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
            >
              Seed Type
            </th>

            <th
            class="py-2 px-3 bg-grey-lightest font-bold uppercase text-sm text-grey-dark border-b border-grey-light"
          >
            Sync Type
          </th>
          </tr>
        </thead>
        <tbody>
          {#each tables as table}
            <tr class="hover:bg-grey-lighter">
              <td class="py-2 px-3 border-b border-grey-light">{table.name}</td>
              <td class="py-2 px-3 border-b border-grey-light">{table.slug}</td>
              <td class="py-2 px-3 border-b border-grey-light">
                {table.description}
              </td>
              <td class="py-2 px-3 border-b border-grey-light">
                <select class="p-1 rounded bg-slate-300">
                  <option>strict</option>
                  <option>lazy</option>
                  <option>none</option>
                </select>
              </td>
              <td class="py-2 px-3 border-b border-grey-light">
                <select class="p-1 rounded bg-slate-300">
                  <option>none</option>
                  <option>data</option>
                  <option>autogen</option>
                </select>
              </td>

              <td class="py-2 px-3 border-b border-grey-light">
                <select class="p-1 rounded bg-slate-300">
                  <option>none</option>
                  <option>event_only</option>
                  <option>event_and_data</option>
                </select>
              </td>
            </tr>
          {/each}
        </tbody>
      </table>
    </div>
  </div>

  <details>
    <summary class="mb-2 text-sm font-medium text-gray-900">Raw Schema</summary>
    <div class="bg-white border p-2">
      Epcot is a theme park at Walt Disney World Resort featuring exciting
      attractions, international pavilions, award-winning fireworks and seasonal
      special events.
    </div>
  </details>
</div>
