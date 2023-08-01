import { AdapterEditorEnv } from "../../../lib/adapter/adapter";
import { AdapterEditorAPI } from "../../../lib/apiv2/admin/adapter_editor";
import { initRegistry } from "../../../lib/engine/putils";
import type { Registry } from "../../../lib/registry/registry";
initRegistry();

window.addEventListener("load", async (ev) => {
  const registry = window["__registry__"] as Registry<any>;
  const loaderOpts = window["__loader_options__"] || {};

  const api = new AdapterEditorAPI(
    loaderOpts["base_url"],
    loaderOpts["tenant_id"],
    loaderOpts["adapter_editor_token"]
  );

  const env = new AdapterEditorEnv(api, loaderOpts["domain_name"] || "");

  console.log("@adapter_editor_loader", loaderOpts, env, registry);

  const adapterType = loaderOpts["adapter_type"] || "";
  let factory = registry.Get(
    "temphia.adapter_editor.loader",
    `${adapterType}.main`
  );

  if (!factory) {
    await registry.WatchLoad(
      "temphia.adapter_editor.loader",
      `${adapterType}.main`,
      20000
    );

    factory = registry.Get(
      "temphia.adapter_editor.loader",
      `${adapterType}.main`
    );
  }

  if (factory) {
    factory({
      env: env,
      target: document.getElementById("adapter-editor-root"),
    });
  } else {
    console.warn("@adapter_editor_loader", "factory not found");
  }
});
