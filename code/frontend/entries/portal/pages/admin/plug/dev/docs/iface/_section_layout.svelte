<script lang="ts">
  export let name: string;
</script>

<section class="flex flex-col p-1 w-full">
  <h3 class="text-2xl font-semibold text-gray-700 capitalize">{name}</h3>

  <div class="h-full w-full overflow-auto p-2">
    <slot />
  </div>
</section>
