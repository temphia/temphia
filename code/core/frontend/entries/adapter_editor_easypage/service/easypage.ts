import type { AdapterEditorEnv } from "../../../lib/adapter/adapter";
import type { AdapterEditorAPI } from "../../../lib/apiv2/admin/adapter_editor";

export class EasypageService {
  env: AdapterEditorEnv;
  api: AdapterEditorAPI;
  constructor(env: AdapterEditorEnv) {
    this.api = env.api;
    this.env = env;
  }

  listPage = () => {
    return this.api.perform_action("list_pages", null);
  };

  updatePages = (data: any[]) => {
    return this.api.perform_action("update_pages", data);
  };

  getPageData = (slug: string) => {
    return this.api.perform_action("get_page_data", slug);
  };

  setPageData = (slug: string, data: string) => {
    return this.api.perform_action("set_page_data", {
      slug,
      data,
    });
  };

  deletePageData = (slug: string) => {
    return this.api.perform_action("delete_page_data", slug);
  };
}
