<script>
  import Draggable from "./_draggable.svelte";
  import Plug from "./_plug.svelte";

  $: _main_plug_pos = [0, 0, 0, 0];
  $: _main_plug_center = [];


</script>

<div class="h-full w-full max-h-screen p-2">
  <div
    class="h-full w-full rounded overflow-scroll border border-slate-900 bg-white"
  >
    <div
      class="relative"
      style="min-width:5000px; min-height:5000px; transform: scale(1); 
      transform-origin: 0% 0% 0px;
      background-image: radial-gradient(rgba(15, 15, 16, 0.33) 1px, transparent 1px); 
            background-size: 13px 13px; background-color: rgba(71, 211, 255, 0.06);
      "
    >


      <Draggable>
        <Plug />
      </Draggable>
    </div>
  </div>
</div>
