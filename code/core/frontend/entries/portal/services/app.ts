import type { SiteUtils } from "../../../lib/utils/site";
import { ApiManager } from "./apm";
import { Navigator } from "./nav";
import type { Notifier } from "./notifier";

export interface AppOptions {
  base_url: string;
  tenant_id: string;
  user_token: string;
  site_utils: SiteUtils;
}

export class App {
  options: AppOptions;

  nav: Navigator;
  toaster: any;

  api_manager: ApiManager;
  notifier: Notifier;

  constructor(opts: AppOptions) {
    this.options = opts;
    this.nav = new Navigator();
  }

  async init() {
    this.api_manager = new ApiManager(
      this.options.base_url,
      this.options.tenant_id,
      this.options.user_token
    );
    await this.api_manager.init();
  }

  get_apm() {
    return this.api_manager;
  }
}
