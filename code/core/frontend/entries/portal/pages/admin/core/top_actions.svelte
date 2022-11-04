<script lang="ts">
  export let actions: { [_: string]: () => void } = {};
</script>

<div class="flex justify-end p-2">
  {#each Object.entries(actions) as [action, aFn]}
    <button
      on:click={aFn}
      class="px-2 py-1 rounded-full bg-white hover:text-white hover:bg-slate-500 border border-slate-600"
      >{action}</button
    >
  {/each}
</div>
