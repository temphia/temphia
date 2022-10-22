import { SiteUtils, baseURL } from "../../../lib/utils/site";
import { App } from "./app";

const build = () => {
  const site = new SiteUtils();

  if (!site.isLogged()) {
    // redirrect here
    console.error("Not logged")
    return null;
  }

  const adata = site.getAuthedData();

  return new App({
    base_url: baseURL(),
    tenant_id: adata.tenant_id,
    user_token: adata.user_token,
    site_utils: site,
  });
};

export default build;
