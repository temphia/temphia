<script lang="ts">
  export let color: string;
  export let name: string;
  export let left: boolean;
</script>

<div
  class="h-4 w-10 absolute {left
    ? '-left-2'
    : '-right-2' } bottom-1 text-{color}-300 border rounded flex bg-white text-sm justify-between items-center"
>
  {#if !left}
    {name}
  {/if}

  <div class="w-2 h-2 rounded-full bg-{color}-400 hover:bg-{color}-400" />
  {#if left}
    {name}
  {/if}
</div>
