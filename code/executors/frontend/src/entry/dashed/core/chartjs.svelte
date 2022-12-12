<script>
  import Chart from "chart.js/dist/Chart";
  import { onMount, afterUpdate } from "svelte";

  let canvasRef;
  let chart;

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

  onMount(async () => draw());

  function draw() {
    let ctx = canvasRef.getContext("2d");
    chart = new Chart(ctx, {
      type,
      data,
      options,
    });
  }

  afterUpdate(() => {
    if (!chart) return;
    chart.data = data;
    chart.type = type;
    chart.options = options;
    //chart.plugins = plugins;
    chart.update();
  });
</script>

<canvas bind:this={canvasRef} />
