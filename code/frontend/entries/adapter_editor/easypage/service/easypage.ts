import type { AdapterEditorEnv } from "../../../../lib/adapter/adapter";
import type { AdapterEditorAPI } from "../../../../lib/apiv2/admin/adapter_editor";

export class EasypageService {
  env: AdapterEditorEnv;
  api: AdapterEditorAPI;
  modal: {
    big_open: any;
    big_close: any;
    small_open: any;
    small_close: any;
  };

  constructor(env: AdapterEditorEnv) {
    this.api = env.api;
    this.env = env;
  }

  load = () => {
    return this.api.perform_action("load", null);
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
