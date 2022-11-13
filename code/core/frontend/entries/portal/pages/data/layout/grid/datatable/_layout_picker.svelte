<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { onMount } from "svelte";
  import { scale } from "svelte/transition";

  let show = false; // menu state
  let menuRef = null; // menu wrapper DOM reference

  onMount(() => {
    const handleOutsideClick = (event) => {
      if (show && !menuRef.contains(event.target)) {
        show = false;
      }
    };

    const handleEscape = (event) => {
      if (show && event.key === "Escape") {
        show = false;
      }
    };

    // add events when element is added to the DOM
    document.addEventListener("click", handleOutsideClick, false);
    document.addEventListener("keyup", handleEscape, false);

    // remove events when element is removed from the DOM
    return () => {
      document.removeEventListener("click", handleOutsideClick, false);
      document.removeEventListener("keyup", handleEscape, false);
    };
  });
</script>

<div class="relative" bind:this={menuRef}>
  <div>
    <button
      on:click={() => (show = !show)}
      class="flex focus:shadow-solid p-1 text-gray-700 font-light rounded hover:bg-gray-200 bg-gray-50"
    >
      <Icon name="view-boards" class="h-5 w-5 pt-1" />
      Layout
    </button>

    {#if show}
      <div
        in:scale={{ duration: 100, start: 0.95 }}
        out:scale={{ duration: 75, start: 0.95 }}
        class="origin-top-right absolute right-0 w-48 py-2 mt-1 bg-white z-50
              rounded shadow-md flex flex-col"
      >
        <slot />
      </div>
    {/if}
  </div>
</div>


