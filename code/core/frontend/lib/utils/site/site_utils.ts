const AUTHED_KEY_PREFIX = "_temphia_authed_key_";

export interface AuthedData {
  user_token: string;
  site_token: string;
  tenant_id: string;
  // fixme => place claim expiry date ?
}

export interface SiteData {
  tenant_id: string;
  site_token: string;
}

export class SiteUtils {
  _site_token: string;

  constructor(site_token?: string) {
    this._site_token = site_token;
  }

  isLogged() {
    return !!this.get();
  }

  gotoLoginPage() {
    window.location.pathname = "/z/auth";
  }

  setAuthedData(data: AuthedData) {
    const pdata = JSON.stringify(data);
    this.set(pdata);
  }

  getAuthedData(): AuthedData {
    const raw = this.get();
    const data = JSON.parse(raw);
    return data;
  }

  clearAuthedData() {
    localStorage.removeItem(this.key());
  }

  get() {
    return localStorage.getItem(this.key());
  }

  set(data: string) {
    localStorage.setItem(this.key(), data);
  }

  private key() {
    // tenantify ?
    return AUTHED_KEY_PREFIX;
  }
}
