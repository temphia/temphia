<script lang="ts">
  export let script: string;
  export let key;

  let loaded = !!window[key] || false;

  const onLoad = (ev) => {
    loaded = true;
  };
</script>

<svelte:head>
  <script on:load={onLoad} src={loaded ? "" : script}></script>
</svelte:head>

{#if loaded}
  <slot />
{/if}
