import { CabinetAPI, DataAPI, RepoAPI, SelfAPI } from "../../../../lib/apiv2";
import { AdminTargetAPI } from "../../../../lib/apiv2/admin/target";
import { ApiBase } from "../../../../lib/apiv2/base";
import type { SelfLoad } from "./stypes";

import {
  AdminBprintAPI,
  AdminCheckAPI,
  AdminDataAPI,
  AdminLensAPI,
  AdminPlugAPI,
  AdminRepoAPI,
  AdminResourceAPI,
  AdminTenantAPI,
  AdminUserAPI,
  AdminUserGroupAPI,
} from "../../../../lib/apiv2/admin";
import { PlugDevTktAPI } from "../../../../lib/apiv2/engine/plug_dev_tkt";
import { EngineAPI } from "../../../../lib/apiv2/engine/engine";

export class ApiManager {
  base_url: string;
  api_base_url: string;

  tenant_id: string;
  user_token: string;
  session_token: string;

  base: ApiBase;
  self_api: SelfAPI;
  self_data: SelfData;

  constructor(base_url: string, tenant_id: string, user_token: string) {
    this.base_url = base_url;
    this.api_base_url = `${base_url}/z/api/${tenant_id}/v2`;
    this.tenant_id = tenant_id;
    this.user_token = user_token;
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
    this.self_data = new SelfData(this, resp.data as SelfLoad);
  }

  // api

  get_self_api() {
    return this.self_api;
  }

  get_repo_api = () => {
    return new RepoAPI(this.base);
  };

  get_engine_api = () => {
    return new EngineAPI(this.base);
  };

  get_cabinet = (source: string) => {
    return new CabinetAPI(source, this.base);
  };

  get_data_api = async (source: string, group: string) => {
    const data = {
      source: source,
      group: group,
    };

    console.log("##", data);

    const resp = await this.self_api.issue_data(data);
    if (!resp.ok) {
      console.log("@@err", resp);
      return;
    }

    return new DataAPI(this.api_base_url, resp.data["data_token"]);
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

  get_admin_lens_api = () => {
    return new AdminLensAPI(this.base);
  };

  get_dev_plug_tkt_api = (dev_ticket: string) => {
    return new PlugDevTktAPI(this.api_base_url, dev_ticket);
  };
}

export class SelfData {
  data_sources: string[];
  repo_sources: { [_: number]: string };
  user_plugs: object[];

  api_manager: ApiManager;

  constructor(apm: ApiManager, data: SelfLoad) {
    this.api_manager = apm;
    this.user_plugs = data["plug_apps"] || [];
  }

  async get_repo_sources() {
    if (this.repo_sources) {
      return this.repo_sources;
    }
    const resp = await this.api_manager.self_api.list_repo_sources();
    if (!resp.ok) {
      return [];
    }

    this.repo_sources = resp.data;
    return this.repo_sources;
  }


  async get_data_sources() {
    if (this.data_sources) {
      return this.data_sources;
    }

    const resp = await this.api_manager.self_api.list_data_sources();
    if (!resp.ok) {
      return [];
    }
    this.data_sources = resp.data;
    return this.data_sources;
  }

  get_user_plugs() {
    if (this.user_plugs) {
      return this.user_plugs;
    }

    return [];
  }
}
