<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";

  export let actions: {
    name: string;
    icon: string;
    info: string;
    action: () => void;
  }[];

  export let title: string;
</script>

<div class="flex items-center justify-between">
  <h4 class="font-medium text-slate-500">{title}</h4>
</div>
<div class="space-y-2 mt-4">
  {#each actions as rt}
    <div
      on:click={() => {
        rt.action();
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
