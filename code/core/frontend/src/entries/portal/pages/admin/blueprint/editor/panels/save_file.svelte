<script lang="ts">
  import { getContext } from "svelte";

  export let finalSave;
  export let filename;
  export let beditor;

  const { close } = getContext("simple-modal");
</script>

<div class="flex flex-col items-start">
  <div class="flex items-center w-full">
    <div class="text-gray-900 font-medium text-lg">Save File!</div>
    <svg
      class="ml-auto fill-current text-gray-700 w-6 h-6 cursor-pointer"
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 18 18"
    >
      <path
        d="M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z"
      />
    </svg>
  </div>
  <hr />
  <div class="">Do you really want save File [{filename}] ?</div>
  <hr />
  <div class="ml-auto">
    <button
      on:click={finalSave}
      class="bg-blue-500 hover:bg-blue-700 text-white font-bold py-2 px-4 rounded"
    >
      Yes
    </button>
    <button
      on:click={close}
      class="bg-transparent hover:bg-gray-500 text-blue-700 font-semibold hover:text-white py-2 px-4 border border-blue-500 hover:border-transparent rounded"
    >
      No
    </button>
  </div>
</div>
