import { DataAPI } from "../apiv2/data";

export const NewDataTableApi = (api_base_url: string, token: string) => {
  return new DataAPI(api_base_url, token);
};