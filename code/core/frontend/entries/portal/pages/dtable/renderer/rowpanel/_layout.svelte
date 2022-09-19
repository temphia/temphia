<script lang="ts">
  export let row_id;
  const MODE_EDIT = 0;
  const MODE_ACTIVITY = 1;
  const MODE_META_EDIT = 2;
  const MODE_RELATIONS = 3;

  const active_css =
    "bg-white inline-block py-1 px-1 md:px-2 text-blue hover:text-blue-darker text-xs md:text-base border-t border-l border-r border-b-0 font-semibold";
  const inactive_css =
    "bg-gray-50 border-gray-100 inline-block border-2 rounded-t text-xs md:text-base py-1 px-1 md:px-2 text-blue-dark font-semibold";

  export let show_editor = false;

  $: _mode = MODE_EDIT;

  const hide = () => {
    show_editor = false;
  };
</script>

{#if show_editor}
  <modal-wrapper>
    <modal-section>
      <div
        on:click={hide}
        class="modal-close absolute top-0 right-0 cursor-pointer flex flex-col items-center mt-4 mr-4 text-white text-sm z-50"
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          width="18"
          height="18"
          viewBox="0 0 18 18"
        >
          <path
            d="M14.53 4.53l-1.06-1.06L9 7.94 4.53 3.47 3.47 4.53 7.94 9l-4.47 4.47 1.06 1.06L9 10.06l4.47 4.47 1.06-1.06L10.06 9z"
          />
        </svg>
      </div>

      <div class="h-full w-full flex flex-col">
        <div class="flex-shrink h-10 w-full pb-1 border-b flex justify-between">
          {#if row_id !== 0}
            <ul
              class="list-reset flex border-t-1 ml-4 divide divide-light-blue-400"
            >
              <li class="border-0 border-t-1 border-r-1 border-l-1">
                <a
                  class={_mode === MODE_EDIT ? active_css : inactive_css}
                  href="#"
                  on:click={() => {
                    _mode = MODE_EDIT;
                  }}
                >
                  Edit
                </a>
                <a
                  class={_mode === MODE_META_EDIT ? active_css : inactive_css}
                  on:click={() => {
                    _mode = MODE_META_EDIT;
                  }}
                  href="#">Properties</a
                >
                <a
                  class={_mode === MODE_RELATIONS ? active_css : inactive_css}
                  on:click={() => {
                    _mode = MODE_RELATIONS;
                  }}
                  href="#">Relations</a
                >
                <a
                  class={_mode === MODE_ACTIVITY ? active_css : inactive_css}
                  on:click={() => {
                    _mode = MODE_ACTIVITY;
                  }}
                  href="#">Activity</a
                >
              </li>
            </ul>
            <div class="pr-8 py-2">
              <h2 class="text-indigo-500 text-sm uppercase font-medium">
                Row {row_id}
              </h2>
            </div>
          {:else}
            <div class="px-3 py-1 uppercase text-gray-800">New Row</div>
          {/if}
        </div>

        {#if _mode === MODE_ACTIVITY}
          <slot name="activity">Empty Activity template</slot>
        {:else if _mode === MODE_EDIT}
          <slot name="edit">Empty Edit template</slot>
        {:else if _mode === MODE_RELATIONS}
          <slot name="relations">Empty Meta template</slot>
        {:else}
          <slot name="meta">Empty Meta template</slot>
        {/if}
      </div>
    </modal-section>
  </modal-wrapper>
{/if}

<style>
  modal-wrapper {
    position: fixed;
    top: 0px;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 7;
    background-color: rgba(0, 0, 0, 0.44);
  }
  modal-section {
    overflow-y: auto;
    position: relative;
    margin: auto;
    margin-top: 1rem;
    height: 98vh;
    max-width: 750px;
    background-color: #fff;
    border-radius: 6px;
    display: block;
    padding: 0.25rem;
  }
</style>
