<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let name;
  export let files = [];
  export let onClick = (file) => {};

  $: __active = false;
</script>

<div
  class="py-2 px-2 flex hover:text-blue-500 cursor-pointer focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500 relative"
>
  <button
    class="flex items-center"
    on:click={() => {
      __active = !__active;
    }}
  >
    {name} &nbsp;&nbsp;

    <Icon name="chevron-down" solid class="w-6 h-6 mt-2" />
  </button>

  {#if __active}
    <div
      class="absolute z-50 -right-4 top-8 py-2 bg-white rounded shadow dark:bg-gray-800 w-36 transition-all ease-out duration-500"
    >
      {#each files as file}
        <button
          on:click={() => {
            __active = false;
            onClick(file);
          }}
          class="block py-2 w-full text-sm text-gray-600 transition-colors duration-200 transform hover:bg-blue-200"
        >
          {file}
        </button>
      {/each}
    </div>
  {/if}
</div>
