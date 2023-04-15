<script lang="ts">
  import type { SheetService } from "../../../../../services/data";

  export let service: SheetService;
  export let rid;

  const state = service.state;
  const refcols = $state.ref_columns;

  const sheets = service.group.sheets;

  const sheetnames = $sheets.reduce((prev, curr) => {
    prev[curr.__id] = curr.name;

    return prev;
  }, {});

  let selected;
  let loading = false;
  let data = {};

  const load = async (refsheet, refcol) => {
    console.log("@columns_refs", refcols);
    const resp = await service.get_relations(rid, refsheet, refcol);
    if (resp.ok) {
      data = resp.data;
    } else {
      data = {};
    }
  };

  $: console.log("@data", data)

</script>

<div class="w-full p-2 flex flex-col relative">
  <div class="flex justify-end p-1 gap-1">
    <select class="p-1 bg-gray-50 border">
      <option>Select Sheet</option>
      {#each refcols as column}
        <option on:click={() => load(column.sheetid, column.__id)}>
          #{sheetnames[column.sheetid]}|>
          {column.name}
        </option>
      {/each}
    </select>

    <div>
      {#if selected && !loading}
        <button
          class="p-1 text-white bg-blue-500 rounded hover:scale-110"
          on:click={() => {}}>Follow</button
        >
      {/if}
    </div>
  </div>

  {#if selected && loading}
    <div>Loading...</div>
  {:else if selected}
    {#if data["rows"]}
      <div>TODO SHOW DATA</div>
    {/if}
  {/if}
</div>
