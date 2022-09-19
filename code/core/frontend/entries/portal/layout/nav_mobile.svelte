<script lang="ts">
    import { Link, navigate } from "svelte-routing";
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

    $: __open_menu = false;
</script>

<nav class="flex justify-between w-screen bg-blue-200 text-gray-700 shadow-lg">
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
    <div class="flex mr-6 items-center cursor-pointer" on:click={noti.toggle_notification}>
        <NotificationIcon />

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
    </div>

    <div
        class="flex mr-6 items-center cursor-pointer"
        on:click={app.navigator.goto_self_profile}
    >
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
                        role="button"
                        aria-label="close menu modal"
                        class="focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-gray-500 rounded-md cursor-pointer rounded border mr-2"
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
                            <Link
                                to={link[2] + ""}
                                getProps={() => ({
                                    class: "h-12 px-4 flex justify-center items-center w-full text-gray-700 focus:text-orange-500 uppercase",
                                })}
                            >
                                <svelte:component this={link[0]} />

                                {link[1]}
                            </Link>
                        </li>
                    {/each}

                    <li class="border rounded">
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
                                <path
                                    d="M9 21H5a2 2 0 0 1-2-2V5a2 2 0 0 1 2-2h4"
                                />
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
