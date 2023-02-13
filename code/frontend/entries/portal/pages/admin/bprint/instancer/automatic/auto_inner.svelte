<script lang="ts">
  import Icon from "@krowten/svelte-heroicons/Icon.svelte";
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../../core";

  export let bid;
  export let bundle_objects;
  export let instancer_type;

  const app = getContext("__app__") as PortalService;
  const bapi = app.api_manager.get_admin_bprint_api();

  let instanceing = false;
  let message = "";
  let instanced_resp = false;
  let instancedData;

  const instance = async () => {
    instanceing = true;
    const resp = await bapi.instance(bid, {
      auto: true,
      instancer_type,
      file: "schema.json",
    });

    if (!resp.ok) {
      message = resp.data;
      instanced_resp = true;
      return;
    }

    instancedData = resp.data;
    instanceing = false;
    instanced_resp = true;
  };

  const action_instance_manual = async (item) => {
    // const api = app.api_manager.get_admin_bprint_api();
    // const resp = await api.get(id);
    // if (!resp.ok) {
    //   console.log("@@");
    //   return;
    // }
    // const bprint = resp.data;
    // const file = bprint["files"].filter(
    //   (v) => v !== "schema.json" || v !== "schema.yaml"
    // )[0];
    // instance_helper(app, bprint["type"], bprint, file, InstanceBundlePicker);
  };

  const iconTypes = {
    data_group: "collection",
    plug: "view-grid-add",
    resource: "cube",
    data_sheet: "table",
  };
</script>

<div class="p-4  bg-white rounded">
  <div class="flex flex-col">
    <h4 class="font-semibold text-lg text-slate-800">Instance Blueprint</h4>

    <p class="text-red-500">{message}</p>
  </div>

  {#if instanceing}
    <LoadingSpinner />
  {:else if instanced_resp}
    <details>
      <summary>Response</summary>
      <pre><code>{JSON.stringify(instancedData, null, 2)}</code></pre>
    </details>
  {:else}
    <div class="space-y-2 mt-4 border">
      {#each bundle_objects["items"] || [] as item}
        <div
          class="flex justify-between space-x-4 bg-white p-3 shadow-sm hover:border border-blue-500 cursor-pointer"
        >
          <div class="flex">
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

          <div>
            <button
              class="text-white inline-flex rounded bg-blue-400 hover:bg-blue-700 p-1 text-sm"
              on:click={action_instance_manual}
            >
              <Icon name="play" class="w-5 h-5" />
              Manual
            </button>
          </div>
        </div>
      {/each}
    </div>

    <div class="flex justify-end py-2">
      <button
        on:click={instance}
        class="p-1 text-white font-semibold flex self-center shadow rounded bg-green-400 hover:bg-green-600 text-base"
      >
        <Icon name="lightning-bolt" class="h-6 w-6" />

        Start</button
      >
    </div>
  {/if}
</div>
