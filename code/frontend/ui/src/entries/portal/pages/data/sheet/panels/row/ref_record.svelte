<script lang="ts">
  import type { SheetService } from "../../../../../services/data";

  export let service: SheetService;
  export let onYes = () => {};
  export let sheet_name: string;
  export let sheet_id: number;
  export let row_id: string;
  export let row_name: string;
</script>

<div class="flex flex-col h-full">
  <div class="mb-4">
    <h1 class="block mb-2 text-2xl text-gray-700">
      Do you want to goto record ?
    </h1>
  </div>
  <div>[ {`${sheet_name} ${sheet_id}`} / {` ${row_name} ${row_id}`} ]</div>

  <div class="flex justify-end gap-1 mt-10 border-t">
    <button
      on:click={onYes}
      class="px-2 py-1 font-bold text-white bg-blue-500 hover:bg-blue-700 focus:outline-none focus:shadow-outline rounded mt-1"
      >Yes</button
    >
    <button
      on:click={() => {
        service.close_small_modal();
      }}
      class="px-2 py-1 font-bold text-white bg-blue-300 hover:bg-blue-400 focus:outline-none focus:shadow-outline rounded mt-1"
      >No</button
    >
  </div>
</div>
