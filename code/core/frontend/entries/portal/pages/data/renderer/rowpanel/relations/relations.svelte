<script lang="ts">
  import { getContext } from "svelte";

  import type {
    PortalApp,
    DataTableService,
  } from "../../../../../../../lib/app/portal";
  import Embed from "../../embed/embed.svelte";

  export let reverse_ref_column: object[];
  export let manager: DataTableService;
  export let row: object;

  const app: PortalApp = getContext("__app__");

  let selected = false;
  let loading = false;
  let data = {};
  const load = async (column: object) => {
    selected = true;
    loading = true;
    const resp = await manager.rev_ref_load(
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
