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
  class="px-2 py-1 font-bold text-white bg-{color}-500 hover:bg-{color}-700 focus:outline-none focus:shadow-outline rounded mt-1"
>
  {#if loading}
    <svg
      class="animate-spin -ml-1 mr-3 h-5 w-5 text-white"
      fill="none"
      viewBox="0 0 24 24"
    >
      <circle
        class="opacity-25"
        cx="12"
        cy="12"
        r="10"
        stroke="currentColor"
        stroke-width="4"
      />
      <path
        class="opacity-75"
        fill="currentColor"
        d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
      />
    </svg>
  {/if}

  {name}
</button>
