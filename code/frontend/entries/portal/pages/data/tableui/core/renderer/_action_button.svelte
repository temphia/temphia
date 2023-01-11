<script>
  export let action;
  export let clx = "";
  let name = action["name"] || "";
  let icon = action["icon"] || "";
</script>

<button
  class="p-1 text-gray-700 inline-flex font-light rounded hover:bg-gray-200 {clx}"
  on:click={action["action"]}
>
  <svg class="h-4 w-4" viewBox="0 0 20 20" fill="currentColor">
    {@html icon}
  </svg>

  {name}
</button>
