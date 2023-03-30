<script lang="ts">
  import Chartjs from "./_chartjs.svelte";
  export let panel: any;
  export let data: any;

  $: console.log("@chartjs", data, panel);

  const data_opts = panel["data_opts"] || {};
  const options = data_opts["data_opts"] || {};
  const chart_type = data_opts["chart_type"];
</script>

<div class="w-full p-1 flex justify-center">
  <h4 class="text-xl leading-none font-bold text-gray-600 mb-5">
    {panel.name}
  </h4>
</div>

<Chartjs {data} {options} type={chart_type} />
