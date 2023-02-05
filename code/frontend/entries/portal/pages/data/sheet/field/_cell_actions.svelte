<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let is_cell_open = false;
  export let onCellClose = undefined;
  export let pick_icon = undefined;
  export let onPick = undefined;
  export let onCellColor;
</script>

<div class="flex gap-1 justify-end">
  {#if pick_icon !== undefined && !is_cell_open}
    <button on:click={() => onPick && onPick()}>
      <Icon
        class="w-6 h-6 p-0.5 rounded border hover:bg-yellow-100"
        name={pick_icon}
      />
    </button>
  {/if}

  <button on:click={onCellColor}>
    <Icon
      name="color-swatch"
      class="h-5 w-5 p-0.5 rounded border hover:bg-yellow-100"
    />
  </button>

  {#if is_cell_open}
    <button on:click={() => onCellClose && onCellClose()}>
      <Icon class="w-6 h-6 p-0.5 rounded border hover:bg-yellow-100" name="x" />
    </button>
  {/if}
</div>
