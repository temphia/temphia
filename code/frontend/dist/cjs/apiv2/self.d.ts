import type { ApiBase } from "./base";
export declare class SelfAPI {
    base: ApiBase;
    constructor(base: ApiBase);
    load(): Promise<import("./http").HttpResponse>;
    self(): Promise<import("./http").HttpResponse>;
    self_ws_url(): string;
    user_profile(userid: string): Promise<import("./http").HttpResponse>;
    user_message(userid: string, message: string): Promise<import("./http").HttpResponse>;
    self_update(data: any): Promise<import("./http").HttpResponse>;
    sessions(): Promise<import("./http").HttpResponse>;
    email_change(data: any): Promise<import("./http").HttpResponse>;
    list_message(): Promise<import("./http").HttpResponse>;
    modify_message(data: any): Promise<import("./http").HttpResponse>;
    issue_data(data: any): Promise<import("./http").HttpResponse>;
    issue_folder(data: any): Promise<import("./http").HttpResponse>;
    issue_ugroup(data: any): Promise<import("./http").HttpResponse>;
    list_cabinet_sources(): Promise<import("./http").HttpResponse>;
    list_data_sources(): Promise<import("./http").HttpResponse>;
    list_adapter_providers(): Promise<import("./http").HttpResponse>;
    list_executors(): Promise<import("./http").HttpResponse>;
    list_modules(): Promise<import("./http").HttpResponse>;
    list_repo_sources(): Promise<import("./http").HttpResponse>;
    list_devices(): Promise<import("./http").HttpResponse>;
    add_device(opts: any): Promise<import("./http").HttpResponse>;
    get_device(id: number): Promise<import("./http").HttpResponse>;
    delete_device(id: number): Promise<import("./http").HttpResponse>;
}
