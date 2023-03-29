<script lang="ts">
  import type { Panel } from "../../service";

  export let panel: Panel;
  export let data;
</script>

<div class="flex flex-col p-4">
  <div class="flex items-center justify-center">
    <div class="inline-flex">
      <svg>
        {@html panel.data_opts["icon"] || ""}
      </svg>

      <span class="text-2xl sm:text-3xl leading-none font-bold text-gray-900 ">
        {data || ""}
      </span>
    </div>
  </div>

  <div class="flex items-center justify-center">
    {panel.info}
  </div>
</div>
