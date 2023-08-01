<script>
  $: _mode_expanded = false;
  const click = () => (_mode_expanded = !_mode_expanded);
  export let actions = [{ name: "Xyz", action: () => {} }, { name: "PQR asasa s", action: () => {} }];
</script>

<div class="fixed bottom-20 right-10">
  {#if _mode_expanded}
    <slot>
      <div class="p-1 flex flex-col gap-1">
        {#each actions as action}
          <div class="bg-white shadow-lg rounded-lg py-1 pl-1 px-4">
            <div class="">
              <button
                class="p-0 w-12 h-12 bg-red-500 text-white rounded-full hover:bg-red-700 active:shadow-lg mouse shadow transition ease-in duration-200 focus:outline-none"
              >
                <svg
                  viewBox="0 0 20 20"
                  enable-background="new 0 0 20 20"
                  class="w-6 h-6 inline-block"
                >
                  <path
                    fill="#FFFFFF"
                    d="M16,10c0,0.553-0.048,1-0.601,1H11v4.399C11,15.951,10.553,16,10,16c-0.553,0-1-0.049-1-0.601V11H4.601
                                    C4.049,11,4,10.553,4,10c0-0.553,0.049-1,0.601-1H9V4.601C9,4.048,9.447,4,10,4c0.553,0,1,0.048,1,0.601V9h4.399
                                    C15.952,9,16,9.447,16,10z"
                  />
                </svg>
              </button>
              <span>{action.name}</span>
            </div>
          </div>
        {/each}
      </div>
    </slot>
  {/if}
</div>

<div class="fixed bottom-10 right-10">
  <button
    on:click={click}
    class="p-0 w-14 h-14 bg-red-600 rounded-full hover:bg-red-700 active:shadow-lg mouse shadow transition ease-in duration-200 focus:outline-none"
  >
    <svg
      viewBox="0 0 20 20"
      enable-background="new 0 0 20 20"
      class="w-6 h-6 inline-block"
    >
      {#if _mode_expanded}
        <path
          fill="#FFFFFF"
          d="M4.293 4.293a1 1 0 011.414 0L10 8.586l4.293-4.293a1 1 0 111.414 1.414L11.414 10l4.293 4.293a1 1 0 01-1.414 1.414L10 11.414l-4.293 4.293a1 1 0 01-1.414-1.414L8.586 10 4.293 5.707a1 1 0 010-1.414z"
          clip-rule="evenodd"
        />
      {:else}
        <path
          fill="#FFFFFF"
          d="M16,10c0,0.553-0.048,1-0.601,1H11v4.399C11,15.951,10.553,16,10,16c-0.553,0-1-0.049-1-0.601V11H4.601
                          C4.049,11,4,10.553,4,10c0-0.553,0.049-1,0.601-1H9V4.601C9,4.048,9.447,4,10,4c0.553,0,1,0.048,1,0.601V9h4.399
                          C15.952,9,16,9.447,16,10z"
        />
      {/if}
    </svg>
  </button>
</div>
