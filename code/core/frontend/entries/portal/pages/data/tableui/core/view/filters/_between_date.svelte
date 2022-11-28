<script>
  export let value = "";
  const __value = value.split("#");

  let from = __value[0] || "";
  let to = __value[1] || "";
</script>

<div class="flex flex-col">
  <label>
    From
    <input type="datetime-local" class="w-full border" bind:value={from} />
  </label>
  <label>
    To
    <input type="datetime-local" class="w-full border" bind:value={to} />
  </label>
</div>
