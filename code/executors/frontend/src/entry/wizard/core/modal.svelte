<script lang="ts">
  import { setContext } from "svelte";
  import type { ModalControl } from "./index";

  let current_big_compo;
  let big_props = {};
  $: _show_big = true;

  let current_small_compo;
  let small_props = {};
  $: _show_small = true;

  const handle: ModalControl = {
    show_big: (_compo, _props) => {
      current_big_compo = _compo;
      big_props = _props;
      _show_big = true;
    },

    close_big: () => {
      _show_big = false;
      current_big_compo = null;
    },

    show_small: (_compo, _props) => {
      current_small_compo = _compo;
      small_props = _props;
      _show_small = true;
    },

    close_small: () => {
      _show_small = false;
      current_small_compo = null;
    },
  };

  setContext("__modal__", handle);
</script>

{#key _show_big}
  {#if _show_big && current_big_compo}
    <modal-big-wrapper>
      <div
        on:click={handle.close_big}
        class="modal-close absolute top-0 right-0 cursor-pointer flex flex-col items-center mt-4 mr-4 bg-white rounded-lg border-4 z-50"
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

      <div class="sm:w-3/4 lg:w-1/2 3xl:w-1/3 modal-big-section" >
        <svelte:component this={current_big_compo} {...big_props} />
      </div>

    </modal-big-wrapper>
  {/if}
{/key}

{#key _show_small}
  {#if _show_small && current_small_compo}
    <modal-small-wrapper>
      <div
        on:click={handle.close_small}
        class="modal-close absolute top-0 right-0 cursor-pointer flex flex-col items-center mt-4 mr-4 bg-white rounded-lg border-4 z-50"
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

      <modal-small-section>
        <svelte:component this={current_small_compo} {...small_props} />
      </modal-small-section>
    </modal-small-wrapper>


  {/if}
{/key}

<slot />

<style>
  modal-big-wrapper {
    position: fixed;
    top: 0px;
    right: 0;
    bottom: 0;
    left: 0;
    overflow: auto;
    z-index: 100;
    background-color: rgba(0, 0, 0, 0.44);
  }
  .modal-big-section {
    position: relative;
    height: 98vh;
    margin-left: auto;
    margin-right: auto;
    margin-top: 0.25rem;
    margin-bottom: 0.25rem;
    overflow: auto;
    background-color: #fff;
    border-radius: 6px;
    display: block;
    padding: 0.5em;
  }
  modal-small-wrapper {
    position: fixed;
    top: 0px;
    right: 0;
    bottom: 0;
    left: 0;
    overflow: auto;
    z-index: 150;
    background-color: rgba(0, 0, 0, 0.44);
  }
  modal-small-section {
    position: relative;
    height: 95vh;
    margin: auto;
    overflow: auto;
    width: 512px;
    background-color: #fff;
    border-radius: 6px;
    display: block;
    padding: 0.5em;
  }
</style>
