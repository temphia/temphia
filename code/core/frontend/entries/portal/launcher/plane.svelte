<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { Launcher } from "../services/engine/launcher";

  export let launcher: Launcher;
  const state = launcher.state;

  $: __instances = $state.instances;
  $: __active_instance = $state.active_instance;
</script>

<div
  class={$state.display === "HIDDEN"
    ? "hidden"
    : $state.display === "FLOATING"
    ? "floating"
    : "not-floating"}
>
  <div class="flex justify-between border border-gray-100">
    <div class="grow h-10 flex flex-row flex-nowrap overflow-hidden">
      {#each __instances as instance}
        <div
          on:click={() => launcher.instance_change(instance.id)}
          class="text-gray-600 py-2 px-2 flex hover:text-blue-500 cursor-pointer focus:outline-none {instance.id ===
          __active_instance
            ? 'text-blue-500 border-b-2 font-medium border-blue-500'
            : ''}"
        >
          {instance.name}

          <span on:click={() => launcher.instance_close(instance.id)}>
            <Icon name="x-circle" class="h-6 w-8 pt-1 hover:text-red-500" />
          </span>
        </div>
      {/each}
    </div>

    <div class="grow-0 h-8 w-8 p-1">
      <svg
        class="w-6 h-6 rounded border"
        fill="none"
        viewBox="0 0 24 24"
        stroke-width="1.5"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          d="M3.75 6A2.25 2.25 0 016 3.75h2.25A2.25 2.25 0 0110.5 6v2.25a2.25 2.25 0 01-2.25 2.25H6a2.25 2.25 0 01-2.25-2.25V6zM3.75 15.75A2.25 2.25 0 016 13.5h2.25a2.25 2.25 0 012.25 2.25V18a2.25 2.25 0 01-2.25 2.25H6A2.25 2.25 0 013.75 18v-2.25zM13.5 6a2.25 2.25 0 012.25-2.25H18A2.25 2.25 0 0120.25 6v2.25A2.25 2.25 0 0118 10.5h-2.25a2.25 2.25 0 01-2.25-2.25V6zM13.5 15.75a2.25 2.25 0 012.25-2.25H18a2.25 2.25 0 012.25 2.25V18A2.25 2.25 0 0118 20.25h-2.25A2.25 2.25 0 0113.5 18v-2.25z"
        />
      </svg>
    </div>
  </div>

  <div
    class="border border-red-600 w-full"
    style="height: calc(100vh - 2.75rem);"
  >
    {#each __instances as instance}
      <iframe
        title={instance.name}
        class="w-full h-full transition-all {instance.id === __active_instance
          ? ''
          : 'hidden'}"
        src="https://picsum.photos/seed/{instance.id}/1300/700"
      />
    {/each}
  </div>
</div>

<style>
  .hidden {
    display: none;
    width: 0px;
    height: 0px;
  }

  .floating {
    display: block;
    position: fixed;
    width: 100%;
    height: 100%;
    z-index: 1000;
  }

  .not-floating {
    display: block;
    width: 100%;
    height: 100%;
  }
</style>
