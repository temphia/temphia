import {
  SelfAPI,
  BprintAPI,
  CabinetAPI,
  PlugAPI,
  ResourceAPI,
  TenantAPI,
  UserAPI,
  DynAPI,
  DtableAPI,
} from "../api";
import { Sockd } from "../sockd";

export interface Options {
  api_url: string;
  user_token: string;
  tenant_id: string;
}

export class ApiBuilder {
  _api_url: string;
  _user_token: string;
  _tenant_id: string;
  _basic_api: SelfAPI;
  _admin_user_api: UserAPI; // change this to per user group

  constructor(opts: Options) {
    this._api_url = opts.api_url;
    this._user_token = opts.user_token;
    this._tenant_id = opts.tenant_id;
  }

  async get_sockd_api(): Promise<Sockd> {
    const sockd = new Sockd(
      `${
        this._api_url
      }/self/user_ws?token=${this._basic_api.get_session_token()}`
    );
    await sockd.init();
    return sockd;
  }

  async get_dyn_api(): Promise<DynAPI> {
    const sapi = new DynAPI(this._api_url, this._user_token);
    await sapi.init();
    return sapi;
  }

  async get_basic_api(): Promise<SelfAPI> {
    const basic_api = new SelfAPI(this._api_url, this._user_token);
    await basic_api.init();
    this._basic_api = basic_api;
    return basic_api;
  }

  async get_dtable_api(source: string, group: string): Promise<DtableAPI> {
    const sapi = new DtableAPI(
      this._api_url,
      this._user_token,
      source,
      group
    );
    await sapi.init();
    return sapi;
  }

  async get_cabinet_api(source: string): Promise<CabinetAPI> {
    const sapi = new CabinetAPI(this._api_url, this._user_token, source);
    await sapi.init();
    return sapi;
  }

  async get_plug_api(): Promise<PlugAPI> {
    const sapi = new PlugAPI(this._api_url, this._user_token);
    await sapi.init();
    return sapi;
  }

  async get_user_api(): Promise<UserAPI> {
    const sapi = new UserAPI(this._api_url, this._user_token);
    await sapi.init();
    return sapi;
  }


  async get_bprint_api(): Promise<BprintAPI> {
    const bapi = new BprintAPI(this._api_url, this._user_token);
    await bapi.init();
    return bapi;
  }

  async get_resource_api(): Promise<ResourceAPI> {
    const rapi = new ResourceAPI(this._api_url, this._user_token);
    await rapi.init();
    return rapi;
  }

  // async get_engine_api(): Promise<EngineAPI> {
  //     const eapi = new EngineAPI(this._api_base_url, this._user_token)
  //     await eapi.init()
  //     return eapi
  // }

  async get_tenant_api(): Promise<TenantAPI> {
    const tapi = new TenantAPI(this._api_url, this._user_token);
    await tapi.init();
    return tapi;
  }
}
