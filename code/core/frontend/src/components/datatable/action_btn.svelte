<script>
  export let action;
  export let clx = "";
  let name = action["name"] || "";
  let icon = action["icon"] || "";
</script>

<button
  class="py-1 px-3 text-gray-700 inline-flex font-light h-8 rounded hover:bg-gray-200 {clx}"
  on:click={action["action"]}
>
  {@html icon}
  {name}
</button>
