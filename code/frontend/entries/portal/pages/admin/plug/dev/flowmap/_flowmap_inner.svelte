<script lang="ts">
  import { LoadingSpinner } from "../../../core";
  import type { PortalService } from "../../../core";
  import Draggable from "./_draggable.svelte";
  import {
    FormatedPlug,
    formatFlowData,
    generateAgentLinkIds,
    hashedPosCalc,
  } from "./formatter";
  import Plug from "./plug.svelte";
  import { getContext, onMount, tick } from "svelte";

  export let pid: string;

  const app = getContext("__app__") as PortalService;

  let rootElem;

  let loading = true;
  let data: FormatedPlug[];

  let links: [string, string][] = [];

  const papi = app.api_manager.get_admin_plug_api();

  const load = async () => {
    const resp = await papi.flowmap(pid);
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    data = formatFlowData(resp.data);
    loading = false;

    console.log("@@FORMATED_FLOWDATA", data);

    links = generateAgentLinkIds(resp.data);
  };

  const middle = { top: 1000, left: 2500 };

  $: _zoom_level = 1;

  load();

  const instances = [];

  $: console.log("@links", links);

  $: {
    if (!loading && links && rootElem) {
      console.log("@rendering", links);
      renderLinks();
    }
  }

  const renderLinks = async () => {
    await tick();

    links.forEach((link) => {
      const from = document.getElementById(link[0]);
      const to = document.getElementById(link[1]);

      console.log("RENDERING", link, from, to);

      if (!from || !to) {
        console.log("SKIPPING", link, from, to);
        return;
      }

      const ln = new window["LeaderLine"](from, to);

      ln.position();

      instances.push(ln);

      console.log("@ln", ln);
    });
  };

  const rePositionLinks = () => instances.forEach((ln) => ln.position())

</script>

<div class="h-full w-full max-h-screen p-2" bind:this={rootElem}>
  <div
    class="h-full w-full rounded border border-slate-900 bg-white overflow-auto"
    on:scroll={rePositionLinks}
    >
    <div class="fixed z-50 bottom-8 md:bottom-1 right-5 p-1">
      <div class="flex gap-1 p-0.5 text-xs bg-gray-100 rounded">
        <button
          class="p-1 rounded text-white bg-gray-600 hover:bg-blue-600"
          on:click={() => {
            if (_zoom_level < 0.2) {
              return;
            }
            _zoom_level = _zoom_level - 0.1;
            rePositionLinks()
          }}>-</button
        >
        <button
          class="p-1 rounded text-white bg-gray-600"
          on:click={() => (_zoom_level = 1)}
        >
          {(_zoom_level * 100).toFixed(0)}%
        </button>

        <button
          class="p-1 rounded text-white bg-gray-600 hover:bg-blue-600"
          on:click={() => {
            if (_zoom_level > 3) {
              return;
            }
            _zoom_level = _zoom_level + 0.1;
          }}>+</button
        >
      </div>
    </div>

    {#if loading}
      <LoadingSpinner />
    {:else}
      <div
        class="relative w-full h-full"
        style="min-width:5000px; min-height:5000px; transform: scale({_zoom_level}); 
    transform-origin: 0% 0% 0px;
    background-image: radial-gradient(rgba(15, 15, 16, 0.33) 1px, transparent 1px); 
          background-size: 13px 13px; background-color: rgba(71, 211, 255, 0.06);
    "
      >
        {#each data as fdata}
          {@const hash = pid + fdata.plug.id}
          {@const pos =
            fdata.plug.id === pid
              ? middle
              : hashedPosCalc(hash, 5000, 5000, middle, 500)}
          <Draggable left={pos.left} top={pos.top} on:card_pos={rePositionLinks}>
            <Plug data={fdata} />
          </Draggable>
        {/each}
      </div>
    {/if}
  </div>
</div>
