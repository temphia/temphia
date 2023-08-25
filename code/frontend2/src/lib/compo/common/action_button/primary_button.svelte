<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  export let label = "";
  export let icon = "";
  export let onClick;

  let loading = false;

  const handler = async () => {
    loading = true;
    const r = onClick();
    if (r instanceof Promise) {
      await r;
    }
    loading = false;
  };
</script>

<button
  on:click={handler}
  class="p-1 rounded bg-green-500 shadow hover:bg-green-900 flex text-white"
>
  {#if loading}
    <svg
      class="animate-spin h-5 w-5 text-white p-0.5"
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
  {:else}
    <Icon name={icon} class="h-6 w-6" solid />
  {/if}

  {label}
</button>
