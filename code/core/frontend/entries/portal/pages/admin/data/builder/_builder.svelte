<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { hslColor, numHash } from "../../../../../../lib/utils";
  import { ActionDeleteButton, ActionEditButton } from "../../../../../xcompo";
  import Cicon from "../../../data/tableui/core/cicon/cicon.svelte";
  import type Schema from "./sample";

  import AddColumn from "./_add_column.svelte";
  import AddTable from "./_add_table.svelte";
  import EditColumn from "./_edit_column.svelte";
  import EditTable from "./_edit_table.svelte";
  export let schema: typeof Schema;

  export let open_modal: (compo: any, opts: any) => void;
  export let close_modal: () => void;

  let collapsed = [];

  $: __schema = { ...schema };

  const get_column_index = (table: string, column: string) => {
    let _tidx = -1;
    let _cidx = -1;

    __schema.tables.forEach((val, tidx) => {
      if (val.slug !== table) {
        return;
      }
      val.columns.forEach((col, cidx) => {
        if (col.slug !== column) {
          return;
        }
        _tidx = tidx;
        _cidx = cidx;
        return;
      });
    });

    return [_tidx, _cidx];
  };

  const get_table_index = (table: string) => {
    let _tidx = -1;
    __schema.tables.forEach((val, tidx) => {
      if (val.slug !== table) {
        return;
      }
      _tidx = tidx;
    });

    return _tidx;
  };

  const action_add_table = (data) => {
    __schema = {
      ...__schema,
      tables: [
        ...__schema.tables,
        {
          name: data.name,
          slug: data.slug,
          description: data.info || "",
          icon: data["icon"] || "",
          main_column: data["main_column"] || "",
          columns: data["columns"] || [],
          column_refs: data["column_refs"] || [],
        },
      ],
    };

    close_modal();
  };

  $: console.log("$SCHEMA", __schema);

  const action_edit_table = (table: string, data: any) => {
    const tidx = get_table_index(table);
    const tbl = __schema.tables[tidx];

    __schema.tables[tidx] = {
      ...tbl,
      ...{
        name: data.name,
        slug: data.slug,
        description: data.info || "",
      },
    };

    __schema = { ...__schema };
  };

  const action_delete_table = (slug: string) => {
    __schema.tables = __schema.tables.filter((val) => val.slug !== slug);
    __schema = { ...__schema };
  };

  const action_add_column = (table: string, data: any) => {
    const tidx = get_table_index(table);

    const tbl = __schema.tables[tidx];

    tbl.columns = [
      {
        name: data["name"] || "",
        slug: data["slug"] || "",
        ctype: data["ctype"],
        description: data["info"],
        icon: data["icon"] || "",
        options: data["options"] || [],
        not_nullable: !!data["not_nullable"],
        pattern: data["pattern"],
        strict_pattern: !!data["strict_pattern"],
      },
      ...tbl.columns,
    ];

    __schema = { ...__schema };
    close_modal();
  };

  const action_edit_column = (table: string, column: string, data: any) => {
    const [tidx, cidx] = get_column_index(table, column);
    const col = __schema.tables[tidx].columns[cidx];

    __schema.tables[tidx].columns[cidx] = {
      ...col,
      name: data["name"] || "",
      slug: data["slug"] || "",
      ctype: data["ctype"],
      description: data["info"],
    };

    __schema = { ...__schema };
  };

  const action_delete_column = (table: string, column: string) => {
    const [tidx, cidx] = get_column_index(table, column);
    const tbl = __schema.tables[tidx];

    tbl.columns = tbl.columns.filter((val) => val.slug != column);
    __schema = { ...__schema };
  };
</script>

<div class="bg-blue-100 p-10 w-full h-full overflow-auto text-gray-800">
  <div class="bg-white rounded p-2 flex content-center justify-between mb-4">
    <div class="flex gap-2 items-center">
      <h1 class="font-medium text-xl text-gray-800 line-clamp-1">
        {__schema.name}
      </h1>

      <span class="px-1 bg-gray-400 text-gray-50 rounded font-semibold"
        >{__schema.slug}</span
      >
    </div>

    <div class="flex gap-1">
      <button
        class="hover:bg-gray-300 rounded inline-flex border p-1"
        on:click={() =>
          open_modal(AddTable, {
            callback: action_add_table,
          })}
      >
        <Icon name="plus" class="h-5 w-5" />
        Add
      </button>
    </div>
  </div>

  <div class="flex flex-wrap gap-2">
    {#each __schema.tables || [] as table, idx}
      {@const table_color = `border-color: hsl(${
        numHash(table.slug) % 360
      }, 100%, 80%)`}
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

              <h1 class="text-gray-800">{table.name}</h1>
              <span class="px-1 bg-blue-400 text-blue-50 rounded font-semibold"
                >{table.slug}</span
              >
            </div>
            <div class="flex gap-1 p-2">
              <button
                class="hover:bg-gray-300 rounded inline-flex border p-1"
                on:click={() =>
                  open_modal(AddColumn, {
                    current_schema: __schema,
                    callback: (data) => action_add_column(table.slug, data),
                  })}
              >
                <Icon name="plus" class="h-5 w-5" />
                Column
              </button>

              <button
                class="hover:bg-gray-300 rounded inline-flex border p-1"
                on:click={() => open_modal(EditTable, {})}
              >
                <Icon name="pencil-alt" class="h-5 w-5" />
                Edit
              </button>

              <button
                on:click={() => action_delete_table(table.slug)}
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
                style={table_color}
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
                          <span>{col.name}</span>
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
                          onClick={() => open_modal(EditColumn, {})}
                        />
                        <ActionDeleteButton
                          onClick={() =>
                            action_delete_column(table.slug, col.slug)}
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
                        ><span
                          class="p-1 rounded-lg"
                          style={`background: hsl(${
                            numHash(cref.target) % 360
                          }, 100%, 80%)`}>{cref.target}</span
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
                        <ActionEditButton onClick={() => {}} />
                        <ActionDeleteButton onClick={() => {}} />
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
