<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  export let loading = undefined;
  export let error = "";
  export let placeholder = "";
  export let label = "";
  export let value = "";
</script>

<div class="relative h-14">
  <label for="plug-slug" class="block mb-2 text-base text-1.5xl font-medium text-gray-900"
    >{label}</label
  >

  <button class="absolute right-1 top-8">
    {#if loading !== undefined}
      {#if loading}
        <svg
          class="animate-spin h-5 w-5 text-blue-700"
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
        <Icon name="check-circle" class="h-5 w-5 text-green-600" />
      {/if}
    {/if}
  </button>

  <input
    type="text"
    bind:value
    id="plug-slug"
    class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-blue-500 focus:border-blue-500 block w-full p-2.5 "
    {placeholder}
  />

  <p class="italic text-red-500 text-sm p-1">{error}</p>
</div>
