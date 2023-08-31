const AUTHED_KEY_PREFIX = "_temphia_authed_key_";

export interface AuthedData {
  user_token: string;
  tenant_id?: string;
}

export interface SiteData {
  tenant_id: string;
  site_token: string;
  user_group?: string
}

export class SiteUtils {
  isLogged() {
    return !!this.get();
  }

  gotoLoginPage() {
    window.location.pathname = "/z/pages/auth";
  }

  gotoPortalPage() {
    window.location.pathname = "/z/pages/portal";
  }

  setAuthedData(data: AuthedData) {
    const pdata = JSON.stringify(data);
    this.set(pdata);
  }

  getAuthedData(): AuthedData | null {
    const raw = this.get();
    if (raw === null) {
      return null
    }

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
