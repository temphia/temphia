<script lang="ts">
  import ImportV2 from "./_import_v2.svelte";
  import type { PortalApp } from "../../../../../../lib/app/portal";

  export let data;
  export let group;
  export let source;
  export let app: PortalApp


  const importBprint = async (_data: any) => {
    const bapi = await app.get_apm().get_bprint_api();

    return bapi.bprint_import(_data);
  };
</script>

<ImportV2 importFunc={importBprint} {data} {group} {source} />
