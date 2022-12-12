<script>
  import Draggable from "./v2/draggable.svelte";
  import Plug from "./v2/_plug.svelte";
  import PlugCard from "./_plug_card.svelte";

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
      <!-- <PlugCard
        name="Xyz"
        on:card_pos={(ev) => {
          _main_plug_pos = [
            ev.detail["top"],
            ev.detail["left"],
            ev.detail["height"],
            ev.detail["width"],
          ];
          _main_plug_pos = _main_plug_pos;
        }}
      /> -->

      <Draggable>
        <Plug />
      </Draggable>
    </div>
  </div>
</div>
