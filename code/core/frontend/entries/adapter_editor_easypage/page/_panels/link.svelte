<script lang="ts">
  import type { EasypageService } from "../../service/easypage";

  export let domain: string;
  export let slug: string;
  export let service: EasypageService;

  let subdomain_fill = "";
  $: _needs_subdomain_fill = domain.includes("*");
  $: _show_button =
    !_needs_subdomain_fill || (_needs_subdomain_fill && subdomain_fill);

  const openLink = () =>
    window.open(`${domain}/${slug}`.replace("*", subdomain_fill), "_blank");

  if (!_needs_subdomain_fill) {
    openLink();
    service.modal.small_close();
  }
</script>

<div class="flex flex-col">
  {#if _needs_subdomain_fill}
    <div class="mb-4">
      <label class="block mb-2 text-sm font-bold text-gray-700" for="subd"
        >Sub Domain Fill</label
      >
      <input
        class="w-full px-3 py-2 text-sm leading-tight text-gray-700 border rounded shadow appearance-none focus:outline-none focus:shadow-outline"
        id="subd"
        type="text"
        bind:value={subdomain_fill}
        placeholder="subdomain"
      />
    </div>
  {/if}

  {#if _show_button}
    <button
      on:click={openLink}
      class="w-full px-4 py-2 font-bold text-white bg-blue-500 rounded-full hover:bg-blue-700 focus:outline-none focus:shadow-outline"
      type="button"
    >
      Go
    </button>
  {/if}
</div>
