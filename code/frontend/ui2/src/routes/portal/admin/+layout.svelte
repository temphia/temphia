<script lang="ts">
  import { AppBar } from "@skeletonlabs/skeleton";
  import { items } from "./admin";

  import { afterNavigate } from "$app/navigation";
  const active =
    "text-blue-500 rounded-none border-b-2 font-medium border-blue-500";

  $: _current_page = location.pathname.split("/")[5];

  afterNavigate(() => {
    _current_page = location.pathname.split("/")[5];
  });
</script>

<div class="w-full h-full bg-indigo-100 overflow-auto">
  <AppBar
    padding="p-0"
    gridColumns="grid-cols-3"
    slotDefault="place-self-center"
    slotTrail="place-content-end"
  >
    <div slot="lead" />

    <div class="flex flex-row">
      {#each items as item}
        <a href="/z/pages/portal/admin{item.path}">
          <span
            class="text-gray-600 cursor-pointer p-2 md:p-3 hover:bg-red-100 block focus:outline-none uppercase {item.id ==
            _current_page
              ? active
              : ''}"
          >
            <svelte:component this={item.icon} />
          </span>
        </a>
      {/each}
    </div>
    <div slot="trail" />
  </AppBar>
  <div class="w-full h-full" style="height: calc(100vh - 3.5rem);">
    <slot>
      <p>Empty slot</p>
    </slot>
  </div>
</div>
