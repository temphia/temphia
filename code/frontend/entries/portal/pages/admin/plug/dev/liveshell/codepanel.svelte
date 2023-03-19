<script lang="ts">
  import { CEditor } from "../../../core";

  export let code = "";
  export let editor;
  export let file;
  export let loading = false;
</script>

{#if file}
  {#if !loading}
    <CEditor {code} bind:editor container_style={"height:100%;"} />
  {/if}
{/if}
