<script lang="ts">
  import type { Instance } from "../../../services/engine/launcher";
  import LoadingSpinner from "../../../../xcompo/common/loading_spinner.svelte";
  import IframeInner from "./iframe_execute.svelte";
  import type { PortalService } from "../../../services";
  import { getContext } from "svelte";
  import type { ExecInstanceOptions } from "../../../services/engine/exec_type";

  export let options: Instance;

  const app = getContext("__app__") as PortalService;
  const eapi = app.api_manager.get_engine_api();

  let loading = true;
  let data: ExecInstanceOptions;

  const load = async () => {
    const resp = await eapi.launch_target({
      target_id: Number(options.target_id),
      target_type: options.target_type,
    });
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    data = resp.data;
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <IframeInner exec_data={data} name={options.name} secret_id={options.id} />
{/if}
