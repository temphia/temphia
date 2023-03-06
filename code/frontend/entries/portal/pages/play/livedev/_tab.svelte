<script lang="ts">
  export let name;
  export let active;
  export let onClick = () => {};
</script>

<button
  on:click={onClick}
  class="text-gray-600 py-2 px-2 flex hover:text-blue-500 cursor-pointer focus:outline-none {active
    ? 'text-blue-500 border-b-2 font-medium border-blue-500'
    : ''}"
>
  {name} &nbsp;&nbsp;
</button>
