export * from "./app";

import { apiURL, baseURL, SiteManager } from "../../core/site";
import { PortalApp } from "./app";

export const buildApp = (modal_open: any, modal_close: any, toaster: any) => {
  const site = new SiteManager();

  if (!site.isLogged()) {
    // redirrect here
    return null;
  }

  const data = site.getAuthedData();

  const __app = new PortalApp({
    api_url: apiURL(data.tenant_id),
    site_token: data.site_token,
    tenant_id: data.tenant_id,
    url_base:baseURL(),
    user_token: data.user_token,
    simple_modal_open: modal_open,
    simple_modal_close: modal_close,
    toaster: toaster,
    siteman: site,
  });
  return __app;
};
