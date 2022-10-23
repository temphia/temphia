<script lang="ts">
  import Inner from "./notification.svelte";

  let show = false;

  export let loading = true;
  export let messages = [];

  export const toggle = () => (show = !show);
</script>

{#if show}
  <div
    class="w-full h-full bg-gray-800 bg-opacity-90 top-0 overflow-y-auto overflow-x-hidden fixed sticky-0 z-50 left-14"
  >
    <div
      class="w-full absolute h-full overflow-x-hidden transform translate-x-0 transition ease-in-out duration-700 "
    >
      <Inner
        on:ndelete
        on:nread
        on:refresh
        on:toggle_npanel
        {loading}
        {messages}
      />
    </div>
  </div>
{/if}
