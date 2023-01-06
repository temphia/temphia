<script lang="ts">
  import { getContext } from "svelte";
  import type { EasypageService } from "../service/easypage";
  import NewPage from "./_panels/new_page.svelte";

  export let onSave;

  const service = getContext("__easypage_service__") as EasypageService;

  const addPage = () => {
    service.modal.small_open(NewPage, { onSave });
  };
</script>

<div class="bg-blue-50 w-full h-screen font-sans flex flex-col">
  <div class="flex p-2 justify-between bg-blue-400">
    <div class="flex">
      <h3 class="text-white mr-4">Easypage Adapter Editor</h3>
    </div>

    <div class="flex gap-2">
      <button
        class="bg-blue-200 hover:bg-blue-500 rounded  font-bold text-white text-sm px-2"
        >↻</button
      >
      <button
        on:click={addPage}
        class="bg-blue-200 hover:bg-blue-500 rounded  font-bold text-white text-sm px-2"
        >+</button
      >
    </div>
  </div>

  <slot />
</div>
