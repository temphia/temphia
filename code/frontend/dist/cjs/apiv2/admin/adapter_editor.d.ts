import { ApiBase } from "../base";
export declare class AdapterEditorAPI {
    base: ApiBase;
    constructor(base_url: string, tenant_id: string, token: string);
    perform_action(name: string, data: any): Promise<import("../http").HttpResponse>;
    self_update(data: any): Promise<import("../http").HttpResponse>;
    self_reset(): Promise<import("../http").HttpResponse>;
    list_apps(): Promise<import("../http").HttpResponse>;
    new_app(data: any): Promise<import("../http").HttpResponse>;
    get_app(id: number): Promise<import("../http").HttpResponse>;
    update_app(id: number, data: any): Promise<import("../http").HttpResponse>;
    delete_app(id: number): Promise<import("../http").HttpResponse>;
    list_hooks(): Promise<import("../http").HttpResponse>;
    new_hook(data: any): Promise<import("../http").HttpResponse>;
    get_hook(id: number): Promise<import("../http").HttpResponse>;
    update_hook(id: number, data: any): Promise<import("../http").HttpResponse>;
    delete_hook(id: number): Promise<import("../http").HttpResponse>;
}
