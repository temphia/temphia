import { AuthAPI } from "../../../lib/apiv2/auth";
import { apiURL, SiteUtils } from "../../../lib/utils/site";
import type {SiteData} from "../../../lib/utils/site";
import { AuthNav } from "./nav";


export class AuthService {
  auth_api: AuthAPI;
  nav: AuthNav;
  site_manager: SiteUtils;
  active_auth_id: number;

  _site_token: string;
  user_group: string;
  user_group_fixed: boolean;
  tenant_id: string;

  constructor() {
    const site_data: SiteData = window["__temphia_site_data__"];
    this.site_manager = new SiteUtils(site_data.site_token);
    this.auth_api = new AuthAPI(
      apiURL(site_data.tenant_id),
      site_data.site_token
    );
    this.active_auth_id = 0;

    this.user_group = site_data.user_group;
    this.user_group_fixed = false;
    this.tenant_id = site_data.tenant_id;
    this.nav = new AuthNav();

    this._site_token = site_data.site_token;
  }

  list_methods = async () => {
    const resp = await this.auth_api.list_methods(this.user_group);
    if (resp.status !== 200) {
      return null;
    }
    return {
      pass_auth: resp.data["pass_auth"],
      open_signup: resp.data["open_signup"],
      alt_auth_method: resp.data["alt_auth_method"],
    };
  };

  login_next = async (uid: string, pass: string) => {
    return this.auth_api.login_next({
      user_ident: uid,
      password: pass,
      site_token: this._site_token,
    });
  };

  login_submit = async (ntoken: string) => {
    return this.auth_api.login_submit({
      site_token: this._site_token,
      next_token: ntoken,
    });
  };

  login_finish = async (pre_auth_token: string, proof?: string) => {
    return this.auth_api.finish({
      site_token: this._site_token,
      preauthed_token: pre_auth_token,
      proof_token: proof,
    });
  };

  generate_alt_auth = async (id: number) => {
    this.active_auth_id = id;
    return this.auth_api.altauth_generate(id, {
      site_token: this._site_token,
      user_group: this.user_group,
    });
  };

  alt_next_first = async (code: string, state: string) => {
    return this.auth_api.altauth_next(this.active_auth_id, "first", {
      auth_code: code,
      auth_state: state,
      site_token: this._site_token,
      user_group: this.user_group,
    });
  };

  alt_next_second = async (first_token: string, signup_data?: object) => {
    return this.auth_api.altauth_next(this.active_auth_id, "second", {
      site_token: this._site_token,
      first_token: first_token,
      signup_data: signup_data,
    });
  };

  submit_alt_auth = async (next_token: string) => {
    return this.auth_api.altauth_submit(this.active_auth_id, {
      next_token: next_token,
      site_token: this._site_token,
    });
  };

  // preauthed data

  get_preauthed_data = () => {
    if (this.nav.nav_options) {
      return this.nav.nav_options;
    }
    const raw = localStorage.getItem("__pre_authed_data");
    return JSON.parse(raw);
  };

  save_preauthed_data = (data: any) => {
    localStorage.setItem("__pre_authed_data", JSON.stringify(data));
  };

  clear_preauthed_data = () => {
    localStorage.removeItem("__pre_authed_data");
  };

  // authed data

  save_authed_data = (user_token) => {
    this.site_manager.setAuthedData({
      site_token: this._site_token,
      tenant_id: this.tenant_id,
      user_token: user_token,
    });
  };

  clear_authed_data = () => {
    this.site_manager.clearAuthedData();
  };

  about = async () => {
    const adata = this.site_manager.getAuthedData();
    const resp = await this.auth_api.about(adata.user_token);
    if (!resp.ok) {
      return;
    }

    return resp.data;
  };
}
