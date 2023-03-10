<script lang="ts">
  import { ColumnResize } from "./column_resize";
  import VirtualList from "./_virtual_list.svelte";
  import { createEventDispatcher } from "svelte";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import Cicon from "../cicon/cicon.svelte";

  export let columns_index: { [_: string]: object };
  export let columns: string[];
  export let rows: number[];
  export let rows_index: { [_: string]: object };
  export let main_column: string;
  export let selected_rows = [];

  const dispatch = createEventDispatcher();
  const onPageButtom = () => dispatch("on_page_buttom");
  const onPageTop = () => dispatch("on_page_top");
  const rowClick = (payload) => dispatch("on_row_click", payload);

  const rowToggleSelect = (payload) =>
    dispatch("on_row_toggle_select", payload);

  let left_ref;
  let head_ref;

  const flipCSS = (index) => (index % 2 === 1 ? "gray" : "");
  const DEFAULT_WIDTH = 15;
  const column_resize = ColumnResize(DEFAULT_WIDTH);

  const scrollHandle = (sTop, sTopMax, sLeft) => {
    head_ref.scrollLeft = sLeft;
    left_ref.scrollTop = sTop;
    if (sTopMax === sTop) {
      onPageButtom();
    } else if (sTop === 0) {
      onPageTop();
    }
  };

  let start;
  let end;

  let heightClass = "h-12";

  $: re_render = 0;
</script>

<!-- left start -->
<div class="flex w-full" style="height:calc(100vh - 7em);">
  <div class="w-40 h-full border border-gray-300">
    <div class="flex flex-col gap-1 h-full block">
      <div
        class="h-10 border-b border-gray-300 flex justify-center text-sm bg-gray-50"
      >
        {#if main_column}
          <div class="p-2">{main_column}</div>
        {/if}
      </div>
      <div class="h-full border-collapse overflow-hidden" bind:this={left_ref}>
        {#each rows as row, index}
          <div
            class="{heightClass} cursor-pointer flex border-t border-r border-l border-gray-200 bg-{flipCSS(
              index
            )}-50"
          >
            <div class="pl-1 flex items-center">
              <label
                class="text-teal-500 inline-flex justify-between items-center hover:bg-gray-200 px-2 py-2 rounded-lg cursor-pointer"
              >
                <input
                  checked={selected_rows.includes(row)}
                  on:click={() => rowToggleSelect(row)}
                  type="checkbox"
                  class="form-checkbox rowCheckbox focus:outline-none focus:shadow-outline"
                />
              </label>
            </div>

            <div
              class="grow flex justify-between text-white hover:text-gray-600"
              on:click={() => rowClick(row)}
            >
              <div class="p-2 text-thin text-gray-700 ">
                {row}
              </div>
              {#if main_column}
                <div class="p-2 text-thin text-gray-700 ">
                  {rows_index[row][main_column]}
                </div>
              {/if}
              <div class="p-2">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  class="h-6 w-6"
                  fill="none"
                  viewBox="0 0 24 24"
                  stroke="currentColor"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M4 8V4m0 0h4M4 4l5 5m11-1V4m0 0h-4m4 0l-5 5M4 16v4m0 0h4m-4 0l5-5m11 5l-5-5m5 5v-4m0 4h-4"
                  />
                </svg>
              </div>
            </div>
          </div>
        {/each}
        <div class="h-20 block" />
      </div>
    </div>
  </div>
  <!-- end left -->

  <div class="w-full h-full overflow-hidden border border-gray-300">
    <!-- right start -->
    <div class="flex flex-col gap-1 h-full">
      <div
        class="h-10 bg-gray-50 border-b border-gray-300 overflow-hidden w-full flex justify-between"
        bind:this={head_ref}
      >
        <!-- COLUMNS -->
        {#each columns as col}
        {@const coldata = columns_index[col]}
          {#if main_column !== col}
            <div
              class="flex justify-center font-sans align-middle"
              style="min-width:{$column_resize[col] || DEFAULT_WIDTH}em;"
            >
              <button
                class="menu font-thin text-gray-800 focus:outline-none focus:shadow-solid inline-flex"
              >
              <Cicon ctype={coldata["ctype"]} classes="h-5 w-5 pt-1" />
                {coldata["name"] || coldata["slug"] || ""}
              </button>

              <span class="p-2">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="10"
                  height="10"
                  fill="currentColor"
                  class="bi bi-caret-down-fill"
                  viewBox="0 0 16 16"
                >
                  <path
                    d="M7.247 11.14 2.451 5.658C1.885 5.013 2.345 4 3.204 4h9.592a1 1 0 0 1 .753 1.659l-4.796 5.48a1 1 0 0 1-1.506 0z"
                  />
                </svg>
              </span>

              <div class="relative h-full" style="right: -33%;">
                <div
                  class="h-full w-1 bg-white hover:bg-blue-400"
                  style="cursor:col-resize;"
                  on:mousedown={column_resize.getHandler(col)}
                />
              </div>
            </div>
          {/if}
        {/each}
      </div>

      <!-- row body / main cell contents -->
      <VirtualList
        items={rows}
        let:item
        let:itemIndex
        bind:start
        bind:end
        onScroll={scrollHandle}
      >
        <div
          class="{heightClass} border-b border-gray-200 bg-{flipCSS(
            itemIndex
          )}-50"
          data-row={item || 0}
        >
          <div class="flex justify-between">
            {#each columns as col}
              {#if main_column !== col}
                <div
                  data-col={col}
                  data-row={item.id || 0}
                  style="min-width:{$column_resize[col] || 15}em;"
                  class="{heightClass} overflow-hidden flex justify-center cursor-pointer bg-{flipCSS(
                    itemIndex
                  )}-50"
                >
                  <slot name="cell" row={item} column={col}>
                    <div
                      class="text-gray-700 truncate overflow-hidden text-sm p-1"
                    >
                      {rows_index[item][col] || ""}
                    </div>
                  </slot>
                </div>
              {/if}
            {/each}
          </div>
        </div>
      </VirtualList>
    </div>
  </div>
  <!-- end right  -->
</div>
