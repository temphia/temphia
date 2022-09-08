<script lang="ts">
  import { getContext } from "svelte";
  export let onNewName;
  const { close } = getContext("simple-modal");

  let value = "";
  const do_action = async () => {
    const resp = await onNewName(value);
    close();
  };
</script>

<div class="w-full h-full">
  <div class="text-center">
    <h2 class="text-xl font-bold text-gray-900">New Folder!</h2>
  </div>
  <div class="mt-4 space-y-3">
    <div class="grid grid-cols-1 space-y-2">
      <label class="text-sm font-bold text-gray-500 tracking-wide">Name</label>
      <input
        bind:value
        class="text-base p-2 border border-gray-300 rounded-lg focus:outline-none focus:border-indigo-500"
        type="text"
        placeholder="new_folder"
      />
    </div>
    <button
      on:click={do_action}
      class="p-2 rounded cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700"
      >Create</button
    >
  </div>
</div>
