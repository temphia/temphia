import { SiteUtils, baseURL } from "../../../../lib/utils/site";
import { PortalService } from "./portal";

const build = () => {
  const site = new SiteUtils();

  if (!site.isLogged()) {
    // redirrect here
    console.error("Not logged")
    return null;
  }

  const adata = site.getAuthedData();

  return new PortalService({
    base_url: baseURL(),
    tenant_id: adata.tenant_id,
    user_token: adata.user_token,
    site_utils: site,
  });
};

export default build;
