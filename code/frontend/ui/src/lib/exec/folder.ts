import { FolderTktAPI } from "../apiv2/cabinet";

export const NewFolderApi = (api_base_url: string, token: string) => {
    return new FolderTktAPI(api_base_url, token)
};
