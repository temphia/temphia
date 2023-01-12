<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { LoadingSpinner, PortalService } from "../../core";
  import { instance_helper } from "./instance";

  export let bid: string;
  export let app: PortalService;
  export let bprint: object;

  let bundle_objects;
  let loading = true;

  (async () => {
    const bapi = app.api_manager.get_admin_bprint_api();
    const resp = await bapi.get_file(bid, "schema.json");
    if (resp.status !== 200) {
      return;
    }
    bundle_objects = resp.data;
    loading = false;
  })();

  const iconTypes = {
    data_group: "collection",
    plug: "view-grid-add",
  };
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <div class="flex items-center justify-between">
    <h4 class="font-semibold text-lg text-slate-800">
      Pick a object to instance.
    </h4>
  </div>

  <div class="space-y-2 mt-4">
    {#each bundle_objects["items"] || [] as item}
      <div
        on:click={() => {
          instance_helper(app, item["type"], bprint, item["file"]);
          app.utils.small_modal_close();
        }}
        class="flex space-x-4 rounded-xl bg-white p-3 shadow-sm hover:border border-blue-500 cursor-pointer"
      >
        <Icon
          name={iconTypes[item["type"]] || "hashtag"}
          class="w-10 h-10 text-zinc-600"
        />

        <div>
          <h4 class="font-semibold text-gray-600">{item["name"]}</h4>
          <p class="text-sm text-slate-400">
            {item["type"]}
          </p>
        </div>
      </div>
    {/each}
  </div>
{/if}
