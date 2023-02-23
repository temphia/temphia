<script lang="ts">
  import { ColumnResize } from "./column_resize";
  import { createEventDispatcher } from "svelte";
  import Cicon from "../cicon/cicon.svelte";
  import * as f from "../fields/field";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { FolderTktAPI } from "../../../../../../../lib/apiv2";

  export let columns_index: { [_: string]: object };
  export let columns: string[];
  export let rows: number[];
  export let rows_index: { [_: string]: object };
  export let selected_rows = [];
  export let folder_api: FolderTktAPI;

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
            class="sticky top-0 border-b  px-6 py-1 font-bold tracking-wider uppercase text-base text-gray-700 bg-gray-100"
          >
            <button class="inline-flex hover:bg-blue-200 rounded px-1">
              <Cicon ctype={coldata["ctype"]} classes="h-5 w-5 pt-1" />
              {coldata["name"] || coldata["slug"] || ""}
            </button>
          </th>
        {/each}
      </tr>
    </thead>

    <tbody>
      {#each rows as item, ridx}
        <tr class="text-left {flipCSS(ridx)}">
          <td class="border-dashed border-t border-gray-200 px-2">
            <label
              class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer"
            >
              <input
                type="checkbox"
                checked={selected_rows.includes(item)}
                on:click={() => rowToggleSelect(item)}
                class="form-checkbox rowCheckbox focus:outline-none focus:shadow-outline"
              />
            </label>

            <button
              class="underline text-blue-500 "
              on:click={() => rowClick(item)}
            >
              <span class="text-xs text-gray-500">{item}</span>
            </button>
          </td>

          {#each columns as col}
            {@const coldata = columns_index[col]}
            {@const ctype = coldata["ctype"]}
            {@const celldata = rows_index[item][col]}

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
                {:else if (ctype === f.CtypeMultiFile || ctype === f.CtypeFile) && celldata}
                  {#each celldata.split(",") as cd}
                    <img
                      class="h-8 w-auto"
                      src={folder_api && folder_api.getFilePreviewUrl(cd)}
                      alt=""
                    />
                  {/each}
                {:else if ctype === f.CtypeDateTime && celldata}
                  <span class="underline text-blue-700"
                    >{new Date(celldata).toLocaleDateString()}</span
                  >
                {:else if ctype === f.CtypeCheckBox}
                  {#if celldata === true}
                    <Icon name="check" class="w-6 h-6 text-green-500" />
                  {:else if celldata === false}
                    <Icon name="x" class="w-6 h-6 text-red-500" />
                  {/if}
                {:else if (ctype === f.CtypeSingleUser || ctype === f.CtypeMultiUser) && celldata}
                  <div class="inline-flex gap-1">
                    {#each celldata.split(",") as cd}
                      <div class="flex">
                        <img
                          alt=""
                          src="/z/assets/static/default_user_profile.png"
                          class="w-6 h-6 p-1 rounded-full bg-green-50"
                        />
                        <!-- fixme => profile image -->
                        <span
                          class="underline text-green-800 bg-green-50 rounded"
                          >{cd}</span
                        >
                      </div>
                    {/each}
                  </div>
                {:else}
                  {celldata || ""}
                {/if}
              </div>
            </td>
          {/each}
        </tr>{/each}
    </tbody>
  </table>
</div>

<div class="hidden w-min md:w-full" />
