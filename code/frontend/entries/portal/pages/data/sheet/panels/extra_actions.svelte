<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { SheetService } from "../../../../services/data";

  export let service: SheetService;

  let actions = [
    { name: "Export", icon: "download", info: "Export data as json." },
    { name: "Import", icon: "upload", info: "Import data as json." },
    { name: "View Options", icon: "adjustments", info: "Set view options." },
  ];
</script>

<div class=" flex flex-col gap-2">
  {#each actions as action}
    <button
      class="group flex gap-y-2 items-center gap-x-5 rounded-md px-2.5 py-2 transition-all duration-75 hover:bg-blue-100"
    >
      <div
        class="flex h-12 w-12 items-center rounded-lg bg-gray-200 text-black group-hover:bg-blue-200"
      >
        <Icon name={action.icon} class="text-gray-700" />
      </div>
      <div
        class="flex flex-col items-start justify-between font-light text-gray-600"
      >
        <p class="text-[15px]">{action.name}</p>
        <span class="text-xs font-light text-gray-400">{action.info}</span>
      </div>
    </button>
  {/each}
</div>
