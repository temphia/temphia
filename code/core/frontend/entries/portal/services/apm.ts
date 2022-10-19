import { SelfAPI } from "../../../lib/apiv2";
import { ApiBase } from "../../../lib/apiv2/base";

export class ApiManager {
  base_url: string;
  tenant_id: string;
  user_token: string;

  cabinet_sources: string[];
  data_sources: string[];
  user_plugs: object[];

  base: ApiBase;
  setting_base: ApiBase;

  self_api: SelfAPI;

  constructor(base_url: string, tenant_id: string, user_token: string) {
    this.base_url = base_url;
    this.tenant_id = tenant_id;
    this.user_token = user_token;
    this.cabinet_sources = [];
    this.data_sources = [];
    this.user_plugs = [];
  }

  async init() {
    this.base = new ApiBase(this.base_url, this.tenant_id, this.user_token);
    await this.base.init()

    this.self_api = new SelfAPI(this.base);

    this.self_api.load()

  }

  get_self_api() {
    this.self_api;
  }
}
