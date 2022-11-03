import { RepoAPI, SelfAPI } from "../../../lib/apiv2";
import {
  AdminBprintAPI,
  AdminCheckAPI,
  AdminDataAPI,
  AdminPlugAPI,
  AdminRepoAPI,
  AdminResourceAPI,
  AdminTenantAPI,
  AdminUserAPI,
  AdminUserGroupAPI,
} from "../../../lib/apiv2/admin";
import { AdminTargetAPI } from "../../../lib/apiv2/admin/target";
import { ApiBase } from "../../../lib/apiv2/base";
import type { SelfLoad } from "./stypes";

export class ApiManager {
  base_url: string;
  api_base_url: string;

  tenant_id: string;
  user_token: string;
  session_token: string;

  cabinet_sources: string[];
  data_sources: string[];
  repo_sources: { [_: number]: string };
  user_plugs: object[];

  base: ApiBase;

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

    this.base = new ApiBase(
      this.api_base_url,
      this.tenant_id,
      this.session_token
    );
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

  async get_repo_sources() {
    if (this.repo_sources) {
      return this.repo_sources;
    }

    const resp = await this.self_api.list_repo_sources();
    this.repo_sources = resp.data;
    return this.repo_sources;
  }

  // api

  get_self_api() {
    this.self_api;
  }

  get_repo_api = () => {
    return new RepoAPI(this.base);
  };

  // admin api

  get_admin_bprint_api = () => {
    return new AdminBprintAPI(this.base);
  };

  get_admin_plug_api = () => {
    return new AdminPlugAPI(this.base);
  };

  get_admin_repo_api = () => {
    return new AdminRepoAPI(this.base);
  };

  get_admin_data_api = () => {
    return new AdminDataAPI(this.base);
  };

  get_admin_resource_api = () => {
    return new AdminResourceAPI(this.base);
  };

  get_admin_target_api = () => {
    return new AdminTargetAPI(this.base);
  };

  get_admin_tenant_api = () => {
    return new AdminTenantAPI(this.base);
  };

  get_admin_ugroup_api = () => {
    return new AdminUserGroupAPI(this.base);
  };

  get_admin_user_api = () => {
    return new AdminUserAPI(this.base);
  };

  get_admin_check_api = () => {
    return new AdminCheckAPI(this.base);
  };
}
