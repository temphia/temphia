<script lang="ts">
  export let loading = true;
  export let onClick;
  export let name;
  export let color = "blue";
</script>

<button
  on:click={async () => {
    loading = true;
    const resp = onClick();
    if (resp instanceof Promise) {
      await resp;
    }

    loading = false;
  }}
  class="px-2 py-1 font-bold text-white bg-{color}-500 hover:bg-{color}-700 focus:outline-none focus:shadow-outline rounded mt-1 {loading
    ? 'animate-bounce'
    : ''} ">{name}</button
>
