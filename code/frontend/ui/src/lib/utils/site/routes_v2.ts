//http://localhost:4000/z/api/:tenant_id/v2
export const apiURL = (tenant_id: string) =>
  `${window.location.origin}/z/api/${tenant_id}/v2`;

//http://localhost:4000
export const baseURL = () => window.location.origin;
export const portalURL = () => `${window.location.origin}/z/portal`;
export const authURL = (opts?: { tenant_id?: string; user_group?: string }) => {
  if (!opts) {
    return `${window.location.origin}/z/auth`;
  }

  return `${window.location.origin}/z/auth?${
    opts.tenant_id ? "tenant_id=" + opts.tenant_id + "&" : ""
  }${opts.user_group ? "ugroup=" + opts.user_group : ""}`;
};
