<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import type { LoadResponse } from "../../service";

  export let data: LoadResponse;

  const dispatch = createEventDispatcher();
</script>

<div class="flex items-center justify-end text-gray-700 gap-2">
  <label for="stage"> Templates </label>
  <select
    id="stage"
    class="p-1 rounded w-40 border"
    value={data["first_stage"]}
    on:change={(ev) => {
      dispatch("change_template", ev.target["value"]);
    }}
  >
    {#each Object.entries(data["stages"]) as [skey, sval]}
      <option>
        {skey}
      </option>
    {/each}
  </select>
</div>
