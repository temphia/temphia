<script>
  import Basic from "../basic/basic.svelte";

  export let message = "";
  export let fields = [];
</script>

<div class="p-2 flex flex-col">
  {#each fields as field}
    <Basic {field} />
  {/each}
</div>
