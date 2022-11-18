<script lang="ts">
  import type { Instance } from "../../../services/engine/launcher";
  import LoadingSpinner from "../../../../xcompo/common/loading_spinner.svelte";
  import IframeInner from "./iframe_inner.svelte";
  import type { PortalService } from "../../../services";
  import { getContext } from "svelte";
  import { iframeTemplateBuild } from "../../../../../lib/engine/template";

  export let options: Instance;

  const app = getContext("__app__") as PortalService;
  const eapi = app.api_manager.get_engine_api();

  let loading = true;
  let src = "";

  const load = async () => {
    const resp = await eapi.launch_target({
      target_id: 0,
      target_type: "",
    });
    if (!resp.ok) {
      console.log("Err", resp);
      return;
    }
    const data = resp.data;
    src = iframeTemplateBuild({
      plug: data["plug"],
      agent: data["agent"],
      base_url: data["base_url"],
      entry_name: data["entry"],
      exec_loader: data["exec_loader"],
      js_plug_script: data["js_plug_script"],
      style_file: data["style"],
      token: data["token"] || "",
      ext_scripts: data["ext_scripts"],
      parent_secret: this._secret,
      startup_payload: {},
    });
    loading = false;
  };

  load();
</script>

{#if loading}
  <LoadingSpinner />
{:else}
  <IframeInner {options} source={src} />
{/if}
