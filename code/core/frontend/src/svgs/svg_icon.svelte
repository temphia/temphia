<script>
  export let name = "";
</script>

<svg
  width="20"
  height="20"
  fill="currentColor"
  class="m-auto"
  viewBox="0 0 24 24"
  stroke="currentColor"
>
  {#if name === "admin"}
    <path
      stroke-linecap="round"
      stroke-linejoin="round"
      stroke-width="2"
      d="M12 6V4m0 2a2 2 0 100 4m0-4a2 2 0 110 4m-6 8a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4m6 6v10m6-2a2 2 0 100-4m0 4a2 2 0 110-4m0 4v2m0-6V4"
    />
  {:else if name === "table_group"}
    <path
      d="M5.45 5.11L2 12v6a2 2 0 0 0 2 2h16a2 2 0 0 0
                  2-2v-6l-3.45-6.89A2 2 0 0 0 16.76 4H7.24a2 2 0 0
                  0-1.79 1.11z"
    />
  {:else}
    <path />
  {/if}
</svg>
