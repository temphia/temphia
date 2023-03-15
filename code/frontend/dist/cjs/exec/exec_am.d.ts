export declare class ExecAM {
    api_base_url: string;
    constructor(api_base_url: string);
    new_data_api: (token: string) => import("../apiv2").DataAPI;
    new_folder_api: (token: string) => import("../apiv2").FolderTktAPI;
    new_sockd_room: (token: string) => Promise<import("../sockd").Sockd>;
    new_plug_state: (token: string) => import("../apiv2").AdminPlugStateTktAPI;
}
