<script lang="ts">
  import Pertable from "./table/percentage.svelte";
  import MetricCard from "./card/metric.svelte";
  import AutoTable from "./table/auto.svelte";
  import ChartGraph from "./chart/chart.svelte";
  export let panel: any;
  export let data: any;
</script>

{#if panel.type === "hero_card"}
  <MetricCard {data} {panel} />

<!-- {:else if panel.type === "auto_table"}
  <AutoTable {data} {panel} /> 
{:else if panel.type === "per_table"}
  <Pertable {data} {panel} />
{:else if panel.type === "chart_graph"}
  <ChartGraph {data} {panel} />
  -->
{:else}
  <div>Not Impl</div>
{/if}

