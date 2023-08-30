<script>
  $: _show = true;
  let current_component;
  let props = {};

  // fixme => export const showmodal and closemodal
  window.showModal = (_compo, _props) => {
    current_component = _compo;
    props = _props;
    _show = true;
  };

  window.closeModal = () => {
    _show = false;
  };
</script>

{#key _show}
  {#if _show && current_component}
    <modal-wrapper>
      <div
        on:click={window.closeModal}
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

      <modal-section>
        <svelte:component this={current_component} {...props} />
      </modal-section>
    </modal-wrapper>
  {/if}
{/key}

<style>
  modal-wrapper {
    position: fixed;
    top: 0px;
    right: 0;
    bottom: 0;
    left: 0;
    overflow: auto;
    z-index: 999;
    background-color: rgba(0, 0, 0, 0.44);
  }
  modal-section {
    position: relative;
    height: 95vh;
    margin: auto;
    overflow: auto;
    width: 90vw;
    background-color: #fff;
    border-radius: 6px;
    display: block;
    padding: 0.5em;
  }
</style>
