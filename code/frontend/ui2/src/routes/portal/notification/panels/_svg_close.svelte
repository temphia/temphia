<script lang="ts">
  export let onClick;
</script>

<div
  tabindex="0"
  aria-label="delete icon"
  role="button"
  class="focus:outline-none cursor-pointer"
  on:click={onClick}
>
  <svg
    width="14"
    height="14"
    viewBox="0 0 14 14"
    fill="none"
    xmlns="http://www.w3.org/2000/svg"
  >
    <path
      d="M10.5 3.5L3.5 10.5"
      stroke="#4B5563"
      stroke-width="1.25"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
    <path
      d="M3.5 3.5L10.5 10.5"
      stroke="#4B5563"
      stroke-width="1.25"
      stroke-linecap="round"
      stroke-linejoin="round"
    />
  </svg>
</div>
