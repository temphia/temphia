<script lang="ts">
  import { fromGeoJson } from "$lib/utils";

  export let value;

  const value_point = fromGeoJson(value);
</script>

<div class="flex gap-1">
  <span class="bg-yellow-100 rounded p-0.5 text-gray-600">
    Lat
    <strong class="font-semibold text-gray-700">{value_point[0].toFixed(3)} </strong>
  </span>

  <span class="bg-yellow-100 rounded p-0.5 text-gray-600">
    Long
    <strong class="font-semibold text-gray-700">{value_point[1].toFixed(3)}</strong>
  </span>
</div>
