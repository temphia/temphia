<script lang="ts">
  import { afterUpdate, getContext } from "svelte";
  import type { Launcher } from "$lib/services/portal/launcher/launcher";
  import type { PortalService } from "$lib/services/portal/portal";

  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import Instance from "./hostproxy/hostproxy.svelte";

  const app = getContext("__app__") as PortalService;

  export let launcher: Launcher;
  const state = launcher.state;
  $: __instances = $state.instances;
  $: __active_instance = $state.active_instance;
  $: __hidden = $state.display === "HIDDEN";
  $: __floating = $state.display === "FLOATING";

  let active_scrolled;

  const fix_tab = () => {
    if (!__active_instance || __hidden) {
      return;
    }

    if (active_scrolled === __active_instance) {
      return;
    }
    const item = document.getElementById(`itab-item-${__active_instance}`);
    item && item.scrollIntoView();
    active_scrolled = __active_instance;
    console.log("@fix_tab");
  };

  afterUpdate(fix_tab);
</script>

<div
  class="bg-white {__hidden
    ? 'hidden'
    : __floating
    ? 'floating'
    : 'not-floating'}"
>
  {#if __active_instance}
    <div class="flex justify-between border border-gray-100">
      <div class="grow h-10 flex flex-row flex-nowrap overflow-hidden">
        {#each __instances as instance}
          <button
            id={`itab-item-${instance.id}`}
            on:click={() => {
              if (__floating) {
                launcher.instance_change(instance.id);
              } else {
                app.nav.launch_target(instance.target_id);
              }
            }}
            class="text-gray-600 py-2 px-2 flex hover:text-blue-500 cursor-pointer focus:outline-none {instance.id ===
            __active_instance
              ? 'text-blue-500 border-b-2 font-medium border-blue-500'
              : ''}"
          >
            {instance.name || ""} [{instance.id || ""}] &nbsp;&nbsp;

            <button on:click={() => launcher.instance_close(instance.id)}>
              <Icon
                name="x-circle"
                class="w-5 pt-1 text-gray-500 hover:text-red-500"
              />
            </button>
          </button>
        {/each}
      </div>

      <div class="grow-0 h-8 w-8 p-1 flex">
        {#if __floating}
          <button
            on:click={() => {
              launcher.plane_hide();
            }}
          >
            <Icon
              name="eye-off"
              class="w-6 h-6 rounded border hover:text-blue-500"
            />
          </button>
        {/if}
      </div>
    </div>

    <div
      class="border border-red-600 w-full"
      style="height: calc(100vh - 2.75rem);"
    >
      {#each __instances as instance}
        <div
          class="w-full h-full {instance.id === __active_instance
            ? ''
            : 'hidden'}"
        >
          <Instance options={instance} />
        </div>
      {/each}
    </div>
  {:else if !__hidden}
    <div>No Apps running</div>
  {/if}
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
    width: calc(100vw - 3em);
    height: 100%;
    z-index: 1000;
  }

  @media screen and (max-width: 768px) {
    .floating {
      width: 100vw;
      /* height: calc(100vh - 3em); */
    }
  }

  .not-floating {
    display: block;
    width: 100%;
    height: 100%;
  }
</style>
