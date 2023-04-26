<script lang="ts">
  import Pertable from "./table/percentage.svelte";
  import HeroCard from "./card/hero.svelte";
  import MetricCard from "./card/metric.svelte";
  import AutoTable from "./table/auto.svelte";
  import Sheet from "./table/sheet.svelte";
  import ChartGraph from "./chart/chart.svelte";
  import Dump from "./dump/dump.svelte";

  import type { Element } from "../service";


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
{:else if element.type === "sheet_table"}
  <Sheet {data} {element} />
{:else if element.type === "dump"}
  <Dump />
{:else}
  <div>Not Impl</div>
{/if}
