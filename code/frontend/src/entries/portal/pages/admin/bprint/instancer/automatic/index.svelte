<script lang="ts">
  import { getContext } from "svelte";
  import { LoadingSpinner, PortalService } from "../../../../admin/core";

  import { params } from "svelte-hash-router";
  import InstancerInner from "./_instancer.svelte";

  export let bid = $params.bid;

  const app = getContext("__app__") as PortalService;
  const bapi = app.api_manager.get_admin_bprint_api();

  let bundle_objects = {};
  let loading = true;
  let instancer_type = "";

  const load = async () => {
    const resp1 = await bapi.get(bid);
    if (!resp1.ok) {
      console.log("err", resp1);
      return;
    }

    // fixme change schema.json -> install.json

    if (resp1.data["type"] === "bundle") {
      const resp = await bapi.get_file(bid, "schema.json");
      if (!resp.ok) {
        console.log("Err", resp);
        return;
      }

      instancer_type = "bundle";
      bundle_objects = resp.data;
      loading = false;
    } else {
      let btype = resp1.data["type"];
      if (btype === "tschema") {
        btype = "data_group";
      }
      instancer_type = btype;
      bundle_objects = {
        items: [
          {
            type: btype,
            file: "schema.json",
            name: resp1.data["name"],
          },
        ],
      };
      loading = false;
    }
  };

  load();
</script>

<div class="p-4">
  {#if loading}
    <LoadingSpinner />
  {:else}
    <InstancerInner {bid} {bundle_objects} {instancer_type} />
  {/if}
</div>
