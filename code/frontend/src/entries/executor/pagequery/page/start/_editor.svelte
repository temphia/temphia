<script lang="ts">
  import Ceditor from "../../../../xcompo/ceditor/ceditor.svelte";
  import type { LoadResponse } from "../../service";

  export let data: LoadResponse;
  export const getCodeValue = () => {
    if (!editor || !modified) {
      return "";
    }
    return editor.getValue();
  };
  
  export let current = data.first_stage;

  let editor;
  let modified = false;

  $: code = (data.stages[current] || {})["script"] || "";
</script>

<div class="p-1 flex-grow">
  <Ceditor
    bind:editor
    {code}
    on:change={() => {
      modified = true;
    }}
  />
</div>

<div class="flex items-center justify-end text-gray-700 gap-2">
  <label for="stage"> Templates </label>
  <select
    id="stage"
    class="p-1 rounded w-40 border"
    value={data["first_stage"]}
    on:change={(ev) => {
      current = ev.target["value"];
      modified = false;
    }}
  >
    {#each Object.entries(data["stages"]) as [skey, sval]}
      <option value={skey}>
        {skey}
      </option>
    {/each}
  </select>
</div>
