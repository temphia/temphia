<script lang="ts">
  import ActionBtn from "./action_btn.svelte";
  import { ColumnResize } from "./column_resize";
  import VirtualList from "./virtual_list.svelte";
  import Loading from "./_loading.svelte";
  import DOMPurify from "dompurify";

  export let columns_index: { [_: string]: object };
  export let columns: string[];
  export let hooks: object[];

  export let rows: number[];
  export let rows_index: { [_: string]: object };

  export let all_tables: object[];
  export let active_table: string;
  export let actions: object[];
  export let main_column: string;
  export let loading: boolean = false;

  export let onChangeDtable;
  export let onPageButtom;
  export let onPageTop;
  export let rowClick;
  export let newRowClick;
  export let rowToggleSelect;

  export let onHookClick;

  export let selectedRows = [];

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

  const gotoDtable = (dtable) => () => {
    onChangeDtable(dtable);
  };

  let start;
  let end;

  let heightClass = "h-12";
  $: _selected_rows = [];

  $: re_render = 0;

  const _onHook = (hook: object) => () => {
    onHookClick(hook)
  };
</script>

<div class="w-full h-full overflow-x-hidden">
  <div class="m-1 pb-1 pl-1 pr-1 shadow bg-white rounded border">
    <div>
      <!-- TABS  start -->
      <ul class="list-reset flex overflow-x-auto border-t-1 ml-4 divide divide-light-blue-400">
        {#each all_tables as tbl}
          <li class="border-0 border-t-1 border-r-1 border-l-1">
            {#if tbl["slug"] !== active_table}
              <span
                class="bg-gray-50 border-gray-100 inline-block border border-gray-300 rounded-t px-1 md:px-2 text-xs md:text-base text-blue-dark font-semibold"
              >
                <button on:click={gotoDtable(tbl["slug"])} class="align-middle rounded h-8 md:h-10"
                  >{tbl["name"]}</button
                >
              </span>
            {:else}
              <span
                class="bg-white inline-block px-1 md:px-2 text-xs md:text-base text-blue hover:text-blue-darker border-t border-l border-r border-b-0 font-semibold"
              >
                <div class="dropdown">
                  <div class="dropdown-toggle align-middle rounded h-8 md:h-10">{tbl["name"]}</div>
                </div>
              </span>
            {/if}
          </li>
        {/each}
      </ul>
    </div>
    <!-- TABS  end -->

    <div
      class="rounded-t-lg flex flex-col shadow md:flex-row justify-between items-center"
    >
      <!-- TOOLBAR  start -->
      <div class="flex flex-wrap p-1 pr-4 gap-x-1">
        {#key re_render}
          {#each actions as action}
            {#if action["type"] === "normal"}
              <ActionBtn
                {action}
                clx={action["active"] ? "bg-blue-200" : "bg-gray-50"}
              />
            {:else if action["type"] === "contextual"}
              {#if selectedRows.length > 0}
                <ActionBtn
                  {action}
                  clx={action["active"] ? "bg-blue-200" : "bg-blue-50"}
                />
              {/if}
            {/if}
          {/each}
          <div class="h-full w-2" />

          {#each hooks as hook}
            {#if hook["type"] === "data_hook"}
              {#if (hook["sub_type"] === "row" && selectedRows.length > 0) || hook["sub_type"] === "table"}
                <ActionBtn
                  action={{
                    name: hook["name"],
                    action: _onHook(hook),
                    icon: hook["icon"] ? DOMPurify.sanitize(hook["icon"]) : "",
                  }}
                  clx={"bg-green-50"}
                />
              {/if}
            {/if}
          {/each}
        {/key}
      </div>
    </div>
    <!-- TOOLBAR  END -->

    <!-- left start -->
    <div class="flex w-full" style="height:calc(100vh - 6em);">
      <div class="w-80 h-full border border-gray-300">
        <div class="flex flex-col gap-1 h-full block">
          <div
            class="h-10 border-b border-gray-300 flex justify-center text-sm bg-gray-50"
          >
            {#if main_column}
              <div class="p-2">{main_column}</div>
            {/if}
          </div>
          <div
            class="h-full border-collapse overflow-hidden"
            bind:this={left_ref}
          >
            {#each rows as row, index}
              <div
                class="{heightClass} cursor-pointer flex border-t border-r border-l border-gray-200 bg-{flipCSS(
                  index
                )}-50"
              >
                <div class="pl-1 flex items-center">
                  <input
                    checked={selectedRows.includes(row)}
                    type="checkbox"
                    on:click={rowToggleSelect(row)}
                    class="relative peer py-2 text-purple-600 rounded-md focus:ring-0"
                  />
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
              {#if main_column !== col}
                <div
                  class="flex justify-center font-sans align-middle"
                  style="min-width:{$column_resize[col] || DEFAULT_WIDTH}em;"
                >
                  <button
                    class="menu font-thin text-gray-800 focus:outline-none focus:shadow-solid inline-flex"
                  >
                    {columns_index[col]["name"] ||
                      columns_index[col]["slug"] ||
                      ""}
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
                      class="h-full w-1 bg-white hover:bg-blue-400 h-full"
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
  </div>
</div>

<div class="fixed bottom-8 z-5 right-10 ">
  {#if loading}
    <Loading />
  {:else}
    <button
      on:click={newRowClick}
      class="p-0 w-8 h-8 md:w-10 md:h-10 bg-blue-600 rounded-full hover:bg-blue-700 active:shadow-lg mouse shadow transition ease-in duration-200 focus:outline-none"
    >
      <svg
        viewBox="0 0 20 20"
        enable-background="new 0 0 20 20"
        class="w-6 h-6 inline-block"
      >
        <path
          fill="#FFFFFF"
          d="M16,10c0,0.553-0.048,1-0.601,1H11v4.399C11,15.951,10.553,16,10,16c-0.553,0-1-0.049-1-0.601V11H4.601 C4.049,11,4,10.553,4,10c0-0.553,0.049-1,0.601-1H9V4.601C9,4.048,9.447,4,10,4c0.553,0,1,0.048,1,0.601V9h4.399 C15.952,9,16,9.447,16,10z"
        />
      </svg>
    </button>
  {/if}
</div>
