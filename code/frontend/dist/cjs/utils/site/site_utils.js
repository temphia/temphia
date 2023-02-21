"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SiteUtils = void 0;
const AUTHED_KEY_PREFIX = "_temphia_authed_key_";
class SiteUtils {
    constructor(site_token) {
        this._site_token = site_token;
    }
    isLogged() {
        return !!this.get();
    }
    gotoLoginPage() {
        window.location.pathname = "/z/auth";
    }
    setAuthedData(data) {
        const pdata = JSON.stringify(data);
        this.set(pdata);
    }
    getAuthedData() {
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
    set(data) {
        localStorage.setItem(this.key(), data);
    }
    key() {
        // tenantify ?
        return AUTHED_KEY_PREFIX;
    }
}
exports.SiteUtils = SiteUtils;
