import { DataSheetAPI } from "./data_sheet";
import { Http } from "./http";
export declare class DataAPI {
    http: Http;
    api_base_url: string;
    token: string;
    constructor(api_base_url: string, token: string);
    load(): Promise<import("./http").HttpResponse>;
    new_row(tid: string, data: any): Promise<import("./http").HttpResponse>;
    get_row(tid: string, rid: string): Promise<import("./http").HttpResponse>;
    update_row(tid: string, rid: string, data: any): Promise<import("./http").HttpResponse>;
    delete_row(tid: string, rid: string): Promise<import("./http").HttpResponse>;
    load_table(tid: string, view?: string): Promise<import("./http").HttpResponse>;
    simple_query(tid: string, query: any): Promise<import("./http").HttpResponse>;
    ref_load(tid: string, data: any): Promise<import("./http").HttpResponse>;
    ref_resolve(tid: string, data: any): Promise<import("./http").HttpResponse>;
    reverse_ref_load(tid: string, data: any): Promise<import("./http").HttpResponse>;
    list_activity(tid: string, rid: string): Promise<import("./http").HttpResponse>;
    comment_row(tid: string, rid: string, data: any): Promise<import("./http").HttpResponse>;
    sockd_url: () => string;
    sheet_api: () => DataSheetAPI;
    list_users: (opts: any) => Promise<import("./http").HttpResponse>;
}
