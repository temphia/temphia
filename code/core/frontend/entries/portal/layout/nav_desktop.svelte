<script lang="ts">
  import { Link } from "svelte-routing";
  import SvelteTooltip from "svelte-tooltip";
  import type { PortalApp } from "../../../src/lib/app/portal";

  import {
    AdminIcon,
    FolderIcon,
    GroupIcon,
    PlugIcon,
    HomeIcon,
    StoreIcon,
    NotificationIcon,
    UserIcon,
  } from "../../../src/svgs";

  import Logo from "../../../src/svgs/logo.svelte";

  const route_links = [
    [HomeIcon, "start", "/z/portal"],
    [GroupIcon, "data tables", "/z/portal/dtable_load"],
    [PlugIcon, "plug apps", "/z/portal/apps_launcher"],
    [AdminIcon, "admin", "/z/portal/admin/plugs"],
    [FolderIcon, "cabinet", "/z/portal/cabinet_load"],
    [StoreIcon, "store", "/z/portal/store"],
  ];

  export let app: PortalApp;

  const noti = app.notifier;
  const noti_pending_read = noti.isPendingRead;
</script>

<nav
  class="flex flex-col items-center bg-blue-200 text-gray-700 shadow h-full w-12 shadow-lg"
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
          <Link
            to={link[2] + ""}
            getProps={() => ({
              class:
                "h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500",
            })}
          >
            <svelte:component this={link[0]} />
          </Link>
        </SvelteTooltip>
      </li>
    {/each}
  </ul>

  <div class="mt-auto flex items-center w-full">
    <div class="flex flex-col items-center w-full">
      <div
        on:click={noti.toggle_notification}
        class="
          h-16 cursor-pointer mx-auto w-full 
          flex justify-center items-center
          focus:text-orange-500 hover:bg-red-200 focus:outline-none relative"
      >
        {#if $noti_pending_read}
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
        on:click={app.navigator.goto_self_profile}
        class="h-12 cursor-pointer mx-auto flex justify-center items-center
        w-full focus:text-orange-500 hover:bg-red-200 focus:outline-none"
      >
        <UserIcon />
      </div>

      <button
        on:click={app.log_out}
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
