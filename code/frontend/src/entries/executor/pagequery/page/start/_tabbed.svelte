<script>
  export let modes = [];
  export let mode = modes[0];
  export let onChangeMode = undefined
</script>

<div class="flex flex-row border flex-nowrap overflow-auto mt-1">
  {#each modes as m}
    {#if mode == m}
      <button
        class="p-2 hover:text-blue-500 focus:outline-none text-blue-500 border-b-2 font-medium border-blue-500 inline-flex"
        >{mode}
      </button>
    {:else}
      <button
        on:click={() => {
          mode = m;
          onChangeMode && onChangeMode(m)
        }}
        class="text-gray-600 p-2 block hover:text-blue-500 focus:outline-none"
        >{m}</button
      >
    {/if}
  {/each}
</div>
