<script lang="ts">
  import type { DataTableService } from "../../../../app";
  import Inner from "./_inner.svelte";

  import Layout from "./_layout.svelte";

  export let show = false;

  export let manager: DataTableService;
  export let columns: object[];

  let rootstate = manager.store.states;
  let navstore = manager.navStore;
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
    view_name = v.name || ""
  };
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
    <div class="p-2">
      <label for="">Views</label>
      <select class="p-1 rounded">
        <option value="__index_0" />
        {#each $rootstate[manager.dtable].views || [] as v}
          <option on:click={onViewChange(v)}>{v["name"] || ""}</option>
        {/each}
      </select>
    </div>

    <button
      on:click={() =>
        manager.applyView(view_name, {
          ...getViewData(),
        })}
      class="bg-blue-400 hover:bg-blue-600 w-20 text-white text-sm rounded"
      >Apply</button
    >
  </div>
</Layout>
