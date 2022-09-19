<script lang="ts">
  import { Link } from "svelte-routing";
  import {
    BluprintIcon,
    EventIcon,
    GroupIcon,
    UserIcon,
    PlugIcon,
    CodeIcon,
    OrgIcon,
    UserGroupIcon,
    ResourceIcon,
    StoreIcon,
    SearchIcon
  } from "../../../_svgs";
  import SvelteTooltip from "svelte-tooltip";

  const active =
    "text-blue-500 rounded-none border-b-2 font-medium border-blue-500";

  export let items = [
    {
      id: "bprints",
      name: "Bluprints",
      icon: BluprintIcon,
      path: "z/portal/admin/bprints",
    },
    {
      id: "plugs",
      name: "Plugs",
      icon: PlugIcon,
      path: "z/portal/admin/plugs",
    },
    {
      id: "resources",
      name: "Resources",
      icon: ResourceIcon,
      path: "z/portal/admin/resources",
    },

    {
      id: "user_groups",
      name: "Users and Groups",
      icon: UserGroupIcon,
      path: "z/portal/admin/user_groups",
    },
    {
      id: "ns",
      name: "Organization",
      icon: OrgIcon,
      path: "z/portal/admin/ns",
    },
    {
      id: "dtable",
      name: "Tables",
      icon: GroupIcon,
      path: "z/portal/admin/dtable",
    },
    {
      id: "repo",
      name: "Repos",
      icon: StoreIcon,
      path: "z/portal/admin/repo",
    },

    {
      id: "lens",
      name: "lens",
      icon: SearchIcon,
      path: "z/portal/admin/lens",
    },
  ];

  export let current_item: string = ""
  export let loading = false;

  $: _current_page = current_item
    ? current_item
    : location.pathname.split("/")[4];
</script>

<div class="w-full h-full bg-indigo-100 overflow-auto">
  <div class="w-full fixed rounded shadow z-20">
    <nav class="flex flex-nowrap bg-white justify-center text-white">
      {#each items as item}
        <Link to={item.path}>
          <SvelteTooltip tip={item.name} right color="#7c3aed">
            <span
              class="text-gray-600 cursor-pointer py-3 px-4 hover:bg-red-100 block focus:outline-none uppercase {item.id ==
              _current_page
                ? active
                : ''}"
            >
              <svelte:component this={item.icon} />
            </span>
          </SvelteTooltip>
        </Link>
      {/each}
    </nav>
  </div>

  <div class="w-full mt-14">
    {#if loading}
      <div class="p-10 flex justify-center">
        <div class="lds-ring">
          <div />
          <div />
          <div />
          <div />
        </div>
      </div>
    {:else}
      <slot>
        <p>Empty slot</p>
      </slot>
    {/if}
  </div>
</div>

<style>
  .lds-ring {
    display: inline-block;
    position: relative;
    width: 2.5rem;
    height: 2.5rem;
  }
  .lds-ring div {
    box-sizing: border-box;
    display: block;
    position: absolute;
    width: 2.5rem;
    height: 2.5rem;
    margin: 0.25rem;
    border: 0.25rem solid #726e6e;
    border-radius: 50%;
    animation: lds-ring 1.2s cubic-bezier(0.5, 0, 0.5, 1) infinite;
    border-color: #5f5c5c transparent transparent transparent;
  }
  .lds-ring div:nth-child(1) {
    animation-delay: -0.45s;
  }
  .lds-ring div:nth-child(2) {
    animation-delay: -0.3s;
  }
  .lds-ring div:nth-child(3) {
    animation-delay: -0.15s;
  }
  @keyframes lds-ring {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
</style>
