<script lang="ts">
  import { CEditor } from "../../../core";
  export let invoker_code = "{}";
  export let invoker_name = "";

  let editor;
</script>

<div class="w-full h-full flex flex-col">
  <div class="p-1 flex justify-between h-14">
    <div class="p-2">
      <label class="uppercase text-gray-700 font-semibold text-lg">
        Name

        <input
          type="text"
          bind:value={invoker_name}
          class="p-2 h-full rounded border w-56"
        />
      </label>
    </div>
  </div>

  <div class="flex flex-col h-96 flex-grow">
    <CEditor
      code={invoker_code}
      bind:editor
      container_style={"min-height:20rem;"}
    />

    <div class="flex justify-end mt-4">
      <button
        class="p-1 rounded bg-blue-500 shadow hover:bg-blue-900 flex text-white uppercase"
      >
        Set
      </button>
    </div>
  </div>
</div>
