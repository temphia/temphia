<script lang="ts">
  $: _code = false;
</script>

<div class="w-full h-full">
  <div class="flex justify-center border bg-white rounded p-2">
    <div class="inline-flex">
      <button
        on:click={() => {
          _code = false;
        }}
        class="text-sm hover:bg-gray-300 text-gray-800 border font-semibold py-2 px-4 rounded-l {_code
          ? 'bg-gray-50'
          : 'bg-gray-200'}"
      >
        View
      </button>
      <button
        on:click={() => {
          _code = true;
        }}
        class="text-sm bg-gray-50 hover:bg-gray-300 text-gray-800 border font-semibold py-2 px-4 rounded-r {_code
          ? 'bg-gray-200'
          : 'bg-gray-50'}"
      >
        Raw
      </button>
    </div>
  </div>

  {#if _code}
    <slot name="code" />
  {:else}
    <slot name="view" />
  {/if}
</div>


