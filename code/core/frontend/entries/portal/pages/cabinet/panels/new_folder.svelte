<script lang="ts">
  export let onNewName;

  let value = "";
</script>

<div class="w-full h-full">
  <div class="text-center">
    <h2 class="text-xl font-bold text-gray-900">New Folder!</h2>
  </div>
  <div class="p-4 flex flex-col gap-2">
    <label class="text-sm font-bold text-gray-500 tracking-wide"
      >Name
      <input
        bind:value
        class="text-base p-2 border border-gray-300 rounded-lg focus:outline-none focus:border-indigo-500 w-full"
        type="text"
        placeholder="new_folder"
      />
    </label>

    <div class="inline-flex justify-end">
      <button
        disabled={value === ""}
        on:click={() => onNewName(value)}
        class="p-2 rounded cursor-pointer uppercase text-sm font-bold bg-blue-500 text-white hover:bg-blue-700 disabled:bg-gray-500"
        >Create</button
      >
    </div>
  </div>
</div>
