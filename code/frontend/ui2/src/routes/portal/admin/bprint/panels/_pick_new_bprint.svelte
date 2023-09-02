<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import type { PortalService } from "$lib/core";
  export let app: PortalService;

  const resourceTypes = [
    {
      name: "From Store",
      icon: "view-grid-add",
      info: "Import blueprint from repo store",
      action: () => app.nav.repo_loader(),
    },
    {
      name: "From zip",
      icon: "archive",
      info: "Upload using bprint zip",
      action: () => app.nav.admin_bprint_new_zip(),
    },
    {
      name: "Empty",
      icon: "sparkles",
      info: "Create Empty Bprint",
      action: () => app.nav.admin_bprint_new(),
    },
  ];
</script>

<div class="flex items-center justify-between">
  <h4 class="font-medium text-slate-500">How do you like your new blueprint?</h4>
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
