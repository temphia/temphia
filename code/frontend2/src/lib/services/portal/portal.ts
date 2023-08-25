import type { SiteUtils } from "../../utils/site";
import { ApiManager } from "./apm";
import { Launcher } from "./launcher/launcher";
import { Navigator } from "./nav";
import { Notifier } from "./notifier";

import { DataService } from "../data";

import { SockdService } from "./sockd/sockd";
import type { Logger } from "./logger";
import { CabinetService } from "../cabinet/cabinet";
import { XtMgr } from "./xtmgr";

export interface AppOptions {
  base_url: string;
  tenant_id: string;
  user_token: string;
  site_utils: SiteUtils;
  registry: any;
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
  sockd_service: SockdService;
  notifier?: Notifier;
  utils: Utils;
  registry: any;
  logger: Logger;
  xtmgr: XtMgr;

  services: Map<string, any>


  launcher: Launcher;
  data_service: DataService;
  cabinet_service: CabinetService;

  constructor(opts: AppOptions) {
    console.log("@portal_service", this);

    this.options = opts;
    this.nav = new Navigator();
    this.api_manager = new ApiManager(
      opts.base_url,
      opts.tenant_id,
      opts.user_token
    );

    this.services = new Map()


    this.launcher = new Launcher();
    this.sockd_service = new SockdService();
    this.cabinet_service = new CabinetService(this.api_manager);
    this.registry = opts.registry;
    this.xtmgr = new XtMgr(this, this.registry);
  }

  async init() {
    const resp = await this.api_manager.init();
    if (resp) {
      return resp;
    }
    this.init_notifier();
    this.xtmgr.init();
  }

  private async init_notifier() {
    const wsurl = this.api_manager.self_api.self_ws_url();

    this.notifier = new Notifier(this.api_manager.self_api);

    const sockd = await this.sockd_service.build(
      wsurl,
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
        close_modal: this.utils.big_modal_close,
        open_modal: this.utils.big_modal_open,
        api_base_url: this.api_manager.api_base_url,
        sockd_builder: this.sockd_service,
        profile_genrator: this.get_user_profile,
      });
    }

    return this.data_service;
  }

  inject(utils: Utils) {
    this.utils = utils;
    this.notifier.toast_open = utils.toast_success;
  }

  get_sockd_service = () => {
    return this.sockd_service;
  };

  get_cabinet_service = () => {
    return this.cabinet_service;
  };

  get_user_profile = (name) => {
    return `${this.options.base_url}/z/assets/system/${this.options.tenant_id}/user/${name}`;
  };

  logout = () => {
    this.options.site_utils.clearAuthedData();
    location.href = "/z/auth";
  };
}
