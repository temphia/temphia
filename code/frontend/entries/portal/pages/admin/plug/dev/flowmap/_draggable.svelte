<script>
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { createEventDispatcher, onMount } from "svelte";
  import { tick } from "svelte/internal";

  export let name = "";
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
  class="absolute  bg-white z-20 hover:border-2 border-sky-500 shadow rounded-lg"
  style="min-width: 15rem; min-height: 15rem; top: {top}px; left: {left}px; resize:both;"
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
      <Icon name="code" class="w-4 h-4 hover:p-0.5 hover:bg-yellow-300" />
    </div>
  </div>

  <!-- debug middle point -->

  <!-- <div
        class="h-2 w-2 rounded-full absolute right-1/2 top-1/2 bg-red-400 hover:bg-red-400"
      /> -->

  <slot />
</div>
