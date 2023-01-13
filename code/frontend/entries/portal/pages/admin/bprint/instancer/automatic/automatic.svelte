<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { ActionButton } from "../../../core";
  import { LoadingSpinner, PortalService } from "../../../core";

  import { params } from "svelte-hash-router";

  export let bid = $params.bid;

  const app = getContext("__app__") as PortalService;
  const bapi = app.api_manager.get_admin_bprint_api();

  let bundle_objects;
  let loading = true;

  const load = async () => {
    const resp = await bapi.get_file(bid, "schema.json");
    if (resp.status !== 200) {
      return;
    }
    bundle_objects = resp.data;
    loading = false;
  };

  const instance = async () => {
    bapi.instance(bid, {
      auto: true,
      instancer_type: "bundle",
      file: "schema.json",
    });
  };

  const iconTypes = {
    data_group: "collection",
    plug: "view-grid-add",
    resource: "cube",
  };

  load();
</script>

<div class="p-4">
  {#if loading}
    <LoadingSpinner />
  {:else}
    <div class="p-4  bg-white rounded">
      <div class="flex items-center justify-between ">
        <h4 class="font-semibold text-lg text-slate-800">Instance Bundle</h4>
      </div>

      <div class="space-y-2 mt-4 border">
        {#each bundle_objects["items"] || [] as item}
          <div
            class="flex space-x-4 bg-white p-3 shadow-sm hover:border border-blue-500 cursor-pointer"
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

      <div class="flex justify-end py-2">
        <button
          class="p-1 text-white text-sm font-semibold flex self-center shadow rounded bg-green-400 hover:bg-green-600"
        >
          <Icon name="lightning-bolt" class="h-5 w-5" />

          Start</button
        >
      </div>
    </div>
  {/if}
</div>
