<script lang="ts">
  import type { PortalApp } from "../../../../app";
  import {
    ActionAddButton,
    ActionDeleteButton,
    ActionEditButton,
  } from "../../../../../_shared/common";

  import type Schema from "./sample";
  import AddColumn from "./_add_column.svelte";
  import AddTable from "./_add_table.svelte";
  export let schema: typeof Schema;
  export let app: PortalApp;

  let collapsed = [];
</script>

<div class="bg-blue-100 p-10 w-full h-full overflow-auto">
  <div
    class="bg-white rounded p-2 flex content-center justify-around mb-6 gap-2"
  >
    <label for="table_name" class="block text-sm font-medium text-gray-900 p-2"
      >Name</label
    >
    <input
      type="text"
      bind:value={schema.name}
      class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
      required
    />

    <label for="table_slug" class="block text-sm font-medium text-gray-900 p-2"
      >Slug</label
    >
    <input
      type="text"
      bind:value={schema.slug}
      class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
      required
    />

    <ActionAddButton
      onClick={() => {
        app.simple_modal_open(AddTable, {});
      }}
      name="Table"
    />
  </div>

  {#each schema.tables || [] as table, idx}
    {#if table !== null}
      <div class="p-2 border bg-white shadow mb-2">
        <div class="flex flex-wrap content-center justify-between">
          <div class="flex p-2 content-center justify-around">
            <button
              class="border p-1 rounded"
              on:click={() => {
                if (collapsed.includes(idx)) {
                  collapsed = collapsed.filter((v) => v !== idx);
                } else {
                  collapsed = [...collapsed, idx];
                }
              }}
            >
              <svg
                xmlns="http://www.w3.org/2000/svg"
                class="h-6 w-6"
                fill="none"
                viewBox="0 0 24 24"
                stroke="currentColor"
                stroke-width="2"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  d="M13 5l7 7-7 7M5 5l7 7-7 7"
                />
              </svg>
            </button>

            <label
              for="table_name"
              class="block text-sm font-medium text-gray-900 p-2">Name</label
            >
            <input
              type="text"
              value={table.name}
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
              required
            />

            <label
              for="table_slug"
              class="block text-sm font-medium text-gray-900 p-2">Slug</label
            >
            <input
              type="text"
              value={table.slug}
              class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5"
              disabled
            />
          </div>
          <div class="flex justify-end content-end justify-items-end">
            <ActionAddButton
              onClick={() => {
                app.simple_modal_open(AddColumn, {});
              }}
              name="Column"
            />
          </div>
        </div>

        {#if !collapsed.includes(idx)}
          <table class="table-auto border-collapse w-full border">
            <thead
              ><tr
                class="rounded-lg text-sm font-medium text-gray-700 text-left"
                ><th
                  class="px-2 py-1"
                  style="background-color: rgb(248, 248, 248);">Name</th
                ><th
                  class="px-2 py-1"
                  style="background-color: rgb(248, 248, 248);">Slug</th
                >

                <th
                  class="px-2 py-1"
                  style="background-color: rgb(248, 248, 248);">ctype</th
                >

                <th
                  class="px-2 py-1"
                  style="background-color: rgb(248, 248, 248);"
                />
              </tr>
            </thead>

            <tbody class="text-sm font-normal text-gray-700">
              {#each table.columns as col}
                <tr
                  class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
                >
                  <td class="px-3 py-1"
                    ><span class="p-1 rounded-lg">{col.name}</span></td
                  >
                  <td class="px-3 py-1"
                    ><span class="p-1 rounded-lg">{col.slug}</span></td
                  >
                  <td class="px-3 py-1"
                    ><span class="p-1 rounded-lg">{col.ctype}</span></td
                  >

                  <td class="px-3 py-1 flex gap-2">
                    <ActionEditButton onClick={() => {}} />
                    <ActionDeleteButton onClick={() => {}} />
                  </td>
                </tr>
              {/each}
            </tbody>
          </table>

          {#if table.column_refs}
            <table class="table-auto border-collapse w-full border">
              <thead
                ><tr
                  class="rounded-lg text-sm font-medium text-gray-700 text-left"
                  ><th
                    class="px-2 py-1"
                    style="background-color: rgb(248, 248, 248);">Slug</th
                  ><th
                    class="px-2 py-1"
                    style="background-color: rgb(248, 248, 248);">Type</th
                  >

                  <th
                    class="px-2 py-1"
                    style="background-color: rgb(248, 248, 248);">Target</th
                  >

                  <th
                    class="px-2 py-1"
                    style="background-color: rgb(248, 248, 248);"
                    >From Columns</th
                  >

                  <th
                    class="px-2 py-1"
                    style="background-color: rgb(248, 248, 248);">To Columns</th
                  >
                  <th
                    class="px-2 py-1"
                    style="background-color: rgb(248, 248, 248);"
                  />
                </tr>
              </thead>

              <tbody class="text-sm font-normal text-gray-700">
                {#each table.column_refs as cref}
                  <tr
                    class="hover:bg-gray-100 border-b border-gray-200 py-10 text-gray-700"
                  >
                    <td class="px-3 py-1"
                      ><span class="p-1 rounded-lg">{cref.slug}</span></td
                    >
                    <td class="px-3 py-1"
                      ><span class="p-1 rounded-lg">{cref.type}</span></td
                    >
                    <td class="px-3 py-1"
                      ><span class="p-1 rounded-lg">{cref.target}</span></td
                    >

                    <td class="px-3 py-1">
                      <div class="flex">
                        {#each cref.from_cols as col}
                          <span class="p-1 rounded bg-slate-400">{col}</span>
                        {/each}
                      </div>
                    </td>

                    <td class="px-3 py-1">
                      <div class="flex">
                        {#each cref.to_cols as col}
                          <span class="p-1 rounded bg-slate-400">{col}</span>
                        {/each}
                      </div>
                    </td>

                    <td class="px-3 py-1 flex gap-2">
                      <ActionEditButton onClick={() => {}} />
                      <ActionDeleteButton onClick={() => {}} />
                    </td>
                  </tr>
                {/each}
              </tbody>
            </table>
          {/if}
        {/if}
      </div>
    {/if}
  {/each}
</div>
