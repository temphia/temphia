import type { AdapterEditorEnv } from "../../../lib/adapter/adapter";
import type { AdapterEditorAPI } from "../../../lib/apiv2/admin/adapter_editor";

export class EasypageService {
  env: AdapterEditorEnv;
  api: AdapterEditorAPI;
  constructor(env: AdapterEditorEnv) {
    this.api = env.api;
    this.env = env;
  }
}
