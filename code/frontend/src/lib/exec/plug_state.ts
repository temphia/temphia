import { AdminPlugStateTktAPI } from "../apiv2/admin/plug_state";

export const NewPlugStateApi = (api_base_url: string, token: string) => {
  return new AdminPlugStateTktAPI(api_base_url, token);
};
