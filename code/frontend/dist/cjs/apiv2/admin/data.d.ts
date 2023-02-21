import type { ApiBase } from "../base";
export declare class AdminDataAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list_group(source: string): Promise<import("../http").HttpResponse>;
    new_group(source: string, data: any): Promise<import("../http").HttpResponse>;
    get_group(source: string, gid: string): Promise<import("../http").HttpResponse>;
    edit_group(source: string, gid: string, data: any): Promise<import("../http").HttpResponse>;
    delete_group(source: string, gid: string): Promise<import("../http").HttpResponse>;
    list_tables(source: string, gid: string): Promise<import("../http").HttpResponse>;
    get_table(source: string, gid: string, tid: string): Promise<import("../http").HttpResponse>;
    edit_table(source: string, gid: string, tid: string, data: any): Promise<import("../http").HttpResponse>;
    delete_table(source: string, gid: string, tid: string): Promise<import("../http").HttpResponse>;
    list_column(source: string, gid: string, tid: string): Promise<import("../http").HttpResponse>;
    get_column(source: string, gid: string, tid: string, cid: string): Promise<import("../http").HttpResponse>;
    edit_column(source: string, gid: string, tid: string, cid: string, data: any): Promise<import("../http").HttpResponse>;
    delete_column(source: string, gid: string, tid: string, cid: string): Promise<import("../http").HttpResponse>;
    list_view(source: string, gid: string, tid: string): Promise<import("../http").HttpResponse>;
    add_view(source: string, data: any, gid: string, tid: string): Promise<import("../http").HttpResponse>;
    get_view(source: string, gid: string, tid: string, id: string): Promise<import("../http").HttpResponse>;
    edit_view(source: string, gid: string, tid: string, id: string, data: any): Promise<import("../http").HttpResponse>;
    delete_view(source: string, gid: string, tid: string, id: string): Promise<import("../http").HttpResponse>;
    seed_table: (source: string, gid: string, tid: string, max: number) => Promise<import("../http").HttpResponse>;
    query: (source: string, gid: string, opts: any) => Promise<import("../http").HttpResponse>;
    list_table_activity: (source: string, gid: string, tid: string, offset: number) => Promise<import("../http").HttpResponse>;
}
