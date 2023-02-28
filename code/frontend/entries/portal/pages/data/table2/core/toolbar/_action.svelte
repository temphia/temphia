<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let name = "";
  export let icon = "";
  export let active = false;
  export let tint = false;
  export let onClick;
</script>

<button
  on:click={onClick}
  class="p-1 text-gray-700 inline-flex font-light rounded hover:bg-gray-200 bg-gray-50 {active ||
  tint
    ? 'bg-blue-400'
    : ''}"
>
  <Icon name={icon} class="h-5 w-5 mt-1 mr-1" />
  {name}
</button>
