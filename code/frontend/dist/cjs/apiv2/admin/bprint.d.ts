import type { ApiBase } from "../base";
export declare class AdminBprintAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    list(): Promise<import("../http").HttpResponse>;
    create(data: any): Promise<import("../http").HttpResponse>;
    create_from_zip(data: any): Promise<Response>;
    import(data: any): Promise<import("../http").HttpResponse>;
    get(bid: string): Promise<import("../http").HttpResponse>;
    update(bid: string, data: any): Promise<import("../http").HttpResponse>;
    delete(bid: string): Promise<import("../http").HttpResponse>;
    list_file(bid: string): Promise<import("../http").HttpResponse>;
    get_file(bid: string, file: string): Promise<import("../http").HttpResponse>;
    add_file(bid: string, file: string, data: any): Promise<Response>;
    update_file(bid: string, file: string, data: any): Promise<Response>;
    delete_file(bid: string, file: string): Promise<import("../http").HttpResponse>;
    instance(bid: string, data: any): Promise<import("../http").HttpResponse>;
    issue(bid: string, data: any): Promise<import("../http").HttpResponse>;
    issue_encoded(bid: string, data: any): Promise<import("../http").HttpResponse>;
    list_plugs(bid: string): Promise<import("../http").HttpResponse>;
}
