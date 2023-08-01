<script lang="ts">
  import Tailwind from "../../../xcompo/common/_tailwind.svelte";

  export let title = "";
</script>

<div class="h-full w-full p-2 bg-blue-50 flex justify-center">
  <div
    class="pb-1 px-2 pt-2 bg-white rounded w-full flex flex-col overflow-auto"
    style="max-width: 700px;"
  >
    <h1
      class="text-gray-800 font-lg font-bold tracking-normal leading-tight mb-4"
    >
      {title}
    </h1>

    <slot />
  </div>
</div>

<Tailwind />
