<script lang="ts">
  export let show = true;

  const hide = () => {
    show = false;
  };
</script>

{#if show}
  <data-plugin-modal-wrapper>
    <data-plugin-modal-section>
      <div
        on:click={hide}
        class="modal-close absolute -top-1 -right-1 bg-gray-300 rounded cursor-pointer flex flex-col items-center mt-4 mr-4 text-white text-sm z-50"
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

      <slot></slot>
    </data-plugin-modal-section>
  </data-plugin-modal-wrapper>
{/if}

<style>
  data-plugin-modal-wrapper {
    position: fixed;
    top: 0px;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 8;
    background-color: rgba(0, 0, 0, 0.44);
    border: rgb(216, 168, 248);
  }
  data-plugin-modal-section {
    overflow-y: auto;
    position: relative;
    margin: auto;
    margin-top: 1rem;
    height: 90vh;
    max-width: 750px;
    background-color: #fff;
    border-radius: 6px;
    display: block;
    padding: 0.25rem;
  }
</style>
