<script lang="ts">
  import { createEventDispatcher } from "svelte";
  import SvelteTooltip from "svelte-tooltip";
  import {
    AdminIcon,
    FolderIcon,
    GroupIcon,
    PlugIcon,
    HomeIcon,
    StoreIcon,
    NotificationIcon,
    UserIcon,
  } from "../../xcompo/svg";

  import Logo from "../../xcompo/svg/logo.svelte";

  export let pending_notification = false;

  const route_links: [any, string, string][] = [
    [HomeIcon, "start", "#/"],
    [GroupIcon, "data tables", "#/data/"],
    [AdminIcon, "admin", "#/admin"],
    [FolderIcon, "cabinet", "#/cabinet/"],
    [StoreIcon, "store", "#/repo/"],
  ];

  const dispatch = createEventDispatcher();

  $: __open_menu = false;
</script>

<div class="h-screen w-screen flex tx-root">
  <!-- DESKTOP NAV -->
  <nav
    class="flex-col items-center bg-blue-200 text-gray-700 h-full w-12 shadow-lg tx-nav-desktop"
  >
    <!-- Side Nav Bar-->

    <div class="h-16 flex items-center w-full">
      <!-- Logo Section -->
      <a class="h-6 mx-auto" href="http://svelte.dev/">
        <Logo />
      </a>
    </div>

    <ul>
      {#each route_links as link}
        <li class="text-white hover:bg-gray-100">
          <SvelteTooltip tip={link[1]} right color="#7c3aed">
            <a
              href={link[2]}
              class="h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500"
            >
              <svelte:component this={link[0]} />
            </a>
          </SvelteTooltip>
        </li>
      {/each}
    </ul>

    <div class="mt-auto flex items-center w-full">
      <div class="flex flex-col items-center w-full">
        <div
          on:click={() => dispatch("notification_toggle")}
          class="
                h-16 cursor-pointer mx-auto w-full 
                flex justify-center items-center
                focus:text-orange-500 hover:bg-red-200 focus:outline-none relative"
        >
          {#if pending_notification}
            <span class="flex absolute -mt-5 ml-4">
              <span
                class="animate-ping absolute inline-flex h-3 w-3 rounded-full bg-pink-400 opacity-75"
              />
              <span
                class="relative inline-flex rounded-full h-3 w-3 bg-pink-500"
              />
            </span>
          {/if}

          <NotificationIcon />
        </div>

        <div
          class="h-12 cursor-pointer mx-auto flex justify-center items-center
              w-full focus:text-orange-500 hover:bg-red-200 focus:outline-none"
        >
          <UserIcon />
        </div>

        <button
          class="h-12 cursor-pointer mx-auto flex justify-center items-center
                        w-full focus:text-orange-500 hover:bg-red-200 focus:outline-none"
        >
          <svg
            class="h-5 w-5 text-red-700"
            xmlns="http://www.w3.org/2000/svg"
            width="24"
            height="24"
            viewBox="0 0 24 24"
            fill="none"
            stroke="currentColor"
            stroke-width="2"
            stroke-linecap="round"
            stroke-linejoin="round"
          >
            <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
            <polyline points="16 17 21 12 16 7" />
            <line x1="21" y1="12" x2="9" y2="12" />
          </svg>
        </button>
      </div>
    </div>
  </nav>

  <nav
    class="justify-between w-screen bg-blue-200 text-gray-700 shadow-lg tx-nav-mobile"
  >
    <div class=" py-2 flex w-full items-center">
      <a
        class="navbar-burger self-center mr-12"
        href="#"
        on:click={() => {
          __open_menu = !__open_menu;
        }}
      >
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6 hover:text-gray-200"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M4 6h16M4 12h16M4 18h16"
          />
        </svg>
      </a>
    </div>
    <!-- Responsive navbar -->
    <div
      class="flex mr-6 items-center cursor-pointer"
      on:click={() => dispatch("notification_toggle")}
    >
      <NotificationIcon />

      {#if pending_notification}
        <span class="flex absolute -mt-5 ml-4">
          <span
            class="animate-ping absolute inline-flex h-3 w-3 rounded-full bg-pink-400 opacity-75"
          />
          <span class="relative inline-flex rounded-full h-3 w-3 bg-pink-500" />
        </span>
      {/if}
    </div>

    <div class="flex mr-6 items-center cursor-pointer">
      <UserIcon />
    </div>
  </nav>

  {#if __open_menu}
    <div class="h-screen w-screen bg-gray-600 bg-opacity-90 fixed z-50">
      <div
        class="h-full w-full absolute transform translate-x-0 transition ease-in-out duration-700 p-5"
      >
        <div class="h-full bg-white rounded p-5">
          <div class="absolute right-4">
            <button
              aria-label="close menu modal"
              class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 cursor-pointer rounded border mr-2"
              on:click={() => {
                __open_menu = !__open_menu;
              }}
            >
              <svg
                class="h-6 w-6"
                viewBox="0 0 24 24"
                fill="none"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  d="M18 6L6 18"
                  stroke="#4B5563"
                  stroke-width="1.25"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
                <path
                  d="M6 6L18 18"
                  stroke="#4B5563"
                  stroke-width="1.25"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                />
              </svg>
            </button>
          </div>

          <ul class="space-y-4 pt-10">
            {#each route_links as link}
              <li class="border rounded">
                <a
                  href={link[2]}
                  on:click={() => {
                    __open_menu = false;
                  }}
                  class="h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500 uppercase"
                >
                  <svelte:component this={link[0]} />
                  {link[1]}
                </a>
              </li>
            {/each}

            <li class="border rounded">
              <button
                class="h-12 cursor-pointer mx-auto flex justify-center items-center
                                  w-full focus:text-orange-500 hover:bg-red-200 focus:outline-none"
              >
                <svg
                  class="h-5 w-5 text-red-700"
                  xmlns="http://www.w3.org/2000/svg"
                  width="24"
                  height="24"
                  viewBox="0 0 24 24"
                  fill="none"
                  stroke="currentColor"
                  stroke-width="2"
                  stroke-linecap="round"
                  stroke-linejoin="round"
                >
                  <path d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4" />
                  <polyline points="16 17 21 12 16 7" />
                  <line x1="21" y1="12" x2="9" y2="12" />
                </svg>
              </button>
            </li>
          </ul>
        </div>
      </div>
    </div>
  {/if}

  <div class="h-screen overflow-auto tx-main">
    <slot />
  </div>
</div>

<style>
  @media screen and (min-width: 768px) {
    .tx-main {
      width: calc(100vw - 3rem);
    }

    .tx-nav-desktop {
      display: flex;
    }

    .tx-nav-mobile {
      display: none;
    }
  }

  @media screen and (max-width: 768px) {
    .tx-root {
      flex-direction: column;
    }

    .tx-nav-desktop {
      display: none;
    }

    .tx-main {
      width: 100vw;
    }

    .tx-nav-mobile {
      display: flex;
    }
  }
</style>
