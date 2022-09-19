<script lang="ts">
  import { getContext } from "svelte";

  import type { PortalApp } from "../../../../../../lib/app/portal";

  import InstallerDgroup from "./_installer_dgroup.svelte";
  import InstallPlug from "./_installer_plug.svelte";
  export let bid;

  const app: PortalApp = getContext("__app__");

  let data = {};
  let schema = {};
  let __loaded = false;
  let sources = [];

  const load = async () => {
    const bapi = await app.get_apm().get_bprint_api();
    const resp = await bapi.bprint_get(bid);

    if (resp.status !== 200) {
      console.log("Err loading", resp); // fixme toast error
      return;
    }
    data = resp.data;
    if (!data["schema"] || data["schema"] === "") {
      const sresp = await bapi.bprint_get_file(bid, "schema.json");
      if (sresp.status === 200) {
        schema = sresp.data;
      }
    } else {
      try {
        schema = JSON.parse(data["schema"]);
      } catch (error) {}
    }

    if (data["type"] === "tschema") {
      sources = await await app.get_dyn_sources();
    }

    __loaded = true;
  };

  load();
</script>

{#if __loaded}
  {#if data["type"] === "plug"}
    <InstallPlug {bid} {data} {schema} />
  {:else if data["type"] === "tschema"}
    <InstallerDgroup {bid} {data} {schema} {sources} />
  {:else}
    <div>Not Implemented</div>
  {/if}
{/if}
