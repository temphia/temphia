<script lang="ts">
  import type { TableService } from "../../../../../../services/data";


  import Embed from "../../embed/embed.svelte";

  export let reverse_ref_column: object[];
  export let table_service: TableService;
  export let row: object;

  const row_service = table_service.get_row_service();

  let selected = false;
  let loading = false;
  let data = {};
  const load = async (column: object) => {
    selected = true;
    loading = true;
    const resp = await row_service.rev_ref_load(
      column["table_id"],
      column["slug"],
      row["__id"]
    );

    data = resp.data;
    loading = false;
  };
</script>

<div class="w-full p-2 flex flex-col relative">
  <div class="flex justify-end p-1 gap-1">
    <select class="p-1 bg-gray-50 border">
      <option>Select Table</option>
      {#each reverse_ref_column as column}
        <option on:click={() => load(column)}>
          #{column["table_id"]}|>
          {column["slug"]}
        </option>
      {/each}
    </select>

    <div>
      {#if selected && !loading}
        <button
          class="p-1 text-white bg-blue-500 rounded hover:scale-110"
          on:click={() => {
            // generate filter condition
            // app.navigator.goto_dtable("", "", "", {});
          }}>Follow</button
        >
      {/if}
    </div>
  </div>

  {#if selected && loading}
    <div>Loading...</div>
  {:else if selected}
    {#if data["rows"]}
      <Embed onRowSelect={() => {}} {data} />
    {/if}
  {/if}
</div>
