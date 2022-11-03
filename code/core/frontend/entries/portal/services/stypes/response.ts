export interface SelfLoad {
  tenant_name: string;
  tenant_id: string;
  user_info: object;
  scopes: string[];
  plug_apps: object[];
}
