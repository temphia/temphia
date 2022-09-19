import { FolderTktAPI } from "../../core/tktapi";
import { DataGroupService } from "../dyn";
import { ApiManager } from "./api_manager";

export interface AppBaseOptions {
  url_base: string;
  api_url: string;
  tenant_id: string;
  site_token: string;
  user_token: string;
}

export class AppBase {
  _url_base: string;
  _tenant_id: string;
  _site_token: string;
  _user_token: string;
  _api_url: string;
  apm: ApiManager;

  constructor(opts: AppBaseOptions) {
    this._url_base = opts.url_base;
    this._api_url = opts.api_url;
    this._tenant_id = opts.tenant_id;
    this._site_token = opts.site_token;
    this._user_token = opts.user_token;
  }

  async init() {
    await this.build_api_manager();
  }

  async build_api_manager() {
    this.apm = new ApiManager({
      api_url: this._api_url,
      tenant_id: this._tenant_id,
      user_token: this._user_token,
    });
    await this.apm.init();
  }

  get_data_service = async (source: string, group: string) => {
    const dapi = await this.apm.get_dtable_api(source, group);
    const svc = new DataGroupService({
      dapi,
      group: group,
      source: source,
      engine_service: null, // fixme
      sockd_service: null, //this.apm.get_sockd_muxer()
    });

    await svc.init();
    return svc;
  };

  get_dyn_sources = async () => {
    const bapi = this.apm.get_basic_api();
    const resp = await bapi.list_dgroup_sources();
    if (resp.status !== 200) {
      console.log("Err loading dyn sources", resp);
      return [];
    }
    return resp.data;
  };

  get_cabinet_sources = async () => {
    const bapi = this.apm.get_basic_api();
    const resp = await bapi.list_cabinet_sources();
    if (resp.status !== 200) {
      console.log("Err loading cabinet sources", resp);
      return [];
    }
    return resp.data;
  };

  get_folder_api = async (source: string, folder: string) => {
    const capi = await this.apm.get_cabinet_api(source);
    const fresp = await capi.get_folder_ticket(folder);
    return new FolderTktAPI(this._api_url, fresp.data);
  };

  user_profile_image_link = (user_id: string) => {
    return `${this._url_base}/api/${this._tenant_id}/v1/user_profile_image/${user_id}`;
  };

  log_out() {}
}
