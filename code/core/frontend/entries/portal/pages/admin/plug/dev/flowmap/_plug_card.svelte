<script>
  import { createEventDispatcher, onMount } from "svelte";
  import { tick } from "svelte/internal";
  import AgentCard from "./_agent_card.svelte";

  export let name;
  export let top = Math.floor(Math.random() * 400);
  export let left = Math.floor(Math.random() * 400);

  const dispatch = createEventDispatcher();

  let pos1 = 0,
    pos2 = 0,
    pos3 = 0,
    pos4 = 0;

  let elmnt;

  function dragMouseDown(e) {
    e = e || window.event;
    e.preventDefault();
    // get the mouse cursor position at startup:
    pos3 = e.clientX;
    pos4 = e.clientY;
    document.onmouseup = closeDragElement;
    // call a function whenever the cursor moves:
    document.onmousemove = elementDrag;
  }

  function elementDrag(e) {
    e = e || window.event;
    e.preventDefault();
    // calculate the new cursor position:
    pos1 = pos3 - e.clientX;
    pos2 = pos4 - e.clientY;
    pos3 = e.clientX;
    pos4 = e.clientY;
    // set the element's new position:
    elmnt.style.top = elmnt.offsetTop - pos2 + "px";
    elmnt.style.left = elmnt.offsetLeft - pos1 + "px";

    tick().then(dispatchPosition);
  }

  function closeDragElement() {
    // stop moving when mouse button is released:
    document.onmouseup = null;
    document.onmousemove = null;
  }

  const dispatchPosition = () => {
    console.log("ELEM", elmnt);

    dispatch("card_pos", {
      name,
      top: elmnt.offsetTop,
      left: elmnt.offsetLeft,
      height: elmnt.offsetHeight,
      width: elmnt.offsetWidth,
    });
  };

  onMount(dispatchPosition);

  // actions

  $: __mouse_over = false;
</script>

<div
  bind:this={elmnt}
  on:mouseenter={() => {
    __mouse_over = true;
  }}
  on:mouseleave={() => {
    __mouse_over = false;
  }}
  class="absolute  bg-white z-20 shadow rounded hover:border-2 border-sky-500 "
  style="min-width: 5rem; min-height: 5rem; resize: both;overflow: auto; top: {top}px; left: {left}px;"
>
  <div class="h-4 cursor-pointer w-full bg-yellow-10 flex justify-between">
    <div
      class="h-4 bg-yellow-50 hover:bg-yellow-300 grow flex justify-center px-2"
      on:mousedown={dragMouseDown}
    >
      <svg fill="currentColor" class="h-4 w-4" viewBox="0 0 16 16">
        <path
          d="M2 8a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm3 3a1 1 0 1 1 0 2 1 1 0 0 1 0-2zm0-3a1 1 0 1 1 0 2 1 1 0 0 1 0-2z"
        />
      </svg>
    </div>

    <div class="flex-none flex bg-yellow-100 gap-x-2 px-1 rounded">
      <svg
        on:click={() => dispatch("edit_block", name)}
        viewBox="0 0 20 20"
        fill="currentColor"
        class="w-4 h-4 hover:p-0.5 hover:bg-yellow-300"
      >
        <path
          d="M5.433 13.917l1.262-3.155A4 4 0 017.58 9.42l6.92-6.918a2.121 2.121 0 013 3l-6.92 6.918c-.383.383-.84.685-1.343.886l-3.154 1.262a.5.5 0 01-.65-.65z"
        />
        <path
          d="M3.5 5.75c0-.69.56-1.25 1.25-1.25H10A.75.75 0 0010 3H4.75A2.75 2.75 0 002 5.75v9.5A2.75 2.75 0 004.75 18h9.5A2.75 2.75 0 0017 15.25V10a.75.75 0 00-1.5 0v5.25c0 .69-.56 1.25-1.25 1.25h-9.5c-.69 0-1.25-.56-1.25-1.25v-9.5z"
        />
      </svg>
    </div>
  </div>

  <div class="relative">
    <div class="absolute h-5 w-20 bg-slate-500 -right-40">
        Hook Targets
    </div>

    <div class="absolute  h-5 w-20 bg-slate-500 -left-40">
        App Targets
    </div>
  </div>



  <!-- debug middle point -->

  <!-- <div
      class="h-2 w-2 rounded-full absolute right-1/2 top-1/2 bg-red-400 hover:bg-red-400"
    /> -->

  <div class="p-2 w-full flex flex-col gap-2">
    <AgentCard />
    <AgentCard />
    <AgentCard />


  </div>
</div>
