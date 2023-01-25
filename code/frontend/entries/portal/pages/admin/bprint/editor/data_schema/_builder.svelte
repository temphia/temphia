<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import {
    ActionDeleteButton,
    ActionEditButton,
  } from "../../../../../../xcompo";
  import Cicon from "../../../../data/table/core/cicon/cicon.svelte";
  import type { Builder } from "./builder";

  import AddColumn from "./panels/_add_column.svelte";
  import AddTable from "./panels/_add_table.svelte";
  import EditColumn from "./panels/_edit_column.svelte";
  import EditTable from "./panels/_edit_table.svelte";
  import { Validate } from "./validator";

  export let open_modal: (compo: any, opts: any) => void;
  export let close_modal: () => void;

  export let builder: Builder;

  let collapsed = [];
  const state = builder.state;
  $: __schema = $state;

  const action_add_table = () => {
    open_modal(AddTable, {
      callback: (data: any) => {
        builder.add_table(data);
        close_modal();
      },
    });
  };

  const action_edit_table = (table: string) => () => {
    open_modal(EditTable, {
      data: table,
      callback: (data) => {
        builder.edit_table(table, data);
        close_modal();
      },
    });
  };

  const action_delete_table = (table: string) => () => {
    builder.delete_table(table);
  };

  const action_add_column = (table: string) => () => {
    open_modal(AddColumn, {
      current_schema: __schema,
      current_table: table,
      callback: (data, ref_data) => {
        builder.add_column(data, ref_data);
        close_modal();
      },
    });
  };

  const action_edit_column = (table: string, col: any) => () => {
    open_modal(EditColumn, {
      data: col,
      callback: (data) => {
        builder.edit_column(table, col.slug, data);
        close_modal();
      },
    });
  };

  const action_delete_column = (table: string, column: string) => () => {
    builder.delete_column(table, column);
  };

  const action_delete_column_ref = (table: string, refidx: number) => {
    // builder.delete_column_ref()
  };

  let validation_message = "";
  const action_validator = () => {
    validation_message = Validate(__schema);
  };
</script>

