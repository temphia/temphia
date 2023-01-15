<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { setContext } from "svelte";
  import type { TableService } from "../../../../../services/data";
  import type { ViewModal } from "./view";

  import Inner from "./_inner.svelte";

  import Layout from "./_layout.svelte";

  export let show = false;

  export let manager: TableService;
  export let columns: object[];
  export let view_modal: ViewModal;

  let rootstate = manager.state.data_store;
  let navstore = manager.state.nav_store;

  let getViewData;

  $: _view = {
    count: 0,
    filter_conds: [],
    main_column: "",
    search_term: "",
    selects: [],
    ...$navstore.active_view,
  };

  let view_name = "";

  const onViewChange = (v) => () => {
    _view = { ...v };
    view_name = v.name || "";
  };

  setContext("__data_context__", {
    modal: view_modal,
    table_service: manager,
  });
</script>

<Layout bind:show>
  <div class="flex-grow flex flex-col h-32 p-2 space-y-1 overflow-y-auto">
    {#key view_name}
      <Inner {columns} data={_view} bind:getViewData />
    {/key}
  </div>

  <div
    class="flex-shrink h-12 w-full flex justify-between gap-x-1 mt-2 pt-2 border-t"
  >
    <div class="p-1 flex gap-2">
      <label class="uppercase text-gray-700" for="">Views</label>
      <select class="p-1 rounded w-36">
        <option value="__index_0" />
        {#each $rootstate.views || [] as v}
          <option on:click={onViewChange(v)}>{v["name"] || ""}</option>
        {/each}
      </select>

      <button class="p-1 rounded bg-gray-50 hover:bg-gray-200">
        <Icon name="cog" class="w-5 h-5" />
      </button>
    </div>

    <button
      on:click={() =>
        manager.apply_view(view_name, {
          ...getViewData(),
        })}
      class="bg-blue-400 hover:bg-blue-600 w-20 text-white text-sm rounded"
      >Apply</button
    >
  </div>
</Layout>
