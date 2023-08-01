<script lang="ts">
  import type { Element } from "../../service";

  export let element: Element;
  export let data = {};
</script>

<div class="flex flex-col py-2">
  <div class="flex items-center justify-center gap-2">
    <svg class="w-6 h-6">
      {@html element.data_opts["icon"] || ""}
    </svg>

    <span class="leading-none font-bold text-gray-900 ">
      {data["value"] || ""}
    </span>
  </div>

  <div class="flex items-center justify-center">
    {element.info}
  </div>
</div>
