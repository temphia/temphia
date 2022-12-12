<script lang="ts">
  import Pertable from "./per_table.svelte";
  import type { Panel } from "../service";
  import MetricCard from "./metric_card.svelte";
  import AutoTable from "./auto_table.svelte";
  import ChartGraph from "./chart_graph.svelte";
  export let panel: Panel;
  export let data: any;
</script>

{#if panel.type === "metric_card"}
  <MetricCard {data} {panel} />
{:else if panel.type === "auto_table"}
  <AutoTable {data} {panel} />
{:else if panel.type === "per_table"}
  <Pertable {data} {panel} />
{:else if panel.type === "chart_graph"}
  <ChartGraph {data} {panel} />
{:else}
  <div>Not Impl</div>
{/if}

<!-- 
  per_table
  metric_card
  chart
    line graph
    bar graph
    pie chart
    scatter
  auto_table
 -->
