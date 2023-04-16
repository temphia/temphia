<script lang="ts">
  import Pertable from "./table/percentage.svelte";
  import HeroCard from "./card/hero.svelte";
  import MetricCard from "./card/metric.svelte";
  import AutoTable from "./table/auto.svelte";

  import type { Element } from "../service";

  import ChartGraph from "./chart/chart.svelte";
  export let element: Element;
  export let data: any;
</script>

{#if element.type === "hero_card"}
  <HeroCard {data} {element} />
{:else if element.type === "metric_card"}
  <MetricCard {data} {element} />
{:else if element.type === "chart"}
  <ChartGraph {data} {element} />
{:else if element.type === "auto_table"}
  <AutoTable {data} {element} />
{:else}
  <div>Not Impl</div>
{/if}
