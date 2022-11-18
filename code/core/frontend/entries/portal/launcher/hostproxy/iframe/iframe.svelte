<script lang="ts">
  import type { Instance } from "../../../services/engine/launcher";
  import LoadingSpinner from "../../../../xcompo/common/loading_spinner.svelte";
  import IframeInner from "./iframe_inner.svelte";
  import type { PortalService } from "../../../services";
  import { getContext } from "svelte";
  import { iframeTemplateBuild } from "../../../../../lib/engine/template";

  export let options: Instance;

  const app = getContext("__app__") as PortalService;
  const papi = app.api_manager.get_admin_plug_api();

  let loading = true;
  let agent = {};

  const load = async () => {
    const resp = await papi.get_agent(options.plug_id, options.agent_id);
    if (!resp.ok) {
      return;
    }
    agent = resp.data;
  };

  const src = iframeTemplateBuild({
    plug: options.plug_id,
    agent: options.agent_id,
    base_url: this._engine_data["base_url"],
    entry_name: this._engine_data["entry"],
    exec_loader: this._engine_data["exec_loader"],
    js_plug_script: this._engine_data["js_plug_script"],
    style_file: this._engine_data["style"],
    token: this._engine_data["token"] || "",
    ext_scripts: this._engine_data["ext_scripts"],
    parent_secret: this._secret,
    startup_payload: {},
  });

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <IframeInner {options} {agent} />
{/if}
