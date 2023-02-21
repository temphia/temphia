import { Http } from "../http";
export declare class ApiBase {
    http: Http;
    api_base_url: string;
    tenant_id: string;
    user_token: string;
    constructor(api_base_url: string, tenant_id: string, token: string);
    init(): Promise<void>;
    get(path: string): Promise<import("../http").HttpResponse>;
    post(path: string, data: any): Promise<import("../http").HttpResponse>;
    postForm(path: string, auth: boolean, data: any): Promise<Response>;
    patchForm(path: string, auth: boolean, data: any): Promise<Response>;
    put(path: string, data: any): Promise<import("../http").HttpResponse>;
    patch(path: string, data: any): Promise<import("../http").HttpResponse>;
    delete(path: string, data?: any): Promise<import("../http").HttpResponse>;
}
