<script lang="ts">
  export let value = "";
  export let column: object;

  const options = column["options"] || [];
</script>

<select bind:value class="p-1 border bg-gray-50 rounded">
  {#each options as opt}
    <option value={opt}>{opt}</option>
  {/each}
</select>
