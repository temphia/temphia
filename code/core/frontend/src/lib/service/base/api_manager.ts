import type { Sockd } from "../../core/sockd";
import type {
  DynAPI,
  DtableAPI,
  CabinetAPI,
  PlugAPI,
  UserAPI,
  BprintAPI,
  ResourceAPI,
  TenantAPI,
  SelfAPI,
} from "../../core/api";
import { ApiBuilder, Options } from "../../core/builder";
import { sleep } from "yootils";
import { SockdService } from "../sockd";

export class ApiManager {
  _api_builder: ApiBuilder;
  _basic_api: SelfAPI;
  _sockd: Sockd;
  _sockd_muxer: SockdService;

  _dtable_apis: Map<string, DtableAPI>;
  _cabinet_apis: Map<string, CabinetAPI>;

  _dyn_api: DynAPI;
  _plug_api: PlugAPI;
  _user_api: UserAPI;
  //_engine_api: EngineAPI;
  _resource_api: ResourceAPI;
  _bprint_api: BprintAPI;
  _tenant_api: TenantAPI;

  constructor(opts: Options) {
    this._api_builder = new ApiBuilder(opts);
    this._dtable_apis = new Map();
    this._cabinet_apis = new Map();
  }

  init = async () => {
    this._basic_api = await this._api_builder.get_basic_api();
    this._sockd = await this._api_builder.get_sockd_api();
    await sleep(500);
    this._sockd_muxer = new SockdService(this._basic_api, this._sockd);
  };

  get_sockd_muxer = () => {
    return this._sockd_muxer;
  };

  async get_dyn_api(): Promise<DynAPI> {
    if (!this._dyn_api) {
      this._dyn_api = await this._api_builder.get_dyn_api();
    }
    return this._dyn_api;
  }

  get_basic_api(): SelfAPI {
    return this._basic_api;
  }

  async get_dtable_api(source: string, group: string): Promise<DtableAPI> {
    const key = `${source}__${group}`;
    if (!this._dtable_apis.has(key)) {
      this._dtable_apis.set(
        key,
        await this._api_builder.get_dtable_api(source, group)
      );
    }
    return this._dtable_apis.get(key);
  }

  async get_cabinet_api(source: string): Promise<CabinetAPI> {
    if (!this._cabinet_apis.has(source)) {
      this._cabinet_apis.set(
        source,
        await this._api_builder.get_cabinet_api(source)
      );
    }
    return this._cabinet_apis.get(source);
  }

  async get_plug_api(): Promise<PlugAPI> {
    if (!this._plug_api) {
      this._plug_api = await this._api_builder.get_plug_api();
    }
    return this._plug_api;
  }

  async get_user_api(): Promise<UserAPI> {
    if (!this._user_api) {
      this._user_api = await this._api_builder.get_user_api();
    }
    return this._user_api;
  }

  async get_bprint_api(): Promise<BprintAPI> {
    if (!this._bprint_api) {
      this._bprint_api = await this._api_builder.get_bprint_api();
    }
    return this._bprint_api;
  }

  async get_resource_api(): Promise<ResourceAPI> {
    if (!this._resource_api) {
      this._resource_api = await this._api_builder.get_resource_api();
    }
    return this._resource_api;
  }

  /*
    async get_engine_api(): Promise<EngineAPI> {
        if (!this._engine_api) {
            this._engine_api = await this._api_builder.get_engine_api()
        }
        return this._engine_api
    }
    */

  async get_tenant_id(): Promise<TenantAPI> {
    if (!this._tenant_api) {
      this._tenant_api = await this._api_builder.get_tenant_api();
    }
    return this._tenant_api;
  }
}
