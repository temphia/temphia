<script lang="ts">
  export let name = "";
  export let sheet_id = 0;
  export let Confirm = () => {};
  export let Deny = () => {};
</script>

<div class="flex flex-col h-full">
  <div class="mb-4">
    <h1 class="block mb-2 text-2xl text-gray-700">
      Are you sure you want to delete the sheet  ?
    </h1>
  </div>

  <div>
    [ {name} / {sheet_id}]
  </div>

  <div class="flex justify-end gap-1 mt-10 border-t">
    <button
      on:click={Confirm}
      class="px-2 py-1 font-bold text-white bg-blue-500 hover:bg-blue-700 focus:outline-none focus:shadow-outline rounded mt-1 "
      >Yes</button
    >
    <button
      on:click={Deny}
      class="px-2 py-1 font-bold text-white bg-blue-300 hover:bg-blue-400 focus:outline-none focus:shadow-outline rounded mt-1 "
      >No</button
    >
  </div>
</div>
