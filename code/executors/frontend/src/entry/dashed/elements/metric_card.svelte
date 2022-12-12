<script lang="ts">
  import type { Panel } from "../service";
  export let panel: Panel;
  export let data;

  let increase = data["increase"];
</script>

<div class="flex items-center">
  <div class="flex-shrink-0">
    <span class="text-2xl sm:text-3xl leading-none font-bold text-gray-900 "
      >{data["value1"] || ""}</span
    >
    <h3 class="text-base font-normal text-gray-500">{panel.name}</h3>
  </div>
  <div
    class="ml-5 w-0 flex items-center justify-end flex-1 text-base font-bold {increase
      ? 'text-green-500'
      : 'text-red-500'}"
  >
    {data["value2"] || ""}

    <svg
      class="w-5 h-5 {increase ? '' : 'rotate-180'}"
      fill="currentColor"
      viewBox="0 0 20 20"
      xmlns="http://www.w3.org/2000/svg"
    >
      <path
        fill-rule="evenodd"
        d="M5.293 7.707a1 1 0 010-1.414l4-4a1 1 0 011.414 0l4 4a1 1 0 01-1.414 1.414L11 5.414V17a1 1 0 11-2 0V5.414L6.707 7.707a1 1 0 01-1.414 0z"
        clip-rule="evenodd"
      />
    </svg>
  </div>
</div>
