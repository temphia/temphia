<script>
  const width_classes = {
    10: "flex-grow",
    1: "w-28",
    2: "w-32",
    3: "w-48",
    4: "w-64",
    5: "w-72",
    6: "w-96",
  };

  const height_classes = {
    10: "h-full",
    1: "h-28",
    2: "h-32",
    3: "h-48",
    4: "h-64",
    5: "h-72",
    6: "h-96",
  }

  export let width;
  export let height;


</script>

<div class="bg-white shadow rounded-lg p-2 {width_classes[width] || ""} {height_classes[height] || ""}">
  <slot />
</div>
