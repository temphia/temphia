import { writable, Writable } from "svelte/store";
import type { SiteUtils } from "../../../lib/utils/site";
import { ApiManager } from "./apm";
import { Navigator } from "./nav";
import { Notifier } from "./notifier";

export interface AppOptions {
  base_url: string;
  tenant_id: string;
  user_token: string;
  site_utils: SiteUtils;
}

export interface Utils {
  toast_success(msg: string): void;
  toast_error(msg: string): void;
  big_modal_open(compo: any, options: object): void;
  big_modal_close(): void;
  small_modal_open(compo: any, options: object): void;
  small_modal_close(): void;
  notification_toggle(): void;
}

export class PortalService {
  options: AppOptions;
  nav: Navigator;
  api_manager: ApiManager;

  notifier: Notifier;
  utils: Utils;
  launcher_active: Writable<boolean>;

  constructor(opts: AppOptions) {
    this.options = opts;
    this.nav = new Navigator();
    this.api_manager = new ApiManager(
      opts.base_url,
      opts.tenant_id,
      opts.user_token
    );
    this.launcher_active = writable(false);
  }

  async init() {
    await this.api_manager.init();

    this.notifier = new Notifier(this.api_manager.self_api);
    await this.notifier.init();
  }

  inject(utils: Utils) {
    this.utils = utils;
  }
}
