<script lang="ts">
  import ActionSaveButton from "../../../../../xcompo/common/action_button/action_save_button.svelte";
  import ViewItem from "../../../data/tableui/core/view/_view_item.svelte";
  import { KvEditor } from "../../core";

  export let columns;
  export let data;
  export let onSave = (_data) => {};

  let name = data["name"];
  let filter_conds = data["filter_conds"] || [];
  let count = data["count"] || 0;
</script>

<div class="flex flex-col px-2 py-3 mt-2 border-b">
  <h2 class="inline-block text-lg  text-slate-800 mb-1">Name</h2>

  <input
    type="text"
    bind:value={name}
    placeholder="a word.."
    class="w-full h-10 px-4 text-sm text-gray-700 bg-white border border-gray-300 rounded-lg duration-300 focus:border-teal-500 focus:outline-none focus:ring focus:ring-primary focus:ring-opacity-40"
  />
</div>

<div class="flex flex-col px-2 py-3 text-slate-600 border-b">
  <h2 class="inline-block text-lg  text-slate-800 mb-1">Filter Conditions</h2>

  <div class="w-full p-2">
    <ViewItem
      {columns}
      bind:filter_conds
      onModify={(_fc) => {
        data["filter_conds"] = _fc;
      }}
    />
  </div>
</div>

<div class="flex flex-col px-2 py-3 mt-2 border-b">
  <h2 class="inline-block text-lg text-slate-800 mb-1">Select Columns</h2>

  <div class="flex flex-wrap text-gray-700 gap-1">
    {#each columns as col}
      <label class="p-1 border bg-red-50 rounded">
        {col.name}
        <input type="checkbox" />
      </label>
    {/each}
  </div>
</div>

<div class="flex flex-col px-2 py-3 mt-2 border-b">
  <h2 class="inline-block text-xl text-slate-800 mb-1">Fetch Row Count</h2>
  <input type="number" class="border rounded w-20" bind:value={count} />
</div>

<div class="flex flex-col px-2 py-3 border-b">
  <h2 class="inline-block text-lg text-slate-800 mb-1">View Tags</h2>
  <KvEditor data={{}} />
</div>

<div class="flex justify-end mt-2">
  <ActionSaveButton
    name="save"
    icon_name="inbox-in"
    onClick={() =>
      onSave({
        name,
        filter_conds,
        count,
      })}
  />
</div>
