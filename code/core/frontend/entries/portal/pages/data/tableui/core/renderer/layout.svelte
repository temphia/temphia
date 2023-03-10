<script lang="ts">
  import ActionBtn from "./_action_button.svelte";
  import Loading from "./_loading.svelte";
  import DOMPurify from "dompurify";
  import { createEventDispatcher } from "svelte";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let hooks: object[];

  export let all_tables: object[];
  export let active_table: string;
  export let actions: object[];
  export let loading: boolean = false;
  export let selected_rows = [];
  export let layout: string = "grid";

  export let rows_loaded_no = 0;
  export let rows_total_no = 0

  const dispatch = createEventDispatcher();
  const onChangeDtable = (payload) => dispatch("on_table_change", payload);
  const newRowClick = (payload) => dispatch("on_new_row", payload);
  const onHookClick = (payload) => dispatch("on_hook_click", payload);
  const onChangeLayout = () => {
    if (layout === "card") {
      dispatch("on_change_to_grid");
    } else {
      dispatch("on_change_to_card");
    }
  };

  const gotoDtable = (dtable) => () => {
    onChangeDtable(dtable);
  };

  $: re_render = 0;

  const _onHook = (hook: object) => () => {
    onHookClick(hook);
  };
</script>

<div class="w-full h-full overflow-x-hidden">
  <div class="m-1 pb-1 pl-1 pr-1 shadow bg-white rounded border">
    <div>
      <!-- TABS  start -->
      <ul
        class="list-reset flex overflow-x-auto border-t-1 ml-4 divide divide-light-blue-400"
      >
        {#each all_tables as tbl}
          <li class="border-0 border-t border-r border-l">
            {#if tbl["slug"] !== active_table}
              <span
                class="bg-gray-100 inline-block border border-gray-300 rounded-t px-1 md:px-2 text-xs md:text-base text-blue-dark font-semibold"
              >
                <button
                  on:click={gotoDtable(tbl["slug"])}
                  class="align-middle rounded h-8 md:h-10">{tbl["name"]}</button
                >
              </span>
            {:else}
              <span
                class="bg-white inline-block p-1 md:px-2 text-xs md:text-base text-blue hover:text-blue-darker font-semibold"
              >
                {tbl["name"]}
              </span>
            {/if}
          </li>
        {/each}
      </ul>
    </div>
    <!-- TABS  end -->

    <div class="rounded-t-lg flex shadow justify-between items-center">
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
              {#if selected_rows.length > 0}
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
              {#if (hook["sub_type"] === "row" && selected_rows.length > 0) || hook["sub_type"] === "table"}
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

      <div class="p-1">
        <button
          on:click={onChangeLayout}
          class="p-1 bg-gray-50 text-gray-700 inline-flex font-light rounded hover:bg-blue-200"
        >
          {#if layout === "card"}
            <Icon name="table" solid class="h-6 w-6 pt-1" />
            Layout
          {:else}
            <Icon name="color-swatch" solid class="h-6 w-6 pt-1" />
            Layout
          {/if}
        </button>
      </div>
    </div>

    <div class="w-full h-full overflow-hidden">
      <slot>Empty slot</slot>

      <div class="flex justify-start p-0.5 text-sm gap-2 bg-slate-100 uppercase">
        <button>
          <Icon name="chevron-double-right" solid class="h-5 w-5 text-slate-700" />
        </button>
        <p class="text-slate-900">
          Rows <span class="text-slate-700 text-base"
            >[ {`${rows_loaded_no}/${rows_total_no}`} ]</span
          >
        </p>
      </div>
    </div>
    <!-- end right  -->
  </div>
</div>

<div class="fixed bottom-4 z-5 right-10 ">
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
