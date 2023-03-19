<script lang="ts">
  import Execute from "../execute/execute.svelte";
  export const onRun = () => {
    __epoch = __epoch + 1;
  };
  $: __epoch = 0;
</script>

{#key __epoch}
  <Execute />
{/key}
