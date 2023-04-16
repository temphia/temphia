<script>
  import Ceditor from "../../../../xcompo/ceditor/ceditor.svelte";
  import Paramform from "./paramform/paramform.svelte";

  let editor;

  let show = true;

  let code = ``;
</script>

<div class="rounded bg-white p-2">
  <div class="flex flex-wrap justify-end text-sm items-center gap-1">
    <label>
      Script

      <select class="p-1 rounded w-32">
        <option>Do Xyz</option>
        <option>Do Xyz</option>
      </select>
    </label>
    <label>
      <input bind:checked={show} type="checkbox" class="" />
    </label>
  </div>

  {#if show}
    <div class="p-1 flex-grow">
      <Ceditor bind:editor {code} />
    </div>
  {/if}

  <Paramform />

  <div class="flex flex-wrap justify-end text-sm items-center">
    <button
      class="text-white bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded text-sm px-4 py-2 flex"
    >
      Search</button
    >
  </div>
</div>
