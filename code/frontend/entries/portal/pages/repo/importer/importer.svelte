<script lang="ts">
  import type { PortalService } from "../../../services";
  import Importer from "./_importer_impl.svelte";

  export let data;
  export let group;
  export let source;
  export let app: PortalService;

  const importBprint = async (_data: any) => {
    const bapi = app.api_manager.get_admin_bprint_api();
    const resp = await bapi.import(_data);
    if (!resp.ok) {
      console.log("@@ERR", resp);
    }

    app.utils.small_modal_close();
  };
</script>

<Importer importFunc={importBprint} {data} {group} {source} />
