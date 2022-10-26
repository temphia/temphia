<script lang="ts">
  import type { PortalService } from "../../../services";
  import Importer from "./_importer_impl.svelte";

  export let data;
  export let group;
  export let source;
  export let app: PortalService;

  const importBprint = async (_data: any) => {
    console.log("@import", _data);
    //   const bapi = await app.get_apm().get_bprint_api();
    //   return bapi.bprint_import(_data);
  };
</script>

<Importer importFunc={importBprint} {data} {group} {source} />