<div class="bg-blue-100 p-10 w-full h-full overflow-auto text-gray-800">
  <div class="bg-white rounded p-2 flex content-center justify-between mb-4">
    <div class="flex gap-2 items-center">
      <h1 class="font-medium text-xl text-gray-800 line-clamp-1">
        {__schema.name || ""}
      </h1>

      <span class="px-1 bg-gray-400 text-gray-50 rounded font-semibold"
        >{__schema.slug || ""}</span
      >
    </div>

    <div class="flex gap-1">
      <button
        class="hover:bg-gray-300 rounded inline-flex border p-1"
        on:click={action_add_table}
      >
        <Icon name="plus" class="h-5 w-5" />
        Add
      </button>

      <button
        class="hover:bg-gray-300 rounded inline-flex border p-1"
        on:click={action_validator}
      >
        <Icon name="check" class="h-5 w-5" />
        Validate
      </button>

    </div>
  </div>

  <div class="bg-white rounded">
    <p class="text-red-500">{validation_message}</p>
  </div>

  <div class="flex flex-wrap gap-2">
    {#each __schema.tables || [] as table, idx}
      
      {#if table !== null}
        <div class="p-2 border bg-white shadow w-90 rounded-lg">
          <div
            class="flex flex-wrap content-center justify-between border-b mb-2"
          >
            <div
              class="flex p-2 content-center justify-around gap-2 items-center"
            >
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
                  fill="currentColor"
                  class="bi bi-triangle-fill w-4 h-4 text-gray-700 transition-all {collapsed.includes(
                    idx
                  )
                    ? 'rotate-90'
                    : 'rotate-180'}"
                  viewBox="0 0 16 16"
                >
                  <path
                    fill-rule="evenodd"
                    d="M7.022 1.566a1.13 1.13 0 0 1 1.96 0l6.857 11.667c.457.778-.092 1.767-.98 1.767H1.144c-.889 0-1.437-.99-.98-1.767L7.022 1.566z"
                  />
                </svg>
              </button>

              <h1 class="text-gray-800">{table.name || ""}</h1>
              <span class="px-1 bg-blue-400 text-blue-50 rounded font-semibold"
                >{table.slug || ""}</span
              >
            </div>
            <div class="flex gap-1 p-2">
              <button
                class="hover:bg-gray-300 rounded inline-flex border p-1"
                on:click={action_add_column(table.slug)}
              >
                <Icon name="plus" class="h-5 w-5" />
                Column
              </button>

              <button
                class="hover:bg-gray-300 rounded inline-flex border p-1"
                on:click={action_edit_table(table.slug)}
              >
                <Icon name="pencil-alt" class="h-5 w-5" />
                Edit
              </button>

              <button
                on:click={action_delete_table(table.slug)}
                class="hover:bg-gray-300 rounded inline-flex border p-1"
              >
                <Icon name="trash" class="h-5 w-5" />
                Delete
              </button>
            </div>
          </div>

          {#if !collapsed.includes(idx)}
            {#if table.columns && table.columns.length > 0}
              <h2
                class="font-medium text-base md:text-lg text-gray-800 line-clamp-1 py-2"
              >
                Table
              </h2>

              <table
                class="table-auto border-collapse w-full border-2 rounded"
              >
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
                      <td class="px-3 py-1">
                        <div class="p-1 rounded-lg inline-flex">
                          <Cicon ctype={col.ctype} />
                          <span>{col.name || ""}</span>
                        </div>
                      </td>
                      <td class="px-3 py-1"
                        ><span class="p-1 rounded-lg">{col.slug}</span></td
                      >
                      <td class="px-3 py-1"
                        ><span class="p-1 rounded-lg">{col.ctype}</span></td
                      >

                      <td class="px-3 py-1 flex gap-2">
                        <ActionEditButton
                          onClick={action_edit_column(table.slug, col)}
                        />
                        <ActionDeleteButton
                          onClick={action_delete_column(table.slug, col.slug)}
                        />
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            {/if}

            {#if table.column_refs && table.column_refs.length > 0}
              <h2
                class="font-medium text-base md:text-lg text-gray-800 line-clamp-1 py-2"
              >
                References
              </h2>

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
                      style="background-color: rgb(248, 248, 248);"
                      >Target Table</th
                    >

                    <th
                      class="px-2 py-1"
                      style="background-color: rgb(248, 248, 248);"
                      >From Columns</th
                    >

                    <th
                      class="px-2 py-1"
                      style="background-color: rgb(248, 248, 248);"
                      >To Columns</th
                    >
                    <th
                      class="px-2 py-1"
                      style="background-color: rgb(248, 248, 248);"
                    />
                  </tr>
                </thead>

                <tbody class="text-sm font-normal text-gray-700">
                  {#each table.column_refs as cref, crefidx}
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
                        ><span
                          class="p-1 rounded-lg"
                          >{cref.target}</span
                        ></td
                      >

                      <td class="px-3 py-1">
                        <div class="flex">
                          {#each cref.from_cols as col}
                            <span class="p-1 rounded bg-slate-400 text-slate-50"
                              >{col}</span
                            >
                          {/each}
                        </div>
                      </td>

                      <td class="px-3 py-1">
                        <div class="flex">
                          {#each cref.to_cols as col}
                            <span class="p-1 rounded bg-red-500 text-red-50"
                              >{col}</span
                            >
                          {/each}
                        </div>
                      </td>

                      <td class="px-3 py-1 flex gap-2">
                        <ActionDeleteButton
                          onClick={action_delete_column_ref(
                            table.slug,
                            crefidx
                          )}
                        />
                      </td>
                    </tr>
                  {/each}
                </tbody>
              </table>
            {/if}
          {:else}
            <div class="text-blue-500 font-semibold">{"{...}"}</div>
          {/if}
        </div>
      {/if}
    {/each}
  </div>
</div>
