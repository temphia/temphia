<script lang="ts">
  export let total_steps: number;
  export let curent_step: number;
</script>

<div class="h-full w-full py-4">
  <div class="m-auto bg-white shadow rounded p-1" style="max-width: 55rem;">
    <div class="flex justify-between gap-2 p-4">
      {#each Array.from({ length: total_steps }) as step}
        {#if (step < curent_step)}
          <div
            class="h-4 bg-green-500 w-full rounded-full hover:bg-green-400 cursor-pointer"
          />
        {:else if step === curent_step}
          <div
            class="h-4 bg-green-400 w-full rounded-full hover:bg-green-300 cursor-pointer"
          />
        {:else}
          <div
            class="h-4 bg-green-100 w-full rounded-full hover:bg-green-300 cursor-pointer"
          />
        {/if}
      {/each}
    </div>
    <slot />
  </div>
</div>
