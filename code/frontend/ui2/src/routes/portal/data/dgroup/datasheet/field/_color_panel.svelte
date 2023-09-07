<script lang="ts">
  export let color = "";
  export let onColorChange;
  const colors = [
    "red",
    "blue",
    "green",
    "gray",
    "yellow",
    "zinc",
    "emerald",
    "orange",
    "lime",
    "amber",
    "teal",
    "cyan",
    "sky",
    "indigo",
    "violet",
    "purple",
    "fuchsia",
    "pink",
  ];
</script>

<div class="border rounded shadow p-4 mt-2 border-blue-500 overflow-auto">
  <label class="flex gap-2"
    >Color

    <select class="p-1 rouned w-20" value={color} on:change={onColorChange}>
      {#each colors as c}
        <option value={c}>{c}</option>
      {/each}
      <option value="#ffff">Custom</option>
    </select>

    {#if color && color.startsWith("#")}
      <input type="color" value={color} on:change={onColorChange} />
    {/if}
  </label>
</div>
