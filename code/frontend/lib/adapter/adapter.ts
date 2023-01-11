import type { AdapterEditorAPI } from "../apiv2/admin/adapter_editor";

export class AdapterEditorEnv {
  api: AdapterEditorAPI;
  domain_name: string;
  constructor( api: AdapterEditorAPI, domain_name: string) {
    this.api = api;
    this.domain_name = domain_name;
  }
}
