<script lang="ts">
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
    SearchIcon,
  } from "../../xcompo/svg";
  import SvelteTooltip from "svelte-tooltip";

  const active =
    "text-blue-500 rounded-none border-b-2 font-medium border-blue-500";

  export let items = [
    {
      id: "bprint",
      name: "Bluprints",
      icon: BluprintIcon,
      path: "#/admin/bprint/",
    },
    {
      id: "plug",
      name: "Plugs",
      icon: PlugIcon,
      path: "#/admin/plug/",
    },
    {
      id: "resource",
      name: "Resources",
      icon: ResourceIcon,
      path: "#/admin/resource/",
    },

    {
      id: "ugroup",
      name: "Users and Groups",
      icon: UserGroupIcon,
      path: "#/admin/ugroup/",
    },
    {
      id: "tenant",
      name: "Organization",
      icon: OrgIcon,
      path: "#/admin/tenant/",
    },
    {
      id: "data",
      name: "Data Tables",
      icon: GroupIcon,
      path: "#/admin/data/",
    },
    {
      id: "repo",
      name: "Repos",
      icon: StoreIcon,
      path: "#/admin/repo/",
    },

    {
      id: "lens",
      name: "lens",
      icon: SearchIcon,
      path: "#/admin/lens/app",
    },
  ];

  $: _current_page = location.hash.split("/")[2];
  window.addEventListener("hashchange", () => {
    _current_page = location.hash.split("/")[2];
  });
</script>

<div class="w-full h-full bg-indigo-100 overflow-auto">
  <div class="w-full fixed rounded shadow z-20">
    <nav class="flex flex-nowrap bg-white justify-center text-white">
      {#each items as item}
        <a href={item.path}>
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
        </a>
      {/each}
    </nav>
  </div>

  <div class="w-full mt-14">
    <slot>
      <p>Empty slot</p>
    </slot>
  </div>
</div>
