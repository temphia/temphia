import { SelfAPI } from "../../../lib/apiv2";
import { ApiBase } from "../../../lib/apiv2/base";

export class ApiManager {
  base_url: string;
  api_base_url: string;

  tenant_id: string;
  user_token: string;
  session_token: string;

  cabinet_sources: string[];
  data_sources: string[];
  user_plugs: object[];

  base: ApiBase;
  setting_base: ApiBase;

  self_api: SelfAPI;

  constructor(base_url: string, tenant_id: string, user_token: string) {
    this.base_url = base_url;
    this.api_base_url = `${base_url}/z/api/${tenant_id}/v2`;
    this.tenant_id = tenant_id;
    this.user_token = user_token;
    this.cabinet_sources = [];
    this.data_sources = [];
    this.user_plugs = [];
  }

  async init() {
    const rresp = await fetch(`${this.api_base_url}/auth/refresh`, {
      body: JSON.stringify({
        user_token: this.user_token,
        options: {},
        old_token: this.session_token,
      }),
      method: "POST",
    });

    if (!rresp.ok) {
      console.log("@ERR", rresp);
      return;
    }

    const rdata = await rresp.json();
    if (!rdata["status_ok"]) {
      console.log("@ERR", rdata);
      return;
    }

    this.session_token = rdata["token"];

    this.base = new ApiBase(this.api_base_url, this.tenant_id, this.session_token);
    await this.base.init();

    this.self_api = new SelfAPI(this.base);

    const resp = await this.self_api.load();
    if (!resp.ok) {
      console.log("@ERR", resp);
      return;
    }

    const data = resp.data as SelfLoad;

    console.log("@data", data);
  }

  get_self_api() {
    this.self_api;
  }
}

interface SelfLoad {
  tenant_name: string;
  tenant_id: string;
  user_info: object;
  scopes: string[];
  plug_apps: object[];
}
