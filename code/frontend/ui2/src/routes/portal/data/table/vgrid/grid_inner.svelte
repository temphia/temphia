<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import Cicon from "../core/cicon/cicon.svelte";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import * as cf from "$lib/services/data/table/column";
  import type { FolderTktAPI } from "$lib/services/apiv2";
  import type { MarkColorType } from "$lib/services/data";
  import UserAvatar from "../../sheet/field/_user_avatar.svelte";

  export let columns_index: { [_: string]: object };
  export let columns: string[];
  export let rows: number[];
  export let rows_index: { [_: string]: object };
  export let selected_rows = [];
  export let folder_api: FolderTktAPI;
  export let marked_rows: { [_: number]: MarkColorType };
  export let profile_generator: (string: any) => string;

  const dispatch = createEventDispatcher();
  const onPageButtom = () => dispatch("on_page_buttom");
  const onPageTop = () => dispatch("on_page_top");
  const rowClick = (payload) => dispatch("on_row_click", payload);

  const rowToggleSelect = (payload) =>
    dispatch("on_row_toggle_select", payload);

  const flipCSS = (index) => (index % 2 === 1 ? "gray" : "");

  const scrollHandle = (ev) => {
    const sTop = ev.target.scrollTop;
    const sTopMax = ev.target.scrollTopMax;

    if (sTopMax === sTop) {
      onPageButtom();
    } else if (sTop === 0) {
      onPageTop();
    }
  };
</script>

<!-- left start -->
<div
  class="w-full overflow-scroll"
  style="height:calc(100vh - 7em);"
  on:scroll={scrollHandle}
>
  <table
    class="border-collapse table-auto w-full whitespace-no-wrap bg-white table-striped relative"
  >
    <thead class="text-gray-600 border-gray-200 bg-gray-100 h-12">
      <tr class="text-left">
        <th class="py-1 px-3 sticky top-0 border-b w-20 bg-gray-100"> # </th>

        {#each columns as col}
          {@const coldata = columns_index[col]}
          <th
            class="sticky top-0 border-b px-6 py-1 font-bold tracking-wider uppercase text-base text-gray-700 bg-gray-100"
          >
            <button class="inline-flex hover:bg-blue-200 rounded px-1">
              <Cicon ctype={coldata["ctype"]} classes="h-5 w-5 pt-1" />
              {coldata["name"] || coldata["slug"] || ""}
            </button>
          </th>
        {/each}

        <th />
      </tr>
    </thead>

    <tbody>
      {#each rows as item, ridx}
        <tr id="data-table-row-{item}" class="text-left bg-{flipCSS(ridx)}-100">
          <td class="border-dashed border-t border-gray-200 px-2">
            <label
              class="text-teal-500 inline-flex gap-1 justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer bg-{marked_rows[
                item
              ]}-100"
            >
              <input
                type="checkbox"
                checked={selected_rows.includes(item)}
                on:click={() => rowToggleSelect(item)}
                class="form-checkbox rowCheckbox focus:outline-none focus:shadow-outline"
              />
              <span class="text-xs text-gray-500">{item}</span>
            </label>
          </td>

          {#each columns as col}
            {@const coldata = columns_index[col]}
            {@const ctype = coldata["ctype"]}
            {@const celldata = (rows_index[item] || {})[col]}

            <td class="border-dashed border-t border-gray-200 p-1">
              <div class="text-gray-700 truncate overflow-hidden text-sm p-1">
                {#if coldata["ref_type"]}
                  <div
                    class="inline-flex bg-yellow-50 text-gray-600 rounded px-1"
                  >
                    <span>Ref</span>
                    <span class="font-semibold text-xs text-gray-800 ml-1"
                      >{celldata}</span
                    >
                  </div>
                {:else if (ctype === cf.CtypeMultiFile || ctype === cf.CtypeFile) && celldata}
                  <div class="flex gap-1">
                    {#each celldata.split(",") as cd}
                      <img
                        class="h-8 w-auto"
                        src={folder_api && folder_api.getFilePreviewUrl(cd)}
                        alt=""
                      />
                    {/each}
                  </div>
                {:else if ctype === cf.CtypeDateTime && celldata}
                  <span class="underline text-blue-700"
                    >{new Date(celldata).toLocaleDateString()}</span
                  >
                {:else if ctype === cf.CtypeCheckBox}
                  {#if celldata === true}
                    <Icon name="check" class="w-6 h-6 text-green-500" />
                  {:else if celldata === false}
                    <Icon name="x" class="w-6 h-6 text-red-500" />
                  {/if}
                {:else if ctype === cf.CtypeLocation}
                  {#if celldata}
                    <div class="flex gap-1">
                      <span class="bg-yellow-100 rounded p-0.5 text-gray-600">
                        Lat
                        <strong class="font-semibold text-gray-700"
                          >{celldata[0].toFixed(3)}
                        </strong>
                      </span>

                      <span class="bg-yellow-100 rounded p-0.5 text-gray-600">
                        Long
                        <strong class="font-semibold text-gray-700"
                          >{celldata[1].toFixed(3)}</strong
                        >
                      </span>
                    </div>
                  {/if}
                {:else if ctype === cf.CtypeSingleUser || ctype === cf.CtypeMultiUser}
                  {#if celldata}
                    <div class="inline-flex gap-1">
                      {#each celldata.split(",") as cd}
                        <div
                          class="p-0.5 rounded bg-gray-50 flex border gap-0.5 text-xs items-center"
                        >
                          <UserAvatar
                            name={cd}
                            url={profile_generator && profile_generator(cd)}
                          />
                          <span>{cd}</span>
                        </div>
                      {/each}
                    </div>
                  {/if}
                {:else}
                  {celldata || ""}
                {/if}
              </div>
            </td>
          {/each}

          <td>
            <button
              class="underline text-blue-600"
              on:click={() => rowClick(item)}
            >
              edit
            </button>
          </td>
        </tr>{/each}
    </tbody>
  </table>
</div>

<div class="hidden w-min md:w-full" />

<style>
  th,
  td {
    max-width: 20rem;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
  }
</style>
