import type { SiteManager } from "../../core/site";
import type { FolderTktAPI } from "../../core/tktapi";
import { AppBase } from "../../service/base";
import type { DataGroupService } from "../../service/dyn";
import { EngineLauncher } from "./elauncher";
import { Navigator } from "./navigator";
import { Notifier } from "./notifier";


export interface Toaster {
  success(message: string): void;
  error(message: string): void;
}

declare global {
  interface Window {
    showModal(c: any, p: any);
    closeModal();
  }
}

export interface PortalAppOptions {
  url_base: string;
  api_url: string;
  tenant_id: string;
  site_token: string;
  user_token: string;
  simple_modal_close: any;
  simple_modal_open: any;
  toaster: Toaster;
  siteman: SiteManager
}

export class PortalApp {
  _base_app: AppBase;
  navigator: Navigator;
  notifier: Notifier;
  toaster: any;
  elauncher: EngineLauncher
  siteman: SiteManager

  _simple_modal_open: any;
  _simple_modal_close: any;

  _current_data_service: DataGroupService;

  _cabinet_sources: string[];
  _dyn_sources: string[];
  _folder_tickets: Map<string, FolderTktAPI>;
  _store_sources: string[];
  _quick_apps: object[];

  constructor(opts: PortalAppOptions) {
    this.navigator = new Navigator(opts.url_base);
    this._base_app = new AppBase({
      api_url: opts.api_url,
      site_token: opts.site_token,
      tenant_id: opts.tenant_id,
      url_base: opts.url_base,
      user_token: opts.user_token,
    });

    this._simple_modal_open = opts.simple_modal_open;
    this._simple_modal_close = opts.simple_modal_close;
    this._folder_tickets = new Map();
    this.elauncher = new EngineLauncher(this._base_app)
    this.siteman = opts.siteman
  }
  async init() {
    await this._base_app.init();
    const bapi = this._base_app.apm.get_basic_api();

    this.notifier = new Notifier({
      basicAPI: bapi,
      sockdMuxer: this._base_app.apm.get_sockd_muxer(),
    });
  }

  get_apm() {
    return this._base_app.apm;
  }

  simple_modal_open = (compo: any, opts: any) => {
    this._simple_modal_open(compo, opts);
  };

  simple_modal_close = () => {
    this._simple_modal_close();
  };

  big_modal_open = (_compo, _props) => {
    window.showModal(_compo, _props);
  };

  big_modal_close = () => {
    window.closeModal();
  };

  user_profile_image_link = (user_id: string) => {
    return this._base_app.user_profile_image_link(user_id);
  };

  log_out = () => {
    this.siteman.clearAuthedData();
    this.siteman.gotoLoginPage()
  }

  is_mobile = () => {
    return screen.width < 700;
  };

  get_data_service = async (source: string, group: string) => {
    if (this._current_data_service) {
      if (
        this._current_data_service.source === source &&
        this._current_data_service.group === group
      ) {
        return this._current_data_service;
      }
      await this._current_data_service.close();
    }
    this._current_data_service = await this._base_app.get_data_service(
      source,
      group
    );

    return this._current_data_service;
  };

  get_dyn_sources = async () => {
    if (!this._dyn_sources) {
      this._dyn_sources = await this._base_app.get_dyn_sources();
    }

    return this._dyn_sources;
  };

  get_cabinet_sources = async () => {
    if (!this._cabinet_sources) {
      this._cabinet_sources = await this._base_app.get_cabinet_sources();
    }
    return this._cabinet_sources;
  };

  get_store_sources = async () => {
    if (this._store_sources) {
      return this._store_sources;
    }

    const bapi = await this._base_app.apm.get_bprint_api();
    const resp = await bapi.repo_sources();

    if (resp.status !== 200) {
      return;
    }

    this._store_sources = resp.data;
    return this._store_sources;
  };

  get_folder_api = async (source: string, folder: string) => {
    const key = `${source}__${folder}`;

    if (!this._folder_tickets.has(key)) {
      const fapi = await this._base_app.get_folder_api(source, folder);
      this._folder_tickets.set(key, fapi);
    }
    return this._folder_tickets.get(key);
  };
}
