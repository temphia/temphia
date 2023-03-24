<script>
  import { onMount, afterUpdate } from "svelte";

  export let type = "line";
  export let options = {
    scales: {
      yAxes: [
        {
          ticks: {
            beginAtZero: true,
          },
        },
      ],
    },
  };
  export let data = {};
  export function addData(params) {}

  let canvasRef;
  let chart;

  onMount(async () => draw());

  const draw = () => {
    if (chart) {
      return;
    }

    if (!window["Chart"]) {
      return;
    }

    let ctx = canvasRef.getContext("2d");
    chart = new Chart(ctx, {
      type,
      data,
      options,
    });
  };
  afterUpdate(() => {
    if (!chart) return;
    chart.data = data;
    chart.type = type;
    chart.options = options;
    //chart.plugins = plugins;
    chart.update();
  });

  const loadChartJs = () => draw();
</script>

<svelte:head>
  <script on:load={loadChartJs} src="/z/assets/lib/chartjs/chartjs.js"></script>
</svelte:head>

<canvas bind:this={canvasRef} />
