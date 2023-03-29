<script lang="ts">
  import Pertable from "./table/percentage.svelte";
  import HeroCard from "./card/hero.svelte";
  import MetricCard from "./card/metric.svelte";
  import AutoTable from "./table/auto.svelte";
  import ChartGraph from "./chart/chart.svelte";
  export let panel: any;
  export let data: any;
</script>

{#if panel.type === "hero_card"}
  <HeroCard {data} {panel} />
{:else if panel.type === "metric_card"}
  <MetricCard {data} {panel} />

  <!-- 
{:else if panel.type === "per_table"}
  <Pertable {data} {panel} />
{:else if panel.type === "chart_graph"}
  <ChartGraph {data} {panel} />
  -->
{:else}
  <div>Not Impl</div>
{/if}
