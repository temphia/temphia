<script lang="ts">
  import type { SheetService } from "../../../../../services/data";
  import SheetInner from "../../_sheet_inner.svelte";
  import { formatRefCells } from "../../../../../services/data/sheet/format";
  export let service: SheetService;
  export let gotoSiblingSheet;
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
  let data: object = {};

  const load = async (refsheet, refcol) => {
    loading = true;
    selected = refsheet;
    const resp = await service.get_relations(rid, refsheet, refcol);
    if (resp.ok) {
      data = formatRefCells(resp.data);
      loading = false;
    }
  };

  $: console.log("@data", data);
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
  </div>

  {#if selected && loading}
    <div>Loading...</div>
  {:else if selected}
    {#if data}
      <div class="p-1 border rounded shadow h-64 mt-2 overflow-auto">
        <SheetInner
          editable={false}
          cells={data["cells"] || {}}
          columns={data["columns"] || []}
          rows={data["rows"] || []}
          selected_rows={[]}
          pick_label="goto"
          on:pick_row={({ detail }) => {
            gotoSiblingSheet &&
              gotoSiblingSheet(selected, Number(detail["__id"]));
            service.close_big_modal();
          }}
        />
      </div>
    {/if}
  {/if}
</div>
