<script lang="ts">
  import { LoadingSpinner, PortalService } from "../../core";
  import { ResourceModule } from "../../../../../../lib/entities";
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { generateId } from "../../../../../../lib/utils";

  export let service: PortalService;

  let slug = generateId();
  let name = "";
  let module = "";
  let loading = true;
  let modules = [];
  let open = false;

  const Next = async () => {
    service.nav.admin_resource_new({
      slug,
      type: ResourceModule,
      sub_type: module,
      name,
    });

    service.utils.small_modal_close();
  };

  const load = async () => {
    const sapi = service.api_manager.get_self_api();
    const resp = await sapi.list_modules();
    if (!resp.ok) {
      console.log("@err", resp);
      return;
    }

    modules = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="p-1 flex flex-col">
    <div class="text-2xl text-indigo-900">New Module</div>
    <p class="text-red-500" />
    <div class="flex-col flex py-3">
      <label for="field-0" class="pb-2 text-gray-700 font-semibold">Id</label>
      <input
        type="text"
        bind:value={slug}
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <div class="flex-col flex py-3">
      <label for="field-1" class="pb-2 text-gray-700 font-semibold"
        >Module</label
      >
      {#if open}
        <input
          type="text"
          list="mod-list"
          bind:value={module}
          class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200 w-full"
        />

        <datalist id="mod-list">
          {#each modules || [] as opt}
            <option value={opt}>{opt}</option>
          {/each}
        </datalist>
      {:else}
        <select class="p-1 rounded border w-full" bind:value={module}>
          {#each modules || [] as opt}
            <option value={opt}>{opt}</option>
          {/each}
        </select>
      {/if}

      <div class="w-10 p-1 text-gray-700">
        <button on:click={() => (open = !open)}
          ><Icon
            name={open ? "lock-open" : "lock-closed"}
            class="w-6 h-6"
          /></button
        >
      </div>
    </div>

    <div class="flex-col flex py-3">
      <label for="field-1" class="pb-2 text-gray-700 font-semibold">Name</label>
      <input
        bind:value={name}
        type="text"
        class="p-2 shadow rounded-lg bg-gray-100 outline-none focus:bg-gray-200"
      />
    </div>

    <button
      on:click={Next}
      class="p-1 bg-blue-600 rounded hover:bg-blue-800 text-white">Next</button
    >
  </div>
{/if}
