<script lang="ts">
  import type { SheetService } from "../../../../../services/data";
  import { LoadingSpinner } from "../../../../admin/core";
  import HistoryCard from "./_history_card.svelte";

  export let service: SheetService;
  export let rid;

  let datas = [];

  let loading = true;

  const load = async () => {
    loading = true;
    const resp = await service.api.get_row_history(service.sheetid, rid);
    if (!resp.ok) {
      return;
    }

    datas = resp.data;
    loading = false;
  };

  load();

  let open;
</script>

<div class="flex flex-col">
  {#if loading}
    <LoadingSpinner />
  {:else}
    {#each datas as data}
      <HistoryCard
        {data}
        onClick={() => {
          open = data["id"];
        }}
        expanded={open == data["id"]}
      />
    {/each}
  {/if}
</div>
