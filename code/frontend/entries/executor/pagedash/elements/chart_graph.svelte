<script lang="ts">
  import Chartjs from "./_chartjs.svelte";
  export let panel: any;
  export let data: any;

  let options = panel["options"] || {};
</script>

<div class="w-full p-1 flex justify-center">
  <h4 class="text-xl leading-none font-bold text-gray-600 mb-5">{panel.name}</h4>
</div>

<Chartjs {data} options={options["chart_options"]} type={options["type"]} />
