import { AdapterEditorEnv } from "../../lib/adapter/adapter";
import { AdapterEditorAPI } from "../../lib/apiv2/admin/adapter_editor";
import { initRegistry } from "../../lib/engine/putils";
import type { Registry } from "../../lib/registry/registry";
initRegistry();

window.addEventListener("load", (ev) => {
  const registry = window["__registry__"] as Registry<any>;
  const loaderOpts = window["__loader_options__"] || {};

  const env = new AdapterEditorEnv({
    api: new AdapterEditorAPI(
      loaderOpts["base_url"],
      loaderOpts["tenant_id"],
      loaderOpts["adapter_editor_token"]
    ),
  });
  

  console.log("@adapter_editor_loader", loaderOpts, env, registry);
});
