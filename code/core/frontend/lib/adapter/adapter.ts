import type { AdapterEditorAPI } from "../apiv2/admin/adapter_editor";

export class AdapterEditorEnv {
  api: AdapterEditorAPI;
  constructor({ api }) {
    this.api = api;
  }
}
