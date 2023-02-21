import { Http } from "./http";
export declare class DataSheetAPI {
    http: Http;
    base_url: string;
    token: string;
    constructor(base_url: string, token: string);
    list_sheet_group(): Promise<import("./http").HttpResponse>;
    load_sheet(sheetid: string, options: any): Promise<import("./http").HttpResponse>;
    list_sheets(): Promise<import("./http").HttpResponse>;
    get_sheet(sid: string): Promise<import("./http").HttpResponse>;
    new_sheet(data: any): Promise<import("./http").HttpResponse>;
    update_sheet(sid: string, data: any): Promise<import("./http").HttpResponse>;
    delete_sheet(sid: string): Promise<import("./http").HttpResponse>;
    list_columns(sid: string): Promise<import("./http").HttpResponse>;
    get_column(sid: string, cid: string): Promise<import("./http").HttpResponse>;
    new_column(sid: string, data: any): Promise<import("./http").HttpResponse>;
    update_column(sid: string, cid: string, data: any): Promise<import("./http").HttpResponse>;
    delete_column(sid: string, cid: string): Promise<import("./http").HttpResponse>;
    new_row_cell(sid: string, data: any): Promise<import("./http").HttpResponse>;
    get_row_cell(sid: string, rid: string): Promise<import("./http").HttpResponse>;
    update_row_cell(sid: string, rid: string, data: any): Promise<import("./http").HttpResponse>;
    delete_row_cell(sid: string, rid: string): Promise<import("./http").HttpResponse>;
}
