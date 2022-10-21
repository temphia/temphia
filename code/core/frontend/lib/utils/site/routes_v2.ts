//http://localhost:4000/z/api/:tenant_id/v2
export const apiURL = (tenant_id: string) =>
  `${window.location.origin}/z/api/${tenant_id}/v2`;

//http://localhost:4000
export const baseURL = () => window.location.origin;
export const portalURL = () => `${window.location.origin}/z/portal`;
export const authURL = () => `${window.location.origin}/z/auth`;
