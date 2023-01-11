<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { PortalService } from "../core";
  export let app: PortalService;

  const resourceTypes = [
    {
      name: "Data Table",
      icon: "collection",
      info: "Create a data table resource to perform operations on data",
      action: () => app.nav.admin_resource_datagroup_new(),
    },
    {
      name: "Room",
      icon: "chat-alt",
      info: "Create a room resource that can publish messages",
      action: () => app.nav.admin_resource_room_new(),
    },
    {
      name: "Folder",
      icon: "folder",
      info: "Create a Folder resource to store and retrive files",
      action: () => app.nav.admin_resource_folder_new(),
    },
    {
      name: "Module",
      icon: "hashtag",
      action: () => {},
      info: "Create a Module resource that can call native extended functionality",
    },
    {
      name: "Custom",
      icon: "cube",
      info: "Create a Custom resource with own scope and attributes",
      action: () => app.nav.admin_resource_new(),
    },
  ];
</script>

<div class="flex items-center justify-between">
  <h4 class="font-medium text-slate-500">Pick a resource type.</h4>
</div>
<div class="space-y-2 mt-4">
  {#each resourceTypes as rt}
    <div
      on:click={() => {
        rt.action();
        app.utils.small_modal_close();
      }}
      class="flex space-x-4 rounded-xl bg-white p-3 shadow-sm hover:border border-blue-500 cursor-pointer"
    >
      <Icon name={rt.icon} class="w-10 h-10 text-zinc-600" />

      <div>
        <h4 class="font-semibold text-gray-600">{rt.name}</h4>
        <p class="text-sm text-slate-400">
          {rt.info}
        </p>
      </div>
    </div>
  {/each}
</div>
