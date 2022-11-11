import type { SiteUtils } from "../../../../lib/utils/site";
import { ApiManager } from "./apm";
import { Launcher } from "../engine/launcher";
import { Navigator } from "./nav";
import { Notifier } from "../notifier";
import { DataService } from "../data";

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

  launcher: Launcher;
  data_service: DataService;

  constructor(opts: AppOptions) {
    this.options = opts;
    this.nav = new Navigator();
    this.api_manager = new ApiManager(
      opts.base_url,
      opts.tenant_id,
      opts.user_token
    );
    this.launcher = new Launcher();
  }

  async init() {
    await this.api_manager.init();
    this.init_notifier();
  }

  private async init_notifier() {
    const sapi = this.api_manager.get_api_sockd();

    this.notifier = new Notifier(this.api_manager.self_api);

    const sockd = await sapi.user(
      this.api_manager.user_token,
      this.notifier.handle_sockd
    );

    this.notifier.set_sockd(sockd);
    await this.notifier.init();
  }

  async get_data_service() {
    if (!this.data_service) {
      this.data_service = new DataService({
        apm: this.api_manager,
        sources: await this.api_manager.self_data.get_data_sources(),
      });
    }

    return this.data_service;
  }

  inject(utils: Utils) {
    this.utils = utils;
  }
}
