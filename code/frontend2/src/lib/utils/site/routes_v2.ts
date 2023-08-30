
//http://localhost:4000
export const baseURL = () => window.location.origin;
//http://localhost:4000/z/api/:tenant_id/v2
export const apiURL = (tenant_id: string) =>
  `${baseURL()}/z/api/${tenant_id}/v2`;

export const portalPage = () => `${baseURL()}/z/pages/portal`;
export const authURL = () => `${baseURL()}/z/auth`;

export const authPage = (opts?: { tenant_id?: string; user_group?: string }) => {
  if (!opts) {
    return `${baseURL()}/z/pages/auth`;
  }

  return `${baseURL()}/z/pages/auth?${
    opts.tenant_id ? "tenant_id=" + opts.tenant_id + "&" : ""
  }${opts.user_group ? "ugroup=" + opts.user_group : ""}`;
};
